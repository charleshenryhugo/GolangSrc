package main

//limitation parameters
const (
	//maximum subject length for email(Bytes)
	MAX_EMAIL_SUBJECT_LEN = 256
)

//config files
const (
	notifyrcFile string = "notifyrc"
	defaultsFile string = "defaults"
)

//Notifiers name
const (
	emailNotifier string = "smtpemailnotifier"
	slackNotifier string = "slackNotifier"
)

//ERR refers to error code(0~255), equals to uint8
type ERR uint8

//common error codes
const (
	ERR_MAX = 255

	//UNIX/LINUX general error code
	NIL                 ERR = 0
	SUCCESS             ERR = 0
	GENERAL_ERR         ERR = 1
	MISS_USE            ERR = 2
	CMD_CANNOT_EXE      ERR = 126
	CMD_NOT_FOUND       ERR = 127
	INVALID_ARG_TO_EXIT ERR = 128
	CTRLC_TERMINATE     ERR = 130

	//file-reading error code
	NOTIFRC_PARSE_ERR ERR = 55 //eror occurs while parsing notifier config file
	DFLTS_PARSE_ERR   ERR = 56 //error occurs while pasing default config file

	//smtpemail error code
	SMTPM_INVAL         ERR = 11 //smtp notif not valid
	SMTPM_SVR_CONN_ERR  ERR = 12 //error occurs while dialing to smtpserver, stop and check network, host and port ..to be added..
	SMTPM_CLT_BLD_ERR   ERR = 13 //error occurs while building a client, stop and check network, host and port
	SMTPM_AUTH_ERR      ERR = 14 //error occurs while authenticating mail account and password, stop and check your account and network
	SMTPM_SENDER_ERR    ERR = 15 //error occurs while adding sender account, stop and check your account
	SMTPM_RCVR_ERR      ERR = 16 //error occurs while adding receivers' accounts, stop and check those account IDs
	SMTPM_CLT_IO_ERR    ERR = 17 //... while initializing or close a iowriter for email client
	SMTPM_CLT_DATA_ERR  ERR = 18 //... while trying to write message in the client
	SMTPM_CLT_CLOSE_ERR ERR = 19 //... while closing an smtpemail client

	//slack error code
	SLK_INVAL       ERR = 29 //slack notif not valid
	SLK_TOKEN_INVAL ERR = 30 //slack token not invalid
	SLK_POST_ERR    ERR = 31 //token is right, just got stuck in posting to one target user(or channel)
)

//
