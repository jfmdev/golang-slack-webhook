package util

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
)

const JOKES_URL = "https://v2.jokeapi.dev/joke/Any?type=single"

type Joke struct {
  Category string `json:"category"`
  Error    bool   `json:"error"`
  Id       int    `json:"id"`
  Joke     string `json:"joke"`
  Lang     string `json:"lang"`
  Type     string `json:"type"`
}

func GetJoke() (Joke, error) {
  var joke Joke

  response, getErr := http.Get(JOKES_URL)
  if getErr != nil {
    return joke, getErr
  }
  defer response.Body.Close()

  body, readErr := ioutil.ReadAll(response.Body)
  if readErr != nil {
    return joke, readErr
  }

  json.Unmarshal(body, &joke)
  return joke, nil
}
