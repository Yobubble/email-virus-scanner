package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"Github.com/Yobubble/email-virus-scanner-server/config"
)

// Ref: https://mailpit.axllent.org/docs/api-v1/view.html#post-/api/v1/send
type Email struct {
	Attachments []attachment      `json:"Attachments,omitempty"`
	Bcc         []string          `json:"bcc,omitempty"`
	Cc          []recipient       `json:"Cc,omitempty"`
	From        recipient         `json:"From"`
	HTML        string            `json:"HTML,omitempty"`
	Headers     map[string]string `json:"Headers,omitempty"`
	ReplyTo     []recipient       `json:"ReplyTo,omitempty"`
	Subject     string            `json:"Subject"`
	Tags        []string          `json:"Tags,omitempty"`
	Text        string            `json:"Text,omitempty"`
	To          []recipient       `json:"To,omitempty"`
}

type attachment struct {
	Content     string `json:"Content"`
	ContentID   string `json:"ContentID,omitempty"`
	ContentType string `json:"ContentType,omitempty"`
	Filename    string `json:"Filename"`
}

type recipient struct {
	Email string `json:"Email"`
	Name  string `json:"Name,omitempty"`
}

type emailHelper struct {
	cfg *config.Cfg
}

func (e *emailHelper) SendEmail(mail Email) error {
	jsonBytes, err := json.Marshal(mail)
	if err != nil {
		return err
	}

	res, err := http.Post(e.cfg.Mp.ApiUrl+"/send", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	log.Printf("Succesfully send an email: %s", body)
	return nil
}

func (e *emailHelper) GetEmailFromID(ID string) {
	// TODO
}

func NewEmailHelper(cfg *config.Cfg) *emailHelper {
	return &emailHelper{
		cfg: cfg,
	}
}
