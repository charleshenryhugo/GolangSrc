package main

import (
	"log"

	"github.com/urfave/cli"
)

//GenNotify operate all possible notifications
func GenNotify() error {
	//parse notifiers from notifiers config file
	ntfs, err := parseNotifiers(notifyrcFile)
	if err != NIL {
		return cli.NewExitError("", int(err))
	}

	//do email notify
	if err := EmailNotify(ToEmailAddrs, Subject, Message, ntfs); err == NIL {
		log.Println("email notification success")
	} else if err == SMTPM_INVAL {
		log.Println("email notification invalid")
	} else if err == SMTPM_NOTGT {
		log.Println("no target email address(es)")
	} else {
		defer cli.OsExiter(int(err))
	}

	//do slack notify
	if _, _, err := SlackNotify(ToSlackUsers, Subject, Message, ntfs); err == NIL {
		log.Println("slack notification success")
	} else if err == SLK_NOTGT {
		log.Println("no target slack users(channels)")
	} else if err == SLK_INVAL {
		log.Println("slack notification invalid")
	} else {
		defer cli.OsExiter(int(err))
	}
	return nil
}
