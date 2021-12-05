package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Action: action,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
