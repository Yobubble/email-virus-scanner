package email

import (
	"Github.com/Yobubble/email-virus-scanner/pkg/email/entities"
	EmailUtils "Github.com/Yobubble/email-virus-scanner/pkg/email/utils"
	"Github.com/Yobubble/email-virus-scanner/utils"
)

type emailController struct {
	eu *emailUseCases
}

func (e *emailController) SendMockEmail() {
	err := e.eu.SendAMessage(EmailUtils.AttachmentMail)
	if err != nil {
		utils.Sugar.Panic(err)
	}
}

// goroutine
func (e *emailController) ReceiveEmailIDAndConvertToEmail(emailIDs chan string, EmailBodies chan entities.GetMessageSummaryEntity) {
	for emailID := range emailIDs {
		utils.Sugar.Infof("Getting email from the id: %s ...", emailID)

		emailBody, err := e.eu.GetMessageSummary(emailID)
		if err != nil {
			utils.Sugar.Panic(err)
		}

		utils.Sugar.Debugf("Email's body: %v", emailBody)

		EmailBodies <- emailBody
	}
}

func NewEmailController(eu *emailUseCases) *emailController {
	return &emailController{
		eu: eu,
	}
}
