package util

import (
  "bytes"
  "encoding/json"
  "net/http"
)

type SlackMessage struct {
  Text string `json:"text"`
}

func PostMessageOnSlack(webhookUrl string, message string) error {
  data := &SlackMessage{Text: message}
  jsonData, _ := json.Marshal(data)
  byteData := []byte(string(jsonData))

  request, reqError := http.NewRequest("POST", webhookUrl, bytes.NewBuffer(byteData))
  request.Header.Set("Content-Type", "application/json; charset=UTF-8")

  client := &http.Client{}
  response, reqError := client.Do(request)
  if reqError != nil {
    return reqError;
  }
  defer response.Body.Close()

  return nil
}