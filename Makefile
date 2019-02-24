# detecting GOPATH and removing trailing "/" if any
GOPATH = $(realpath $(shell go env GOPATH))
IMPORT_PATH = $(subst $(GOPATH)/src/,,$(PWD))

export SERVICE_NAME ?= $(subst github.com/dohernandez/,,$(IMPORT_PATH))

APP_PATH ?= $(PWD)
APP_SCRIPTS_PATH ?= $(APP_PATH)/resources/app/scripts

branch = $(shell git symbolic-ref HEAD 2>/dev/null)
VERSION ?= $(shell git symbolic-ref -q --short HEAD || git describe --tags --exact-match)
revision = $(shell git log -1 --pretty=format:"%H")
build_user = $(USER)
build_date = $(shell date +%FT%T%Z)

VERSION_PKG = $(APP_PATH)/pkg/version
export LDFLAGS = -X $(VERSION_PKG).version=$(VERSION) -X $(VERSION_PKG).branch=$(branch) -X $(VERSION_PKG).revision=$(revision) -X $(VERSION_PKG).buildUser=$(build_user) -X $(VERSION_PKG).buildDate=$(build_date)

BUILD_DIR ?= bin

## -- Misc --

## Init the application
init: envfile deps

## Build binary
build:
	@echo ">> building binary"
	@go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/${SERVICE_NAME} cmd/${SERVICE_NAME}/*

## Run application (before exec this command make sure `make init` was exec before)
run:
	@echo ">> running app"
	@go run -ldflags "${LDFLAGS}" cmd/"${SERVICE_NAME}"/*

## Run application with CompileDaemon (automatic rebuild on code change)
run-compile-daemon:
	@test -s $(shell go env GOPATH)/bin/CompileDaemon || (echo ">> installing CompileDaemon" && go get -u github.com/githubnemo/CompileDaemon)
	@echo ">> running app with CompileDaemon"
	@$(shell go env GOPATH)/bin/CompileDaemon -exclude-dir=vendor -build='make build migrate' -command='$(BUILD_DIR)/${SERVICE_NAME}' -graceful-kill

## Check with golangci-lint
lint:
	@$(APP_SCRIPTS_PATH)/check-lint.sh

## Apply goimports and gofmt
fix-lint:
	@$(APP_SCRIPTS_PATH)/fix-style.sh

## Ensure dependencies according to toml file
deps:
	@echo ">> ensuring dependencies"
	@test -s $(GOPATH)/bin/dep || GOBIN=$(GOPATH)/bin go get -u github.com/golang/dep/cmd/dep
	@$(GOPATH)/bin/dep ensure
	@git add ${APP_PATH}/Gopkg.lock

## Ensure dependencies according to lock file
deps-vendor:
	@echo ">> ensuring dependencies"
	@test -s $(GOPATH)/bin/dep || GOBIN=$(GOPATH)/bin go get -u github.com/golang/dep/cmd/dep
	@$(GOPATH)/bin/dep ensure --vendor-only
	@git add ${APP_PATH}/Gopkg.lock
	
## -- Environment modifiers --

## Run command with .env vars (before exec this command make sure `make init` was exec before)
env:
	@echo ">> running with .env"
	@$(APP_SCRIPTS_PATH)/env-run.sh make $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
	@kill -3 $$PPID
	@echo "Job done, stopping make, please disregard following 'make: *** [env] Error 1'"
	@exit 1

## Generate .env file based on .env.template
envfile:
	@test -s ./.env || (echo ">> copying .env.template to .env" && cp .env.template .env)

## -- Test --

## Run tests
test: test-unit test-integration

## Run unit tests
test-unit:
	@echo ">> unit test"
	@test -s $(GOPATH)/bin/overalls || GOBIN=$(GOPATH)/bin go get -u github.com/go-playground/overalls
	@$(GOPATH)/bin/overalls -project=${IMPORT_PATH} -covermode=atomic -- -race

## Run integration tests
##
## Arguments:
##   TAGS     Optional tag(s) to run. Filter scenarios by tags:
##            - "@dev": run all scenarios with wip tag
##            - "~@notImplemented": exclude all scenarios with wip tag
##            - "@dev && ~@notImplemented": run wip scenarios, but exclude new
##            - "@dev,@undone": run wip or undone scenarios
##   FEATURE  Optional feature to run. Run only the specified feature.
##
## Examples:
##   only scenarios: 'make test-integration TAGS=@dev'
##   only one feature: 'make test-integration FEATURE=Dev'
test-integration:
	@echo ">> integration test"
	@test -s $(GOPATH)/bin/overalls || GOBIN=$(GOPATH)/bin go get -u github.com/go-playground/overalls
	@$(GOPATH)/bin/overalls -project=${IMPORT_PATH}/features -- -coverpkg ${IMPORT_PATH}/internal/... -godog -stop-on-failure -tag="${TAGS}" -feature="${FEATURE}"

ifeq ($(findstring run,$(MAKECMDGOALS)),run)
    DOCKER_SERVICE_PORTS=--service-ports
endif

## -- Docker --

## Run command with docker-compose (before exec this command make sure `make init` was exec before)
docker:
	@echo ">> running with docker-compose"
	@docker-compose run $(DOCKER_SERVICE_PORTS) --rm app make $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
	@kill -3 $$PPID
	@echo "Job done, stopping make, please disregard following 'make: *** [docker-tool] Error 1'"
	@exit 1

.PHONY: init build run run-compile-daemon lint fix-lint deps test test-unit test-integration docker help

.DEFAULT_GOAL := help
HELP_SECTION_WIDTH="      "
HELP_DESC_WIDTH="                       "
help:
	@printf "form3-service routine operations\n";
	@awk '{ \
			if ($$0 ~ /^.PHONY: [a-zA-Z\-\_0-9]+$$/) { \
				helpCommand = substr($$0, index($$0, ":") + 2); \
				if (helpMessage) { \
					printf "  \033[32m%-20s\033[0m %s\n", \
						helpCommand, helpMessage; \
					helpMessage = ""; \
				} \
			} else if ($$0 ~ /^[a-zA-Z\-\_0-9.]+:/) { \
				helpCommand = substr($$0, 0, index($$0, ":")); \
				if (helpMessage) { \
					printf "  \033[32m%-20s\033[0m %s\n", \
						helpCommand, helpMessage; \
					helpMessage = ""; \
				} \
			} else if ($$0 ~ /^##/) { \
				if (helpMessage) { \
					helpMessage = helpMessage"\n"${HELP_DESC_WIDTH}substr($$0, 3); \
				} else { \
					helpMessage = substr($$0, 3); \
				} \
			} else { \
				if (helpMessage) { \
					print "\n"${HELP_SECTION_WIDTH}helpMessage"\n" \
				} \
				helpMessage = ""; \
			} \
		}' \
		$(MAKEFILE_LIST)
	@printf "\nUsage\n";
	@printf "  make  [flags] [options]";