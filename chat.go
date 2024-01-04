package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func (r ChatRunner) SendMessage(webhookUrl string, message *Message) {
	if ChatDebugable {
		r.logger.Debugf("messaeg: %#v", message)
		body, _ := jsonMarshal(message)
		r.logger.Debugf(string(body))
	}
	messageBytes, err := jsonMarshal(message)
	if err != nil {
		r.logger.Errorf("Error occurred while marshalling message: %s", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", webhookUrl, bytes.NewBuffer(messageBytes))
	if err != nil {
		r.logger.Errorf("Error occurred while creating HTTP request: %s", err)
	}

	// Set the content type to application/json
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		r.logger.Errorf("Error occurred while sending request to webhook: %s", err)
	}
	defer resp.Body.Close()

	// // Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		msg, _ := io.ReadAll(resp.Body)
		r.logger.Errorf("Non-OK HTTP status: %s \n %s", resp.Status, msg)
	}
}

func jsonMarshal(v any) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
