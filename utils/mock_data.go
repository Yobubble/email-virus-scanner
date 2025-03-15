package utils

var (
	BasicMail = Email{
		From: recipient{
			Email: "sender1@example.com",
		},
		To: []recipient{
			{
				Email: "receiver1@example.com",
			},
		},
		Subject: "Basic Email",
		Text:    "no attachment only email's body.",
	}
	AttachmentMail = Email{
		Attachments: []attachment{
			{
				Content:  Base64Encode(GetFile("./assets/mock_files/plain_text.txt")),
				Filename: "plain_text.txt",
			},
			{
				Content:  Base64Encode(GetFile("./assets/mock_files/image.jpg")),
				Filename: "image.jpg",
			},
		},
		From: recipient{
			Email: "sender2@example.com",
		},
		To: []recipient{
			{
				Email: "receiver2@example.com",
			},
		},
		Subject: "Email with Attachments",
		Text:    "Email with attachments and body text.",
	}
	// More mock email possible
)
