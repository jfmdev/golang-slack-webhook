package util

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
)

const BITCOIN_RATE_URL = "https://api.coinbase.com/v2/exchange-rates?currency=BTC"

type cbResponse struct {
  Data cbData `json:"data"`
}

type cbData struct {
  Currency string         `json:"currency"`
  Rates map[string]string `json:"rates"`
}

func GetBitcoinPrice() (string, error) {
  var cbRes cbResponse

  response, getErr := http.Get(BITCOIN_RATE_URL)
  if getErr != nil {
    return "", getErr
  }
  defer response.Body.Close()

  body, readErr := ioutil.ReadAll(response.Body)
  if readErr != nil {
    return "", readErr
  }

  json.Unmarshal(body, &cbRes)
  return cbRes.Data.Rates["USD"], nil
}
