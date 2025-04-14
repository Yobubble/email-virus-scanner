package email

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"Github.com/Yobubble/email-virus-scanner/config"
	"Github.com/Yobubble/email-virus-scanner/pkg/email/entities"
	"Github.com/Yobubble/email-virus-scanner/utils"
)

type EmailUseCases struct {
	cfg *config.Cfg
}

func (e *EmailUseCases) SendAMessage(mail entities.SendAMessageEntity) error {
	jsonBytes, err := json.Marshal(mail)
	if err != nil {
		utils.Sugar.Error("Convert to json error")
		return err
	}

	res, err := http.Post(e.cfg.Mp.ApiUrl+"/send", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		utils.Sugar.Error("Call api error")
		return err
	}
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	if err != nil {
		utils.Sugar.Error("Read body error")
		return err
	}

	utils.Sugar.Debugf("Successfully send a message with subject: %s", mail.Subject)
	return nil
}

func (e *EmailUseCases) GetMessageSummary(ID string) (entities.GetMessageSummaryEntity, error) {
	var email entities.GetMessageSummaryEntity

	res, err := http.Get(e.cfg.Mp.ApiUrl + "/message/" + ID)
	if err != nil {
		utils.Sugar.Error("Cal api error")
		return entities.GetMessageSummaryEntity{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		utils.Sugar.Error("Read body error")
		return entities.GetMessageSummaryEntity{}, err
	}

	err = json.Unmarshal(body, &email)
	if err != nil {
		utils.Sugar.Error("Convert body to object error")
		return entities.GetMessageSummaryEntity{}, err
	}

	return email, nil
}

func (e *EmailUseCases) GetAttachmentContent(emailID string, partID string) ([]byte, error) {
	url := e.cfg.Mp.ApiUrl + "/message/" + emailID + "/part/" + partID

	res, err := http.Get(url)
	if err != nil {
		utils.Sugar.Errorf("Call API error for attachment content: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(res.Body)
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
	return body, nil

	// --- Example if content is base64 within JSON ---
	// var attachmentData struct {
	// 	Content string `json:"Content"` // Assuming a JSON field named 'Content' holds base64 data
	// }
	// if err := json.Unmarshal(body, &attachmentData); err != nil {
	// 	utils.Sugar.Errorf("Failed to unmarshal attachment JSON: %v", err)
	//	return nil, err
	// }
	// decodedContent, err := base64.StdEncoding.DecodeString(attachmentData.Content)
	// if err != nil {
	//	 utils.Sugar.Errorf("Failed to decode base64 attachment content: %v", err)
	//	 return nil, err
	// }
	// return decodedContent, nil
	// --- End Example ---
}

func NewEmailUseCases(cfg *config.Cfg) *EmailUseCases {
	return &EmailUseCases{
		cfg: cfg,
	}
}
