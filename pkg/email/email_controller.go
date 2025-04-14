package email

import (
	"Github.com/Yobubble/email-virus-scanner/pkg/email/entities"
	EmailUtils "Github.com/Yobubble/email-virus-scanner/pkg/email/utils"
	"Github.com/Yobubble/email-virus-scanner/utils"
)

type emailController struct {
	eu *EmailUseCases
}

func (e *emailController) SendAttachmentEmail() {
	err := e.eu.SendAMessage(EmailUtils.AttachmentMail)
	if err != nil {
		utils.Sugar.Panic(err)
	}
}

func (e *emailController) SendVirusEmail() {
	err := e.eu.SendAMessage(EmailUtils.VirusEmail)
	if err != nil {
		utils.Sugar.Panic(err)
	}
}

// goroutine
func (e *emailController) ReceiveEmailIDAndConvertToEmail(emailIDs chan string, EmailBodies chan entities.GetMessageSummaryEntity) {
	for emailID := range emailIDs {
		emailBody, err := e.eu.GetMessageSummary(emailID)
		if err != nil {
			utils.Sugar.Panic(err)
		}
		EmailBodies <- emailBody
	}
}

func NewEmailController(eu *EmailUseCases) *emailController {
	return &emailController{
		eu: eu,
	}
}
