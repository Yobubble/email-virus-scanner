package virusScanner

import "Github.com/Yobubble/email-virus-scanner/config"

type virusScannerUseCase struct {
	cfg *config.Cfg
}

// TODO

func NewVirusScannerUseCase(cfg *config.Cfg) *virusScannerUseCase {
	return &virusScannerUseCase{
		cfg: cfg,
	}
}
