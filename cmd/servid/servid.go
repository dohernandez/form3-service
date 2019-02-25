package main

import (
	"fmt"
	"os"

	"github.com/dohernandez/form3-service/internal/platform/app"
	"github.com/dohernandez/form3-service/pkg/version"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println(version.Info().String())

		return
	}

	cfg, err := app.LoadEnv()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	c, err := app.NewAppContainer(cfg)
	if err != nil {
		panic("failed init application container: " + err.Error())
	}
}
