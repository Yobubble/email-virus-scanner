package mailpitWebsocket

import "Github.com/Yobubble/email-virus-scanner/utils"

type mailpitWebsocketControlller struct {
	mu *mailpitWebsocketUseCases
}

func (m *mailpitWebsocketControlller) EstablishMailpitWebsocket(emailIDs chan string) {
	err := m.mu.OpenMailpitWebsocketClient(emailIDs)
	if err != nil {
		utils.Sugar.Panic(err)
	}
}

func NewMailpitWebsocketController(mu *mailpitWebsocketUseCases) *mailpitWebsocketControlller {
	return &mailpitWebsocketControlller{
		mu: mu,
	}
}
