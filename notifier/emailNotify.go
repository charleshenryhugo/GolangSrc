package main

import (
	"bytes"
	"crypto/tls"
	"errors"
	"log"
	"net/smtp"
	"strings"
)

type Mail struct {
	senderID string
	toIds    []string
	subject  string
	body     string
}

type SmtpServer struct {
	host      string
	port      string
	tlsconfig *tls.Config
}

func (s *SmtpServer) ServerName() string {
	return s.host + ":" + s.port
}

//Email has specific form so we need to build it
func (mail *Mail) BuildMessage() string {
	message := bytes.NewBufferString("From: ")
	message.WriteString(mail.senderID)
	message.WriteString("\r\n")
	message.WriteString("To: ")
	if len(mail.toIds) > 0 {
		message.WriteString(strings.Join(mail.toIds, ";"))
		message.WriteString("\r\n")
	}
	message.WriteString("Subject: ")
	message.WriteString(mail.subject)
	message.WriteString("\r\n\r\n")
	message.WriteString(mail.body)

	return message.String()
}

func newMail(from string, to []string, subject string, body string) *Mail {
	mail := new(Mail)
	mail.senderID = from
	mail.toIds = to
	mail.subject = subject
	mail.body = body

	return mail
}
func newSMTPServer(host string, port string) *SmtpServer {
	smtpServer := new(SmtpServer)
	smtpServer.host = host
	smtpServer.port = port
	smtpServer.tlsconfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.host,
	}

	return smtpServer
}

//send email using SMTP with a specific SMTP server and an account
func smtpEmail(mail *Mail, smtpServer *SmtpServer, pwd string) error {
	msgBody := mail.BuildMessage()
	log.Println("connecting server", smtpServer.ServerName())

	//build an authentication
	auth := smtp.PlainAuth("", mail.senderID, pwd, smtpServer.host)

	conn, err := tls.Dial("tcp", smtpServer.ServerName(), smtpServer.tlsconfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		return err
	}

	//Use Auth
	if err = client.Auth(auth); err != nil {
		return err
	}
	//add all from and to
	if err = client.Mail(mail.senderID); err != nil {
		return err
	}
	for _, k := range mail.toIds {
		if err = client.Rcpt(k); err != nil {
			return err
		}
	}

	//Data
	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(msgBody))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}
	client.Quit()

	return nil
}

func emailNotifyHelp(from string, to []string, subject string, msg string, SMTPHost string, SMTPPort string, pwd string) error {
	mail := newMail(from, to, subject, msg)
	smtpServer := newSMTPServer(SMTPHost, SMTPPort)
	return smtpEmail(mail, smtpServer, pwd)
}

//EmailNotify(to []string, subject, msg string, ntfs Notifiers)
//send an email with subject and message provided with parameters
//to the email address stored in(to []string)
func EmailNotify(to []string, subject, msg string, ntfs Notifiers) error {
	ntf := ntfs.SMTPEmailNotifier

	//check the notification type "smtpemail" and find if the state is "on"
	//if no type of "smtpemail" or the state is "off", do nothing
	if ntf.Type == "smtpemail" && ntf.State == "on" {
		return emailNotifyHelp(ntf.Account, to, subject, msg,
			ntf.SMTPHost, ntf.SMTPPort, ntf.Pwd)
	}

	return errors.New("Email Notification is invalid")
}
