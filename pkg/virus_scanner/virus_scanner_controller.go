package virusScanner

import (
	"Github.com/Yobubble/email-virus-scanner/pkg/email/entities"
	"Github.com/Yobubble/email-virus-scanner/utils"
)

type virusScannerController struct {
	vsu *virusScannerUseCase
}

// goroutine
func (v *virusScannerController) EmailScanning(emailBodies chan entities.GetMessageSummaryEntity) {
	for emailBody := range emailBodies {
		utils.Sugar.Debugf("Scanning...\n %v", emailBody)
		// TODO
	}
}

func NewVirusScannerController(vsu *virusScannerUseCase) *virusScannerController {
	return &virusScannerController{
		vsu: vsu,
	}
}
