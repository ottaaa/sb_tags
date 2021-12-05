package main

import (
	"github.com/urfave/cli/v2"
)

func action(c *cli.Context) error {
	var arg = c.Args().Get(0)

	project := NewProject(arg)
	for _, tag := range project.tags() {
		println(tag)
	}

	return nil
}
