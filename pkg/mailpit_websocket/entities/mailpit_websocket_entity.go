package entities

type MailpitWebsocketEntity struct {
	Type string `json:"Type"`
	Data struct {
		ID string `json:"ID"`
	} `json:"Data"`
}
