package entities

// Ref: https://mailpit.axllent.org/docs/api-v1/view.html#post-/api/v1/send
type SendAMessageEntity struct {
	Attachments []SendAMessageAttachment `json:"Attachments,omitempty"`
	Bcc         []string                 `json:"bcc,omitempty"`
	Cc          []SendAMessageRecipient  `json:"Cc,omitempty"`
	From        SendAMessageRecipient    `json:"From"`
	HTML        string                   `json:"HTML,omitempty"`
	Headers     map[string]string        `json:"Headers,omitempty"`
	ReplyTo     []SendAMessageRecipient  `json:"ReplyTo,omitempty"`
	Subject     string                   `json:"Subject"`
	Tags        []string                 `json:"Tags,omitempty"`
	Text        string                   `json:"Text,omitempty"`
	To          []SendAMessageRecipient  `json:"To,omitempty"`
}

type SendAMessageAttachment struct {
	Content     string `json:"Content"`
	ContentID   string `json:"ContentID,omitempty"`
	ContentType string `json:"ContentType,omitempty"`
	FileName    string `json:"FileName"`
}

type SendAMessageRecipient struct {
	Email string `json:"Email"`
	Name  string `json:"Name,omitempty"`
}
