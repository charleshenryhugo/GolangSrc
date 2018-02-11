package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func show(values ...interface{}) {
	for _, value := range values {
		fmt.Println(value)
	}
}

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:  "test",
			Usage: "just a test",
		},
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "hot",
			Usage: "if hot?",
		},
	}

	app.Action = func(c *cli.Context) {
		log.Println(c.Bool("hot"))
	}
	app.Run(os.Args)
}
