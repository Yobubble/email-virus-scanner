package utils

import (
	"Github.com/Yobubble/email-virus-scanner/pkg/email/entities"
	"Github.com/Yobubble/email-virus-scanner/utils"
)

var (
	BasicMail = entities.SendAMessageEntity{
		From: entities.SendAMessageRecipient{
			Email: "sender1@example.com",
		},
		To: []entities.SendAMessageRecipient{
			{
				Email: "receiver1@example.com",
			},
		},
		Subject: "Basic Email",
		Text:    "no attachment only email's body.",
	}
	AttachmentMail = entities.SendAMessageEntity{
		Attachments: []entities.SendAMessageAttachment{
			{
				Content:  utils.Base64Encode(utils.GetFile("./assets/mock_files/plain_text.txt")),
				FileName: "plain_text.txt",
			},
			{
				Content:  utils.Base64Encode(utils.GetFile("./assets/mock_files/image.jpg")),
				FileName: "image.jpg",
			},
		},
		From: entities.SendAMessageRecipient{
			Email: "sender2@example.com",
		},
		To: []entities.SendAMessageRecipient{
			{
				Email: "receiver2@example.com",
			},
		},
		Subject: "Email with Attachments",
		Text:    "Email with attachments and body text.",
	}
)
