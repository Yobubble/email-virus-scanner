package email

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"Github.com/Yobubble/email-virus-scanner/config"
	"Github.com/Yobubble/email-virus-scanner/pkg/email/entities"
	"Github.com/Yobubble/email-virus-scanner/utils"
)

type emailUseCases struct {
	cfg *config.Cfg
}

func (e *emailUseCases) SendAMessage(mail entities.SendAMessageEntity) error {
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

func (e *emailUseCases) GetMessageSummary(ID string) (entities.GetMessageSummaryEntity, error) {
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

func NewEmailUseCases(cfg *config.Cfg) *emailUseCases {
	return &emailUseCases{
		cfg: cfg,
	}
}
