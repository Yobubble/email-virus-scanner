package virusScanner

import (
	"bytes"
	"encoding/base64"

	"Github.com/Yobubble/email-virus-scanner/config"
	"Github.com/Yobubble/email-virus-scanner/utils"
)

type virusScannerUseCase struct {
	cfg            *config.Cfg
	// Simple list of "virus signatures" (e.g., specific strings or byte patterns)
	// In a real-world scenario, use a proper virus definition database (e.g., ClamAV)
	virusSignatures [][]byte
}

// Initialize the use case with predefined virus signatures
func NewVirusScannerUseCase(cfg *config.Cfg) *virusScannerUseCase {
	// Example signatures (replace with actual signatures or database)
	// These are just placeholder strings for demonstration.
	signatures := [][]byte{
		[]byte("X5O!P%@AP[4\\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*"), // EICAR test string
		[]byte("!!VIRUS_SIGNATURE_EXAMPLE_1!!"),
		[]byte("恶意软件"), // Example non-ASCII signature
	}
	return &virusScannerUseCase{
		cfg:            cfg,
		virusSignatures: signatures,
	}
}

// ScanAttachment checks if the provided content contains any known virus signatures.
// Content is expected to be raw bytes after base64 decoding if necessary.
func (v *virusScannerUseCase) ScanAttachment(content []byte) (bool, string) {
	utils.Sugar.Debugf("Scanning content of size: %d bytes", len(content))
	for _, signature := range v.virusSignatures {
		if bytes.Contains(content, signature) {
			utils.Sugar.Warnf("Virus signature detected: %s", string(signature))
			// Return true (virus found) and the signature that matched
			return true, string(signature)
		}
	}
	utils.Sugar.Debugf("No virus signatures found in content.")
	// Return false (no virus found)
	return false, ""
}

// Helper function (potentially): Decode attachment content if it's base64 encoded.
// Depending on how you fetch the content, it might already be decoded.
func (v *virusScannerUseCase) decodeContent(encodedContent string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(encodedContent)
	if err != nil {
		utils.Sugar.Errorf("Failed to decode base64 content: %v", err)
		return nil, err
	}
	return decoded, nil
}