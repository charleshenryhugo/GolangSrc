package main

import (
	"flag"
	"io/ioutil"
	"strings"
)

//define some command line options
var (
	subjectOpt         = flag.String("s", "", "Specify the title/subject of your notification")
	messageOpt         = flag.String("m", "", "Specify the message of your notification")
	messageFileOpt     = flag.String("f", "", "Specify the file that stores your notification-message")
	toEmailAddrOpt     = flag.String("e", "", "Specify the target emails. Do nothing if the email state is off")
	toEmailAddrFileOpt = flag.String("el", "", "Specify the file that stores target email list.\n Do nothing if the email state is off")
	toSlackUserOpt     = flag.String("k", "", "Specify the target slack userIDs. Do nothing if the slack state is off")
	toSlackUserFileOpt = flag.String("kl", "", "Specify the file that stores target slack userID list.\n Do nothing if the email state is off")

	//do some default settings

	//to be added ... ...
)

//parse command line options and return these options
//to be added ... ...
func cmdParse() (emails, slacks []string, subject string, msg string) {
	flag.Parse()
	//the notification subject, default "No Subject"
	subject = *subjectOpt

	//the notification message, default "No Message"
	msg = *messageOpt

	//file(maybe error.log) that stores the notification message, default "No File"
	msgFile := *messageFileOpt

	//email list containing email-addrs that will be notified
	to := *toEmailAddrOpt
	emails = strings.Fields(to)

	//slack ID list containing slack IDs that will be notified
	to = *toSlackUserOpt
	slacks = strings.Fields(to)

	//append those email addrs stored in the file, only if the file is available
	emailAddrFile := *toEmailAddrFileOpt
	if fileBytes, err := ioutil.ReadFile(emailAddrFile); err == nil {
		emails = append(strings.Fields(string(fileBytes)), emails...)
		//emails = strings.Fields(string(fileBytes))
	}
	//append those slack user IDs stored in the file, only if the file is available
	slackUserFile := *toSlackUserFileOpt
	if fileBytes, err := ioutil.ReadFile(slackUserFile); err == nil {
		slacks = append(strings.Fields(string(fileBytes)), slacks...)
		//slacks = strings.Fields(string(fileBytes))
	}

	//get message from the file(usually error.log), only if the file is available
	if fileBytes, err := ioutil.ReadFile(msgFile); err == nil { //only when the msg file is read successfully, we rewrite the msg
		msg = string(fileBytes)
	}

	//apply the default settings to message, subject, emails or slacks if any of them is empty
	dflt := parseDefaults("defaults.xml")
	if msg == "" {
		msg = dflt.GetDfltmsg()
	}
	if subject == "" {
		subject = dflt.GetDfltSbjt()
	}
	if len(emails) == 0 {
		emails = dflt.GetDfltEmailList()
	}
	if len(slacks) == 0 {
		slacks = dflt.GetDfltSlackList()
	}
	//return email addr list, slack user IDs list
	//and subject and msgs of notifications
	return emails, slacks, subject, msg
}
