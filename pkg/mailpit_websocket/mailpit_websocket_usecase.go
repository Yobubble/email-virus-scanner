package mailpitWebsocket

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"time"

	"Github.com/Yobubble/email-virus-scanner/config"
	"Github.com/Yobubble/email-virus-scanner/pkg/mailpit_websocket/entities"
	"Github.com/Yobubble/email-virus-scanner/utils"
	"github.com/gorilla/websocket"
)

type mailpitWebsocketUseCases struct {
	cfg  *config.Cfg
	done chan struct{}
	conn *websocket.Conn
}

func (m *mailpitWebsocketUseCases) receiveMessages(emailIDs chan string, errCh chan<- error) {
	defer func() {
		close(emailIDs)
		close(m.done)
	}()
	for {
		_, msg, err := m.conn.ReadMessage()
		if err != nil {
			utils.Sugar.Errorf("Error reading incoming message")
			errCh <- err
		}

		var m entities.MailpitWebsocketEntity
		if err := json.Unmarshal(msg, &m); err != nil {
			utils.Sugar.Errorf("Error unmarshaling message")
			errCh <- err
		}

		if m.Type == "new" {
			emailIDs <- m.Data.ID
			utils.Sugar.Infof("Receiving email's ID: %s", m.Data.ID)
		} else {
			utils.Sugar.Infof("Others: %s", msg)
		}
	}
}

func (m *mailpitWebsocketUseCases) OpenMailpitWebsocketClient(emailIDs chan string) error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	errCh := make(chan error, 1)
	var err error
	m.conn, _, err = websocket.DefaultDialer.Dial(m.cfg.Mp.WebsocketUrl, nil)
	if err != nil {
		return err
	}
	defer m.conn.Close()

	go m.receiveMessages(emailIDs, errCh)

	for {
		select {
		case <-errCh:
			return err
		case <-m.done:
			return nil
		case <-interrupt:
			log.Println("interrupt")
			err := m.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Error write close:", err)
				return err
			}
			select {
			case <-m.done:
			case <-time.After(time.Second):
			}
			return nil
		}
	}
}

func NewMailpitWebsocketUseCases(cfg *config.Cfg) *mailpitWebsocketUseCases {
	return &mailpitWebsocketUseCases{
		cfg:  cfg,
		done: make(chan struct{}),
	}
}
