package mailer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Mailer struct {
	key     string
	address string
}

type message struct {
	Sender   string   `json:"sender"`
	To       []string `json:"to"`
	Subject  string   `json:"subject"`
	TextBody string   `json:"text_body"`
	HtmlBody string   `json:"html_body"`
}

func New(key, address string) *Mailer {
	return &Mailer{
		key,
		address,
	}
}

func (m *Mailer) Send(
	destinationAddress,
	subject,
	textContent,
	htmlContent string) error {

	body := message{
		Sender:  m.address,
		To:      []string{destinationAddress},
		Subject: subject,
	}

	if textContent != "" {
		body.TextBody = textContent
	} else if htmlContent != "" {
		body.HtmlBody = htmlContent
	} else {
		return errors.New("No message body was provided")
	}

	jsonBytes, err := json.Marshal(body)

	r, err := http.NewRequest(
		"POST",
		"https://api.smtp2go.com/v3/email/send",
		bytes.NewBuffer(jsonBytes),
	)

	if err != nil {
		return err
	}

	headers := map[string]string{
		"X-Smtp2go-Api-Key": m.key,
		"Content-Type":      "application/json",
		"accept":            "application/json",
	}

	for k, v := range headers {
		r.Header.Add(k, v)
	}

	client := http.Client{}

	defer r.Body.Close()

	response, err := client.Do(r)
	_, err = io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	fmt.Println(response)

	fmt.Println(r)
	return nil
}
