package virusScanner

import (
	"Github.com/Yobubble/email-virus-scanner/pkg/email"
	"Github.com/Yobubble/email-virus-scanner/pkg/email/entities"
	"Github.com/Yobubble/email-virus-scanner/utils"
)

type virusScannerController struct {
	vsu *virusScannerUseCase
	eu  *email.EmailUseCases
}

func NewVirusScannerController(vsu *virusScannerUseCase, eu *email.EmailUseCases) *virusScannerController {
	return &virusScannerController{
		vsu: vsu,
		eu:  eu,
	}
}

// goroutine
func (v *virusScannerController) EmailScanning(emailBodies chan entities.GetMessageSummaryEntity) {
	for emailBody := range emailBodies {
		utils.Sugar.Infof("Scanning: %s", emailBody.Subject)

		if len(emailBody.Attachments) == 0 {
			utils.Sugar.Warnf("Email ID: %s has no attachments to scan.", emailBody.ID)
			continue
		}

		for _, attachment := range emailBody.Attachments {
			if v.eu == nil {
				utils.Sugar.Error("EmailUseCase is not initialized in VirusScannerController. Cannot fetch attachment content.")
				continue
			}

			attachmentContent, err := v.eu.GetAttachmentContent(emailBody.ID, attachment.PartID)
			if err != nil {
				utils.Sugar.Errorf("Failed to fetch content for attachment %s (PartID: %s): %v", attachment.FileName, attachment.PartID, err)
				continue
			}

			isVirus, detectedSignature := v.vsu.ScanAttachment(attachmentContent) // Use decodedContent if decoding was needed

			if isVirus {
				utils.Sugar.Warnf("VIRUS DETECTED in attachment '%s' (Email Subject: %s)! Signature: %s",
					attachment.FileName, emailBody.Subject, detectedSignature)
			} else {
				utils.Sugar.Infof("Attachment '%s' (Email Subject: %s) scanned clean.", attachment.FileName, emailBody.Subject)
			}
		}
	}
}
