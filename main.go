package main

import (
  "fmt"
  "os"
  "github.com/joho/godotenv"
)

import "main/util"

const WEBHOOK_ENV_KEY = "WEBHOOK_URL"

func initialize() (string, error) {
  envErr := godotenv.Load(".env")
  if envErr != nil {
    return "", envErr
  }

  return os.Getenv(WEBHOOK_ENV_KEY), nil
}

// TODO: Execute requests in parallel.
// TODO: Add thresholds to not send messages too often in case this application is run by a Cron job.
func main() {
  webhookUrl, initErr := initialize()
  if initErr != nil {
    fmt.Printf("Application couldn't be initialized | %s", initErr)
    return
  }

  jokeRes, jokeErr := util.GetJoke()
  if jokeErr == nil {
    util.PostMessageOnSlack(webhookUrl, fmt.Sprintf("*Joke of the day:* %s", jokeRes.Joke))
  } else {
    fmt.Printf("Joke couldn't be fetched | %s \n", jokeErr)
  }

  bitcoinPrice, bitcoinErr := util.GetBitcoinPrice()
  if jokeErr == nil {
    util.PostMessageOnSlack(webhookUrl, fmt.Sprintf("Current Bitcoing price (in USD): *%s*", bitcoinPrice))
  } else {
    fmt.Printf("Bitcoing price couldn't be fetched | %s \n", bitcoinErr)
  }
}
