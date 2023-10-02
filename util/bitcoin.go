package util

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
  "strconv"
)

const BITCOIN_RATE_URL = "https://api.coinbase.com/v2/exchange-rates?currency=BTC"

type cbResponse struct {
  Data cbData `json:"data"`
}

type cbData struct {
  Currency string         `json:"currency"`
  Rates map[string]string `json:"rates"`
}

func GetBitcoinPrice() (float64, error) {
  var cbRes cbResponse

  response, getErr := http.Get(BITCOIN_RATE_URL)
  if getErr != nil {
    return 0, getErr
  }
  defer response.Body.Close()

  body, readErr := ioutil.ReadAll(response.Body)
  if readErr != nil {
    return 0, readErr
  }

  json.Unmarshal(body, &cbRes)
  price, parseErr := strconv.ParseFloat(cbRes.Data.Rates["USD"], 64)
  if parseErr != nil {
    return 0, parseErr
  }

  return price, nil
}
