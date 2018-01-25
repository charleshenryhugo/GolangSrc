package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type Notifier struct {
	Type  string `yaml:"type"`
	State string `yaml:"state"`
}

type NotifConfig struct {
	Notifiers Notifiers `yaml:"notifiers"`
}

type Notifiers struct {
	SMTPEmailNotifier SmtpEmailNotifier `yaml:"smtpemailnotifier"`
	SlackNotifier     SlackNotifier     `yaml:"slacknotifier"`
}

type SmtpEmailNotifier struct {
	Type     string `yaml:"type"`
	State    string `yaml:"state"`
	Account  string `yaml:"account"`
	Pwd      string `yaml:"pwd"`
	SMTPHost string `yaml:"SMTPHost"`
	SMTPPort string `yaml:"SMTPPort"`
}

type SlackNotifier struct {
	Type      string `yaml:"type"`
	State     string `yaml:"state"`
	Token     string `yaml:"token"`
	AsUser    bool   `yaml:"asUser"`
	UserName  string `yaml:"userName"`
	IconEmoji string `yaml:"iconEmoji"`
}

//parse notifiers objects from the *.yaml file
func parseNotifiers(file string) Notifiers {
	//open yamlFile
	yamlFile, err := os.Open(file)
	if err != nil {
		log.Panic(err)
	}

	defer yamlFile.Close()

	//read yamlFile as a byte array
	byteValue, _ := ioutil.ReadAll(yamlFile)

	//initialize Notifiers
	var notifCfg NotifConfig
	//unmarshall byteValue into Notifiers
	yaml.Unmarshal(byteValue, &notifCfg)

	return notifCfg.Notifiers
}

type DfltConfig struct {
	Defaults Defaults `yaml:"defaults"`
}

type Defaults struct {
	EmailListFile string `yaml:"emailListFile"`
	SlackListFile string `yaml:"slackListFile"`
	Subject       string `yaml:"subject"`
	Message       string `yaml:"message"`
	MessageFile   string `yaml:"messageFile"`
}

//parse the Defaults object from *.yaml file
func parseDefaults(file string) Defaults {
	//open yamlFile
	yamlFile, err := os.Open(file)
	if err != nil {
		log.Panic(err)
	}

	defer yamlFile.Close()

	//read yamlFile as a byte array
	byteValue, _ := ioutil.ReadAll(yamlFile)

	//initialize Defaults object
	var dfltCfg DfltConfig
	//unmarshall byteValue into Defaults object
	yaml.Unmarshal(byteValue, &dfltCfg)

	return dfltCfg.Defaults
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
