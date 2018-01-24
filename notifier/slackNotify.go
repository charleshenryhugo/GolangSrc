package main

import (
	"errors"

	"github.com/nlopes/slack"
)

//parse tokens from "notifyrc.xml"
func getToken(ntf SlackNotifier) string {
	if ntf.Type == "slack" && ntf.State == "on" {
		return ntf.Token
	}
	return ""
}

//get all channels using your token
func getSlackChannels(token string) (channels []slack.Channel, err error) {
	api := slack.New(token)
	channels, err = api.GetChannels(true)
	return channels, err
}

//get all users using your token
func getSlackUsers(token string) (users []slack.User, err error) {
	api := slack.New(token)
	users, err = api.GetUsers()
	return users, err
}

//build a slack attachment for slack message parameter and return it
//(to be extended......)
func buildAttachment(title, pretext, text string) slack.Attachment {
	attachment := slack.Attachment{
		Title:   title,
		Pretext: pretext,
		Text:    text,
	}

	return attachment
}

//build a slack message parameter and return it
func buildMessageParameters(attachment slack.Attachment, ntf SlackNotifier) slack.PostMessageParameters {
	params := slack.PostMessageParameters{}
	params.Attachments = []slack.Attachment{attachment}
	params.AsUser = ntf.AsUser
	params.Username = ntf.UserName
	params.IconEmoji = ":" + ntf.IconEmoji + ":"
	return params
}

//send message to channels using your token parsed from SlackNotifier
func postMsgChannels(ntf SlackNotifier, channelIDs []string, msgTitle, attachTitle, attachPretext, attachText string) ([]string, string, error) {
	token := ntf.Token
	if token == "" {
		return []string{}, "", errors.New("Invalid Token")
	}
	api := slack.New(token)
	msgAttachment := buildAttachment(attachTitle, attachPretext, attachText)
	params := buildMessageParameters(msgAttachment, ntf)

	var (
		timestamp string
		err       error
	)
	for _, channelID := range channelIDs {
		_, timestamp, err = api.PostMessage(channelID, msgTitle, params)
	}

	return channelIDs, timestamp, err
}

//send message to users using your token parsed from SlackNotifier
func postMsgUsers(ntf SlackNotifier, userIDs []string, msgTitle string, attachment slack.Attachment) ([]string, string, error) {
	return postMsgChannels(ntf, userIDs, msgTitle,
		attachment.Title, attachment.Pretext, attachment.Text)
}

//SlackNotify (to []string, subject, msg string, ntfs Notifiers)
//post a notification with subject and message provided with parameters
//to the slack userIDs(ChannelIDs) stored in(to []string)
//ChannelID and UserID are both available
func SlackNotify(to []string, subject, msg string, ntfs Notifiers) ([]string, string, error) {
	ntf := ntfs.SlackNotifier
	if ntf.Type == "slack" && ntf.State == "on" {
		attachment := slack.Attachment{Text: msg}
		return postMsgUsers(ntf, to, subject, attachment)
	}
	return []string{}, "", errors.New("Slack Notification is invalid")
}
