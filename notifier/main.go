package main

import (
	"os"
	"sort"

	"github.com/urfave/cli"
)

func main() {
	//build a new app with cli package, and specify some info
	app := appInit()

	//define flags
	app.Flags = appFlags()

	//define commands and subcommands(including actions)
	app.Commands = appCommands()

	//define app action
	app.Action = func(ctx *cli.Context) error {
		ToEmailAddrs = ctx.StringSlice("email-addrs")
		ToSlackUsers = ctx.StringSlice("slack-ids")
		return appAction(ctx)
	}

	//sort flags and commands
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}
