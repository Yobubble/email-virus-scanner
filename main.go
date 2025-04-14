package main

import (
	"log"
	"os"

	"Github.com/Yobubble/email-virus-scanner/config"
	"Github.com/Yobubble/email-virus-scanner/pkg/email"
	"Github.com/Yobubble/email-virus-scanner/pkg/email/entities"
	mailpitWebsocket "Github.com/Yobubble/email-virus-scanner/pkg/mailpit_websocket"
	virusScanner "Github.com/Yobubble/email-virus-scanner/pkg/virus_scanner"
	"Github.com/Yobubble/email-virus-scanner/utils"
)

var cfg *config.Cfg

func init() {
	cfg = config.InitConfig()
	utils.InitLogger()
}

func main() {
	emailIDs := make(chan string, 10)
	emailBodies := make(chan entities.GetMessageSummaryEntity, 10)

	eu := email.NewEmailUseCases(cfg)
	ec := email.NewEmailController(eu)

	mwu := mailpitWebsocket.NewMailpitWebsocketUseCases(cfg)
	mwc := mailpitWebsocket.NewMailpitWebsocketController(mwu)

	vsu := virusScanner.NewVirusScannerUseCase(cfg)
	vsc := virusScanner.NewVirusScannerController(vsu, eu)

	args := os.Args
	switch args[1] {
	case "sendmail":
		ec.SendAttachmentEmail()
		ec.SendVirusEmail()

	case "scanmail":
		go ec.ReceiveEmailIDAndConvertToEmail(emailIDs, emailBodies)
		go vsc.EmailScanning(emailBodies)
		mwc.EstablishMailpitWebsocket(emailIDs)

	default:
		log.Panic("Need at least 1 argument...")
	}
}
