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
			{
				Content:  utils.Base64Encode(utils.GetFile("./assets/mock_files/doc.pdf")),
				FileName: "doc.pdf",
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
	VirusEmail = entities.SendAMessageEntity{
		Attachments: []entities.SendAMessageAttachment{
			{
				Content:  utils.Base64Encode([]byte("X5O!P%@AP[4\\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*")),
				FileName: "eicar_test_file.txt",
			},
			{
				Content:  utils.Base64Encode([]byte("!!VIRUS_SIGNATURE_EXAMPLE_1!!")),
				FileName: "virus_sample_1.txt",
			},
			{
				Content:  utils.Base64Encode([]byte("恶意软件")),
				FileName: "non_ascii_virus.txt",
			},
		},
		From: entities.SendAMessageRecipient{
			Email: "sender3@example.com",
		},
		To: []entities.SendAMessageRecipient{
			{
				Email: "receiver3@example.com",
			},
		},
		Subject: "Email with Virus",
		Text:    "This email contains attachments with known virus signatures for testing purposes.",
	}
)
