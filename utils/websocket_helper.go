package utils

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"time"

	"Github.com/Yobubble/email-virus-scanner-server/config"
	"github.com/gorilla/websocket"
)

type message struct {
	Type string `json:"Type"`
	Data struct {
		ID string `json:"ID"`
	} `json:"Data"`
}

type websocketHelper struct {
	cfg      *config.Cfg
	done     chan struct{}
	conn     *websocket.Conn
	emailIDs chan string
}

func (w *websocketHelper) receiveMessages() {
	defer func() {
		close(w.emailIDs)
		close(w.done)
	}()
	for {
		_, msg, err := w.conn.ReadMessage()
		if err != nil {
			log.Println("Error read message:", err)
			return
		}

		var m message
		if err := json.Unmarshal(msg, &m); err != nil {
			log.Println("Error unmarshaling message:", err)
			return
		}

		if m.Type == "new" {
			w.emailIDs <- m.Data.ID
			log.Printf("recv: %s, ID: %s", m.Type, m.Data.ID)
		} else {
			log.Printf("recv: %s", m.Type)
		}
	}
}

func (w *websocketHelper) OpenWebsocket() error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	var err error
	w.conn, _, err = websocket.DefaultDialer.Dial(w.cfg.Mp.WebsocketUrl, nil)
	if err != nil {
		return err
	}
	defer w.conn.Close()

	go w.receiveMessages()

	for {
		select {
		case <-w.done:
			return nil
		case <-interrupt:
			log.Println("interrupt")
			err := w.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Error write close:", err)
				return err
			}
			select {
			case <-w.done:
			case <-time.After(time.Second):
			}
			return nil
		}
	}
}

func NewWebsocketHelper(cfg *config.Cfg, emailIDs chan string) *websocketHelper {
	return &websocketHelper{
		cfg:      cfg,
		done:     make(chan struct{}),
		emailIDs: emailIDs,
	}
}
