package virusScanner

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"Github.com/Yobubble/email-virus-scanner/config"
	"Github.com/Yobubble/email-virus-scanner/utils"
)

type virusScannerUseCase struct {
	cfg             *config.Cfg
	virusSignatures [][]byte
}

func NewVirusScannerUseCase(cfg *config.Cfg) *virusScannerUseCase {
	signatures := [][]byte{
		[]byte("X5O!P%@AP[4\\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*"), // EICAR test string
		[]byte("!!VIRUS_SIGNATURE_EXAMPLE_1!!"),
		[]byte("恶意软件"),
	}

	return &virusScannerUseCase{
		cfg:             cfg,
		virusSignatures: signatures,
	}
}

func (v *virusScannerUseCase) ScanAttachment(content []byte) (bool, string) {
	utils.Sugar.Debugf("Scanning content of size: %d bytes", len(content))
	for _, signature := range v.virusSignatures {

		if bytes.Contains(content, signature) {
			return true, string(signature)
		}
	}
	return false, ""
}

func (e *virusScannerUseCase) GetAttachmentContent(emailID string, partID string) ([]byte, error) {
	url := e.cfg.Mp.ApiUrl + "/message/" + emailID + "/part/" + partID

	res, err := http.Get(url)
	if err != nil {
		utils.Sugar.Errorf("Call API error for attachment content: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(res.Body) // Read body for error details
		utils.Sugar.Errorf("Failed to fetch attachment content, status code: %d, response: %s", res.StatusCode, string(bodyBytes))
		return nil, fmt.Errorf("failed to fetch attachment content, status: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		utils.Sugar.Errorf("Read attachment body error: %v", err)
		return nil, err
	}

	// IMPORTANT: Check Mailpit API documentation. The content might be base64 encoded
	// within a JSON structure, or it might be raw binary data.
	// If it's JSON with a base64 string, you'll need to unmarshal and decode.
	// If it's raw binary, you can return `body` directly.
	// Assuming raw binary for simplicity here:
	// utils.Sugar.Debugf("Successfully fetched attachment content, size: %d bytes", len(body))
	return body, nil

	// --- Example if content is base64 within JSON ---
	// var attachmentData struct {
	// 	Content string `json:"Content"` // Assuming a JSON field named 'Content' holds base64 data
	// }
	// if err := json.Unmarshal(body, &attachmentData); err != nil {
	// 	utils.Sugar.Errorf("Failed to unmarshal attachment JSON: %v", err)
	// 	return nil, err
	// }
	// decodedContent, err := base64.StdEncoding.DecodeString(attachmentData.Content)
	// if err != nil {
	// 	utils.Sugar.Errorf("Failed to decode base64 attachment content: %v", err)
	// 	return nil, err
	// }
	// return decodedContent, nil
	// --- End Example ---
}
