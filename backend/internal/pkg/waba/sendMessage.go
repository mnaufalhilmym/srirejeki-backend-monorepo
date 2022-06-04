package waba

import (
	"bytes"
	"errors"
	"fmt"
	"greenhouse-monitoring-iot/internal/pkg/errorHandler"
	"io"
	"log"
	"net/http"
)

func SendMessage(phoneNumber, text string) error {
	token := "EAAudNu4qXHkBABJvfeGLV9yYVEKGYzn7Din5jcxHjZBVHtvIKYbmXt1rCPeZAdzhcRP4WAxcOZCq0TxfkpSzwPQtBPC0JEK5dkO75L5yrCx0JrvJZCPA3hC6a6yD265xlIDKBKeQoJZCZBYTgbjhT3IJSm9IqhKhhPZCuXVUjQGMFtgSCgaEguu"
	phoneNumberID := "105828275488538"
	url := "https://graph.facebook.com/v14.0/" + phoneNumberID + "/messages"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(fmt.Sprintf(`{ "messaging_product": "whatsapp", "recipient_type": "individual", "to": "%s", "type": "text", "text": { "preview_url": false, "body": "%s" } }`, phoneNumber, text))))
	defer func() {
		if err := req.Body.Close(); err != nil {
			errorHandler.LogErrorThenContinue("waba/SendMessageDeferReq", err)
		}
	}()
	if err != nil {
		errorHandler.LogErrorThenContinue("waba/SendMessage1", err)
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	defer func() {
		if err := res.Body.Close(); err != nil {
			errorHandler.LogErrorThenContinue("waba/SendMessageDeferRes", err)
		}
	}()
	if err != nil {
		errorHandler.LogErrorThenContinue("waba/SendMessage2", err)
		return err
	}
	if res.Status[:1] != "2" {
		response, err := io.ReadAll(res.Body)
		if err != nil {
			errorHandler.LogErrorThenContinue("waba/SendMessage3", err)
			return err
		}
		err = errors.New("Failed to send WhatsApp chat: " + text + "\nRespons Body: " + string(response))
		errorHandler.LogErrorThenContinue("waba/SendMessage4", err)
		return err
	}
	log.Println("Successfully sent WhatsApp chat: " + text)
	return nil
}
