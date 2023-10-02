package util

import (
  "fmt"
)

func PostMessageOnSlack(url string, message string) error {
  fmt.Printf("TODO | Post to %s the message: %s \n", url, message)
  return nil
}