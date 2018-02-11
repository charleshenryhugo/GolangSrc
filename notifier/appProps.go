package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/urfave/cli"
)

//global input parameters
//to be added for more notifiers
var (
	Subject          string
	Message          string
	MessageFile      string
	ToEmailAddrs     []string
	ToEmailAddrsFile string
	ToSlackUsers     []string
	ToSlackUsersFile string
)

//usage of global input parameters
//to be added for more notifiers

const (
	subjectFlgUsg          = "Specify the title/subject of your notification (UTF-8, maximum 256 bytes for email notification)"
	messageFlgUsg          = "Specify the message of your notification (UTF-8)"
	msgFileFlgUsg          = "Specify the file that stores your notification message (UTF-8)"
	toEmailAddrsFlgUsg     = "Specify the target email address(es). Do nothing if the email state is off"
	toSlackUsersFlgUsg     = "Specify the target slack userID(s). Do nothing if the slack state is off"
	toEmailAddrsFileFlgUsg = "Specify the file that stores target email address list (one address per line). Do nothing if the email state is off"
	toSlackUsersFileFlgUsg = "Specify the file that stores target slack userID list (one address per line). Do nothing if the email state is off"
)

func appInit() *cli.App {
	app := cli.NewApp()

	app.Name = "Notifier"
	app.Usage = "A simple tool for notification"
	app.HelpName = "Notifier"
	app.Version = "1.0.0"
	app.Author = "ZHU YUE"

	return app
}

func appAction(ctx *cli.Context) error {
	//append those email addrs stored in the file, only if the file is available
	if fileBytes, err := ioutil.ReadFile(ToEmailAddrsFile); err == nil {
		ToEmailAddrs = append(strings.Fields(string(fileBytes)), ToEmailAddrs...)
	}
	//append those slack user IDs stored in the file, only if the file is available
	if fileBytes, err := ioutil.ReadFile(ToSlackUsersFile); err == nil {
		ToSlackUsers = append(strings.Fields(string(fileBytes)), ToSlackUsers...)
	}
	//get message from the file(usually error.log), only if the file is available
	if fileBytes, err := ioutil.ReadFile(MessageFile); err == nil {
		Message = string(fileBytes)
	}
	//apply the default settings to message, subject, emails or slacks
	//if any of them is empty
	dflt, err := parseDefaults(defaultsFile)
	if err == NIL {
		//Apply default settings for any empty CLI flags
		if Message == "" {
			Message = dflt.GetDfltmsg()
		}
		if Subject == "" {
			Subject = dflt.GetDfltSbjt()
		}
		if len(ToEmailAddrs) == 0 {
			ToEmailAddrs = dflt.GetDfltEmailList()
		}
		if len(ToSlackUsers) == 0 {
			ToSlackUsers = dflt.GetDfltSlackList()
		}
	}

	//parse notifiers from notifiers config file
	ntfs, err := parseNotifiers(notifyrcFile)
	if err != NIL {
		return cli.NewExitError("", int(err))
	}

	//do email notify
	if err := EmailNotify(ToEmailAddrs, Subject, Message, ntfs); err == NIL {
		fmt.Println("email notification success")
	} else if err == SMTPM_INVAL {
		fmt.Println("email notification invalid")
	} else {
		defer cli.OsExiter(int(err))
	}

	//do slack notify
	if _, _, err := SlackNotify(ToSlackUsers, Subject, Message, ntfs); err == NIL {
		fmt.Println("slack notification success")
	} else if err == SLK_INVAL {
		fmt.Println("slack notification invalid")
	} else {
		defer cli.OsExiter(int(err))
	}
	return nil
}

func appFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "subject, s",
			Usage:       subjectFlgUsg,
			Destination: &Subject,
		},
		cli.StringFlag{
			Name:        "msg, m",
			Usage:       messageFlgUsg,
			Destination: &Message,
		},
		cli.StringFlag{
			Name:        "msgfile, mf",
			Usage:       msgFileFlgUsg,
			Destination: &MessageFile,
		},
		cli.StringFlag{
			Name:        "emails-file, ef",
			Usage:       toEmailAddrsFileFlgUsg,
			Destination: &ToEmailAddrsFile,
		},
		cli.StringFlag{
			Name:        "slacks-file, kf",
			Usage:       toSlackUsersFileFlgUsg,
			Destination: &ToSlackUsersFile,
		},
		cli.StringSliceFlag{
			Name:  "email-addrs, e",
			Usage: toEmailAddrsFlgUsg,
		},
		cli.StringSliceFlag{
			Name:  "slack-ids, k",
			Usage: toSlackUsersFlgUsg,
		},
	}
}

func appCommands() []cli.Command {
	return []cli.Command{
		//change default settings(modify config file)
		{
			Name:    "setdefault",
			Aliases: []string{"default", "def"},
			Usage:   "Change(set) default settings (with some subcommands)",
			Subcommands: []cli.Command{
				{
					Name:    "message",
					Aliases: []string{"msg"},
					Usage:   "Change(set) default message to be sent",
					Action: func(c *cli.Context) error {
						newMsg := c.Args().First()
						return CfgDfltMsg(newMsg)
					},
				},
				{
					Name:        "subject",
					Aliases:     []string{"title", "sbjt"},
					Usage:       "Change(set) default subject/title to be sent",
					Description: "hahahah",
					Action: func(c *cli.Context) error {
						newSbjt := c.Args().First()
						return CfgDfltSbjt(newSbjt)
					},
				},
				{
					Name:    "messageFile",
					Aliases: []string{"msgFile", "msgf"},
					Usage:   "Change(set) default file name which stores message",
					Action: func(c *cli.Context) error {
						newMsgFile := c.Args().First()
						return CfgDfltMsgFile(newMsgFile)
					},
				},
				{
					Name:    "slackListFile",
					Aliases: []string{"kfile", "kf"},
					Usage:   "Change(set) default file name which stores target slack userID(s)",
					Action: func(c *cli.Context) error {
						newSlackListFile := c.Args().First()
						return CfgDfltSlackListFile(newSlackListFile)
					},
				},
				{
					Name:    "emailListFile",
					Aliases: []string{"efile", "ef"},
					Usage:   "Change(set) default file name which stores target email address(es)",
					Action: func(c *cli.Context) error {
						newEmailListFile := c.Args().First()
						return CfgDfltEmailListFile(newEmailListFile)
					},
				},
			},
		},
		//change notifiers settings(modify config file)
		{
			Name:        "setnotif",
			Aliases:     []string{"notif"},
			Usage:       "Change(set) notifiers settings, (e.g. slack token, email account)",
			Subcommands: []cli.Command{},
		},
		{
			Name:    "toggle",
			Aliases: []string{"tog"},
			Usage:   "toggle notifier state between 'on' and 'off' ",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "email",
					Usage: "toggle email notifier state",
				},
				cli.BoolFlag{
					Name:  "slack",
					Usage: "toggle slack notifier state",
				},
			},
			Action: func(ctx *cli.Context) error {
				if ctx.Bool("email") {
					CfgToggStat(emailNotifier)
				}
				if ctx.Bool("slack") {
					CfgToggStat(slackNotifier)
				}
				return nil
			},
		},
	}
}
