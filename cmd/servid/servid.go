package main

import (
	"fmt"
	"os"

	"github.com/dohernandez/form3-service/pkg/version"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println(version.Info().String())

		return
	}
}
