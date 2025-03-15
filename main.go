package main

import (
	"os"

	"Github.com/Yobubble/email-virus-scanner-server/config"
	"Github.com/Yobubble/email-virus-scanner-server/utils"
)

var cfg *config.Cfg

func init() {
	cfg = config.InitConfig()
}

func main() {
	emailIDs := make(chan string, 10)

	emailHelper := utils.NewEmailHelper(cfg)
	websocketHelper := utils.NewWebsocketHelper(cfg, &emailIDs)

	args := os.Args
	switch args[1] {
	case "sendmail":

		// err := emailHelper.SendEmail(utils.BasicMail)
		// if err != nil {
		// 	panic(err)
		// }

		// NOTE: uncomment to send an email with attachments
		err := emailHelper.SendEmail(utils.AttachmentMail)
		if err != nil {
			panic(err)
		}

	case "scanmail":
		err := websocketHelper.OpenWebsocket()
		if err != nil {
			panic(err)
		}
	}
}
