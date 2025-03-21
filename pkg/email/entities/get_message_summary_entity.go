package entities

import (
	"time"
)

// Ref: https://mailpit.axllent.org/docs/api-v1/view.html#get-/api/v1/message/-ID-
type GetMessageSummaryEntity struct {
	Attachments []GetMessageSummaryAttachment `json:"Attachments,omitempty"`
	Bcc         []GetMessageSummaryRecipient  `json:"Bcc,omitempty"`
	Cc          []GetMessageSummaryRecipient  `json:"Cc,omitempty"`
	Date        time.Time                     `json:"Date,omitempty"`
	From        GetMessageSummaryRecipient    `json:"From,omitempty"`
	HTML        string                        `json:"HTML,omitempty"`
	ID          string                        `json:"ID,omitempty"`
	Inline      []GetMessageSummaryAttachment `json:"Inline,omitempty"`
	MessageID   string                        `json:"MessageID,omitempty"`
	ReplyTo     []GetMessageSummaryRecipient  `json:"ReplyTo,omitempty"`
	ReturnPath  string                        `json:"ReturnPath,omitempty"`
	Size        int                           `json:"Size,omitempty"`
	Subject     string                        `json:"Subject,omitempty"`
	Tags        []string                      `json:"Tags,omitempty"`
	Text        string                        `json:"Text,omitempty"`
	To          []GetMessageSummaryRecipient  `json:"To,omitempty"`
}

type GetMessageSummaryAttachment struct {
	ContentID   string `json:"ContentID,omitempty"`
	ContentType string `json:"ContentType,omitempty"`
	FileName    string `json:"FileName,omitempty"`
	PartID      string `json:"PartID,omitempty"`
	Size        int    `json:"Size,omitempty"`
}

type GetMessageSummaryRecipient struct {
	Address string `json:"Address,omitempty"`
	Name    string `json:"Name,omitempty"`
}
