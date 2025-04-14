package virusScanner

import (
	"Github.com/Yobubble/email-virus-scanner/pkg/email" // Assuming email usecases are accessible
	"Github.com/Yobubble/email-virus-scanner/pkg/email/entities"
	"Github.com/Yobubble/email-virus-scanner/utils"
)

type virusScannerController struct {
	vsu *virusScannerUseCase
	// Add email use case to fetch attachment content
	// You'll need to pass this in when creating the controller in main.go
	emailUseCase *email.EmailUseCases // Or the interface type if you define one
}

// Modify NewVirusScannerController to accept the email use case
func NewVirusScannerController(vsu *virusScannerUseCase, emailUseCase *email.EmailUseCases) *virusScannerController {
	return &virusScannerController{
		vsu:          vsu,
		emailUseCase: emailUseCase, // Store the email use case
	}
}

// goroutine
func (v *virusScannerController) EmailScanning(emailBodies chan entities.GetMessageSummaryEntity) {
	for emailBody := range emailBodies {
		utils.Sugar.Infof("Scanning email ID: %s, Subject: %s", emailBody.ID, emailBody.Subject)

		if len(emailBody.Attachments) == 0 {
			utils.Sugar.Infof("Email ID: %s has no attachments to scan.", emailBody.ID)
			continue // Skip to the next email if there are no attachments
		}

		// Iterate through attachments
		for _, attachment := range emailBody.Attachments {
			utils.Sugar.Infof("Scanning attachment: %s (PartID: %s, Size: %d)", attachment.FileName, attachment.PartID, attachment.Size)

			// --- TODO: Fetch Attachment Content ---
			// This part requires a new function in email_usecases.go
			// Example: GetAttachmentContent(emailID string, partID string) ([]byte, error)
			// This function would call the Mailpit API to get the raw attachment data.

			// Placeholder check - ensure emailUseCase is initialized
			if v.emailUseCase == nil {
				utils.Sugar.Error("EmailUseCase is not initialized in VirusScannerController. Cannot fetch attachment content.")
				continue // Skip this attachment if fetching isn't possible
			}

			// Fetch the actual content (assuming GetAttachmentContent exists and returns raw bytes)
			// Note: Mailpit API might return base64 encoded content, requiring decoding.
			// Adjust GetAttachmentContent accordingly or decode here.
			attachmentContent, err := v.emailUseCase.GetAttachmentContent(emailBody.ID, attachment.PartID)
			if err != nil {
				utils.Sugar.Errorf("Failed to fetch content for attachment %s (PartID: %s): %v", attachment.FileName, attachment.PartID, err)
				continue // Skip this attachment if content fetching fails
			}

			// If content is base64 encoded (adjust based on GetAttachmentContent implementation)
			// decodedContent, err := v.vsu.decodeContent(string(attachmentContent)) // Example if content is string
			// if err != nil {
			// 	utils.Sugar.Errorf("Failed to decode content for attachment %s: %v", attachment.FileName, err)
			// 	continue
			// }
			// --- End Fetch Attachment Content ---


			// --- TODO: Scan the fetched content ---
			// Pass the raw attachment bytes (decoded if necessary) to the scanner
			isVirus, detectedSignature := v.vsu.ScanAttachment(attachmentContent) // Use decodedContent if decoding was needed

			// --- Handle Scan Result ---
			if isVirus {
				// Take action: log, quarantine, notify, etc.
				utils.Sugar.Warnf("VIRUS DETECTED in attachment '%s' (Email ID: %s)! Signature: %s",
					attachment.FileName, emailBody.ID, detectedSignature)
				// Example action: Add a tag to the email via Mailpit API (requires another API call)
				// Example action: Send a notification
			} else {
				utils.Sugar.Infof("Attachment '%s' (Email ID: %s) scanned clean.", attachment.FileName, emailBody.ID)
			}
			// --- End Handle Scan Result ---
		}
	}
	utils.Sugar.Info("EmailScanning goroutine finished.")
}