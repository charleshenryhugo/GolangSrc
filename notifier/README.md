## Name
Notifier

## Overview
Notifier is a simple command line tool written in GO and can be used to send notifications through email and slack.

## Description
Notifier is a command line tool, which can send email and slack notifications.(More notification methods to be added)

Users will have to obtain a valid email account and slack token before using Notifier. 

Users can only send notifications to their own slack group specified by the slack token. 

Get slack token from: 

https://api.slack.com/custom-integrations/legacy-tokens 

## Requirements
- darwin (UNIX-like, Mach, BSD)

## Install
download the following files and put them in the same directory.

Necessary:
- Notifier
- defaults.yml
- notifyrc.yml

Optional: (you can ignore these optional files)
- error.log
- slackListFile
- emailListFile

Easier installing methods(eg. npm, brew) will be added.

## Usage
### command line option
Just type `./notifier --help` , out comes the usage of all command line options:
```
Usage of `./notifier`:
  -e string
    	Specify the target emails. Do nothing if the email state is off
  -el string
    	Specify the file that stores target email list. Do nothing if the email state is off
  -f string
    	Specify the file that stores your notification-message
  -k string
    	Specify the target slack userIDs. Do nothing if the slack state is off
  -kl string
    	Specify the file that stores target slack userID list. Do nothing if the email state is off
  -m string
    	Specify the message of your notification
  -s string
    	Specify the title/subject of your notification
```
### files
- defaults.yml

This file is used for configuring default settings such as default notification message and subject.

(Find more details in the file itself.)

- notifyrc.yml

This file is used for configuring the notification methods such as slack token and email account

(Find more details in the file itself.)

### some examples

e.g.1
```
./notifier -s AnalysisFailed! -m Some\ error\ happened! -e google@gmail.com\ yahoo@gmail.com -kl slackListFile
```
This will send a notification with subject:"AnalysisFailed!" and message:"Some error happened!" 
to google@gmail.com and yahoo@gmail.com as well as slack users(or channels) that have IDs stored in "slackListFile",
which looks like this:

slackListFile
```
U7BL3HC86
U7BL3IC87
U7BL3IC88
U7BL3IC89
U7BL3IC90
```

e.g.2
```
./notifier -el emailListFile -k U7BL3HC86
```
This will send a notification to slack ID U7BL3HC86 and the email accounts stored in emailListFile 
which looks like this:
emailListFile
```
google@gmail.com
yahoo@gmail.com
```
In addition, subject and message will be taken from `defaults.yml` because no message or subject option is specified.


e.g.3
```
./notifier
```
There are no command line options specified, so all the settings will be taken from `defaults.yml`
So the trick is, write all necessary default settings in advance and things become easy. 

That is:
- create a file(e.g emailListFile) and write all the target email accounts. 
- create a file(e.g slackListFile) and write all the target slack IDs(they must be in your group). 
- create a file(e.g error.log) and write the default message you want to send in the next minute or in the future.
- configure the files you have just created (or downloaded) in `defaults.yml`.
- do some other default settings(please refer to `defaults.yml`)


## Demo

## Author
ZHU YUE
