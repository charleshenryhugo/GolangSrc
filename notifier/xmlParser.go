package main

/*
import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Notifier struct {
	Type  string `xml:"type,attr"`
	State string `xml:"state"`
}

type Notifiers struct {
	XMLName           xml.Name          `xml:"notifiers"`
	SMTPEmailNotifier SmtpEmailNotifier `xml:"smtpemailnotifier"`
	SlackNotifier     SlackNotifier     `xml:"slacknotifier"`
}

type SmtpEmailNotifier struct {
	XMLName  xml.Name `xml:"smtpemailnotifier"`
	Type     string   `xml:"type,attr"`
	State    string   `xml:"state"`
	Account  string   `xml:"account"`
	Pwd      string   `xml:"pwd"`
	SMTPHost string   `xml:"SMTPHost"`
	SMTPPort string   `sml:"SMTPPort"`
}

type SlackNotifier struct {
	XMLName   xml.Name `xml:"slacknotifier"`
	Type      string   `xml:"type,attr"`
	State     string   `xml:"state"`
	Token     string   `xml:"token"`
	AsUser    bool     `xml:"asUser"`
	UserName  string   `xml:"userName"`
	IconEmoji string   `xml:"iconEmoji"`
}

//parse notifiers objects from the *.xml file
func parseNotifiers(file string) Notifiers {
	//open xmlFile
	xmlFile, err := os.Open(file)
	if err != nil {
		log.Panic(err)
	}

	defer xmlFile.Close()

	//read xmlFile as a byte array
	byteValue, _ := ioutil.ReadAll(xmlFile)

	//initialize Notifiers
	var notifiers Notifiers
	//unmarshall byteValue into Notifiers
	xml.Unmarshal(byteValue, &notifiers)

	return notifiers
}

type Defaults struct {
	XMLName       xml.Name `xml:"defaults"`
	EmailListFile string   `xml:"emailListFile"`
	SlackListFile string   `xml:"slackListFile"`
	Subject       string   `xml:"subject"`
	Message       string   `xml:"message"`
	MessageFile   string   `xml:"messageFile"`
}

//parse the Defaults object from *.xml file
func parseDefaults(file string) Defaults {
	//open xmlFile
	xmlFile, err := os.Open(file)
	if err != nil {
		log.Panic(err)
	}

	defer xmlFile.Close()

	//read xmlFile as a byte array
	byteValue, _ := ioutil.ReadAll(xmlFile)

	//initialize Defaults object
	var defaults Defaults
	//unmarshall byteValue into Defaults object
	xml.Unmarshal(byteValue, &defaults)

	return defaults
}

func (dflt *Defaults) GetDfltSbjt() string {
	return dflt.Subject
}
func (dflt *Defaults) GetDfltmsg() string {
	//get message from the default message file, only if the file is available
	if fileBytes, err := ioutil.ReadFile(dflt.MessageFile); err == nil { //only when the msg file is read successfully, we rewrite the msg
		return string(fileBytes)
	}
	//if file reading is failed, then return default message directly
	return dflt.Message
}

func (dflt *Defaults) GetDfltSlackList() []string {
	//return those slack user IDs stored in the default file, only if the file is available
	if fileBytes, err := ioutil.ReadFile(dflt.SlackListFile); err == nil {
		return strings.Fields(string(fileBytes))
	}
	return []string{}
}
func (dflt *Defaults) GetDfltEmailList() []string {
	//return those email addrs stored in the default file, only if the file is available
	if fileBytes, err := ioutil.ReadFile(dflt.EmailListFile); err == nil {
		return strings.Fields(string(fileBytes))
	}
	return []string{}
}
*/
