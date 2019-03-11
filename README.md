# Form3 Service

The form3-service (The service) is a service responsible to manage payment resources. It allows to create, delete payments; to update payment's beneficiary and to retrieve a single payment or a list of payment.

The service persist each payment's state, so that we can have a history of all the changes made to a payment.   

## Table of Contents

- [Getting started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Getting the source](#getting-the-source)
    - [Development](#development)
        - [Application run](#application-run)
        - [Generate documentation](#generate-documentation)
    - [Testing](#testing)
- [Troubleshooting](#troubleshooting)
    - [Known issues](#known-issues)
    
## Getting started

### Prerequisites

You need to make sure that you have `go1.11` or later, `make` and `docker` installed

```
$ which make
/usr/bin/make
$ which go
/usr/local/bin/go
$ which docker
/usr/local/bin/docker
```

There is no other prerequisite needed in order to setup this project for development.

[[table of contents]](#table-of-contents)

### Getting the source

Setup the project structure and fetch the repo like so:
 
```bash
go get github.com/dohernandez/form3-service
```

### Development

This project follows the following structure:

```markdown
|-- cmd # MUST be used as a main entrypoint, one folder for each binary
	|-- servid # For simple application logic setup is done here
		|-- servid.go
|-- internal # contains application specific non-reusable by any other projects code 
	|-- domain # domain packages
		|-- transaction
			|-- payment.go
			|-- charges.go
	|-- platform # foundational packages specific to the project
		|-- app # MUST contains base standard definitions to setup service.
		|-- event
			|-- store # contains event aggregate type
				|-- payment.go
		|-- http
			|-- handlers # http handler grouped by domain bundle
				|-- transaction
				        |-- payment
                            |-- decode.go # MUST contains decode request func
                            |-- encode.go # MUST contains encode response func
                            |-- post.go # MUST contains handler func
			|-- routes.go  # MUST contains routes
			|-- config.go # MUST contains the service configuration
			|-- container.go # MUST contains service resources
			|-- init.go # MUST initialize the service resources
		|-- projection
			|-- handler # projection handler grouped by projection
				|-- message
				        |-- payment.go
		|-- storage # MUST contains the abstraction of data (removing, updating, and selecting items from collection)
		    |-- payment.go	            
|-- pkg # MUST NOT import internal packages. Packages placed here should be considered as vendor.
	|-- http
		|-- rest
			|-- request
	|-- log
|-- resources # RECOMMENDED service resources. Shell helper scripts, additional files required for development, documentations.
	|-- migrations # Migration files
	|-- docs # MUST contains project documentation in human and/or machine readable format
|-- features # OPTIONAL, place to store specification definitions in gherkin format
|-- .env.template # MUST contains the env variables used by the service.
|-- .editorconfig # OPTIONAL https://editorconfig.org
|-- .gitignore
|-- docker-compose.yml
|-- Dockerfile
|-- Gopkg.lock
|-- Gopkg.toml
|-- README.md
```

**Package Design**

`cmd/`
    
    * Packages that provide support for a specific program that is being built.
    * Can only import package from `internal/platform` and `pgk`.
    * Can't import package from `internal/domain`.
    * Allowed to panic an application.
    * Wrap errors with context if not being handled.
    * Majority of handling errors happen here.
    * Can recover any panic.
    * Only if system can be returned to 100% integrity.
    
`pkg`
    
    * Can't import import `internal` packages. 
    * Packages placed here should be considered as vendor.
    * Stick to the testing package in go.
    * NOT allowed to panic an application.
    * NOT allowed to wrap errors.
    * Return only root cause error values.
    * NOT allowed to set policy about any application concerns.
    * NOT allowed to log, but access to trace information must be decoupled.
    * Configuration and runtime changes must be decoupled.
    * Retrieving metric and telemetry values must be decoupled.
    * Stick to the testing package in go.
    * Test files belong inside the package.
    * Focus more on unit than integration testing.
    
`internal\domain`
    
    * NOT allowed to panic an application.
    * Allowed to wrap errors when domain concern.
    * Wrap errors with context if not being handled.
    * Allowed to set policy about any application concerns.
    * Allowed to log and handle configuration natively.
    * Minority of handling errors happen here.
    * Stick to the testing package in go.
    * Test files belong inside the package.
    * Focus more on unit than integration testing.
    * Package at the same level are not allowed to import each other.
    * Package root can import subpackages.
    * Can't import `internal\platform` package

`internal\platform`
    
    * NOT allowed to panic an application.
    * NOT allowed to set policy about any application concerns.
    * NOT allowed to log, but access to trace information must be decoupled.
    * Configuration and runtime changes must be decoupled.
    * Retrieving metric and telemetry values must be decoupled.
    * Return only root cause error values.
    * Stick to the testing package in go.
    * Test files belong inside the package.
    * Focus more on unit than integration testing.
    * Packages can import each other.
    * Can import `internal\domain` package
    
This structure design is mostly inspired by [Package Oriented Design](https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html) by William Kennedy.

Routine operations are defined in `Makefile`.

```bash
form3-service routine operations

  init:                 Init the application, usage: "make init API_PORT=<service-api-port> POSTGRES_PORT=<postgres-port>"
                                               
                        Requirement:
                          export FORM3_SERVICE_HOST_PORT=<service-api-port>
                          export FORM3_POSTGRES_HOST_PORT=<postgres-port>
                       
                        Arguments:
                          API_PORT              Requires port to run the service
                          POSTGRES_PORT         Requires port to run the postgres


       -- Misc --

  build:                Build binary
  run:                  Run application (before exec this command make sure `make init` was executed)
  run-compile-daemon:   Run application with CompileDaemon (automatic rebuild on code change)
  lint:                 Check with golangci-lint
  fix-lint:             Apply goimports and gofmt
  deps:                 Ensure dependencies according to toml file
  deps-vendor:          Ensure dependencies according to lock file

       -- Environment modifiers --

  env:                  Run command with .env vars (before exec this command make sure `make init` was executed)
  envfile:              Check/Generate .env file based on .env.template if not exists

       -- Test --

  test:                 Run tests
  test-unit:            Run unit tests
  test-integration:     Run integration tests, usage: make test-integration [TAGS=<tags-splitted-by-comma>] [FEATURE=<tags-splitted-by-comma>]
                       
                        Arguments:
                          TAGS     Optional tag(s) to run. Filter scenarios by tags:
                                   - "@dev": run all scenarios with wip tag
                                   - "~@notImplemented": exclude all scenarios with wip tag
                                   - "@dev && ~@notImplemented": run wip scenarios, but exclude new
                                   - "@dev,@undone": run wip or undone scenarios
                          FEATURE  Optional feature to run. Run only the specified feature.
                       
                        Examples:
                          only scenarios: "make test-integration TAGS=@dev"
                          only one feature: "make test-integration FEATURE=Dev"

       -- Documentation --

  docs:                 Generate api documentation (raml)

       -- Database migrations --

  create-migration:     Create database migration file, usage: "make create-migration NAME=<migration-name>"
  migrate:              Apply migrations
  migrate-cli:          Check/install migrations tool

       -- Docker --

  docker:               Run command with docker-compose (before exec this command make sure `make init` was executed)
                       
                        Examples:
                          run migration: "make docker migrate"
                          run test: "make docker test"

       -- API service --

  servid-start:         Start API service (before exec this command make sure `make init` was executed)
  servid-stop:          Stop API service
  servid-api-log:       Log API service

Usage
  make <flags> [options]
```

[[table of contents]](#table-of-contents)

#### Application run

The first thing you need to do is, init the application, create the `.env` file with the server configuration and set up the environment variable `FORM3_SERVICE_HOST_PORT`. To do so, run the command

```bash
make init API_PORT=8008 POSTGRES_PORT=5434
```

```bash
>> initializing .env file
>> ensuring dependencies
>> run those commands to set the value to the variables env
export FORM3_SERVICE_HOST_PORT=8008
export FORM3_POSTGRES_HOST_PORT=5434

```

It will create the `.env` file for you based on `.env.template` file and print how to set the environment variable require in your environment.

After init the application and export the environment variables, you are able to start/stop the service at any time. 

```bash
make servid-start
```

```bash
>> starting API service in port 8008 and postgres in port 5434
Creating network "form3-service_default" with the default driver
...
WARNING: Image for service api was built because it did not already exist. To rebuild this image you must use `docker-compose build` or `docker-compose up --build`.
Creating form3-service_postgres_1 ... done
Creating form3-service_api_1      ... done

```

**Note** Wait a bit until the service is up and running, run `make servid-api-log` to check when the service is ready

```bash
api_1       | >> checking/installing migrations tool
api_1       | >> installing migrate cli
api_1       | >> running migrations
api_1       | 20190227153745/u create_table_events_transaction_stream (63.9807ms)
api_1       | 20190303224329/u create_table_transaction_payment (102.8258ms)
api_1       | 20190303234557/u create_table_transaction_projections (161.2766ms)
api_1       | >> running app with CompileDaemon
api_1       | 2019/03/08 01:13:52 Running build command!
api_1       | 2019/03/08 01:14:02 Build ok.
api_1       | 2019/03/08 01:14:02 Restarting the given command.
api_1       | 2019/03/08 01:14:02 stderr: {"level":"info","message":"Creating routers","timestamp":"2019-03-08T01:14:02Z"}
api_1       | 2019/03/08 01:14:02 stderr: {"level":"debug","message":"added `/` route","timestamp":"2019-03-08T01:14:02Z"}
api_1       | 2019/03/08 01:14:02 stderr: {"level":"debug","message":"added `/version` route","timestamp":"2019-03-08T01:14:02Z"}
api_1       | 2019/03/08 01:14:02 stderr: {"level":"debug","message":"added `/status` route","timestamp":"2019-03-08T01:14:02Z"}
api_1       | 2019/03/08 01:14:02 stderr: {"level":"debug","message":"added `/health` route","timestamp":"2019-03-08T01:14:02Z"}
api_1       | 2019/03/08 01:14:02 stderr: {"level":"debug","message":"added `/docs` route","timestamp":"2019-03-08T01:14:02Z"}
api_1       | 2019/03/08 01:14:02 stderr: {"level":"debug","message":"added `POST /v1/transaction/payments` route","timestamp":"2019-03-08T01:14:02Z"}
api_1       | 2019/03/08 01:14:02 stderr: {"level":"debug","message":"added `PATCH /v1/transaction/payments/{id}/beneficiary` route","timestamp":"2019-03-08T01:14:02Z"}
api_1       | 2019/03/08 01:14:02 stderr: {"level":"debug","message":"added `DELETE /v1/transaction/payments/{id}` route","timestamp":"2019-03-08T01:14:02Z"}
api_1       | 2019/03/08 01:14:02 stderr: {"level":"debug","message":"added `GET /v1/transaction/payments/{id}` route","timestamp":"2019-03-08T01:14:02Z"}
api_1       | 2019/03/08 01:14:02 stderr: {"level":"debug","message":"added `GET /v1/transaction/payments` route","timestamp":"2019-03-08T01:14:02Z"}
api_1       | 2019/03/08 01:14:02 stderr: {"level":"info","message":"Starting server at port http://0.0.0.0:8000","timestamp":"2019-03-08T01:14:02Z"}
```

To stop the service run

```bash
make servid-stop
```

```bash
>> stop API service in port 8008 and postgres in port 5434
Stopping form3-service_api_1 ... done
Stopping fform3-service_postgres_1 ... done
Removing form3-service_api_1 ... done
Removing form3-service_postgres_1 ... done
Removing network form3-service_default
```

[[table of contents]](#table-of-contents)

#### Generate documentation

Documentation items are generated using raml generator. RAML file is located `resources/raml/api.raml`. 

To update api documentation, run `make docs`.

To see the api documentation generated, you can access to the root of the service [http://localhost:8008](http://localhost:8008), it will show the link to the api documentation.

```html
Welcome to form3-service. Please read API <a href="http://localhost:8008/docs/api.html">documentation</a>.
```


[[table of contents]](#table-of-contents)

### Testing 

Before you can run the complete suite tests (unit test and behavioral test), make sure `.env` file is created and  `docker-compose` services to your `/etc/hosts`:

```
127.0.0.1 postgres
```                                                                                                                                               

then you can run

```
make env test
```

otherwise see routine operations defined in `Makefile` to run each suite independently.


Another way to run the complete suite tests is using docker where there is no need to add any entry into your `/etc/hosts`:

```
make docker test
```

This is the most simple way to quick start testing your app after cloning a repo, though it has low performance and is harder to debug.

[[table of contents]](#table-of-contents)

## Troubleshooting

### Known issues

There are no known issues.

[[table of contents]](#table-of-contents)