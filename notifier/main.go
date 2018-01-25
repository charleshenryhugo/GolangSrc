package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getStdin() string {
	fmt.Println("Input message to be sent. Press enter if you want to use the default message")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	ntfs := parseNotifiers(notifyrcFile)
	emails, slacks, subject, msg := cmdParse()

	// scanMsg := getStdin()
	// if scanMsg != "" {
	// 	msg = scanMsg
	// }

	if err := EmailNotify(emails, subject, msg, ntfs); err == nil {
		log.Println("Email Notification Successfully")
	} else {
		log.Println(err)
	}
	if _, _, err := SlackNotify(slacks, subject, msg, ntfs); err == nil {
		log.Println("Slack Notification Successfully")
	} else {
		log.Println(err)
	}

}
