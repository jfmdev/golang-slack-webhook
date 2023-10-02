package main

import (
  "fmt"
  "os"
  "strconv"
  "sync"
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

// TODO: Add thresholds to not send messages too often in case this application is run by a Cron job.
func main() {
  webhookUrl, initErr := initialize()
  if initErr != nil {
    fmt.Printf("Application couldn't be initialized | %s", initErr)
    return
  }

  var wg sync.WaitGroup

  wg.Add(1)
  go func() {
    defer wg.Done()

    jokeRes, jokeErr := util.GetJoke()
    if jokeErr == nil {
      util.PostMessageOnSlack(webhookUrl, fmt.Sprintf("*Joke of the day:* %s", jokeRes.Joke))
    } else {
      fmt.Printf("Joke couldn't be fetched | %s \n", jokeErr)
    }
  }()

  wg.Add(1)
  go func() {
    defer wg.Done()

    bitcoinPrice, bitcoinErr := util.GetBitcoinPrice()
    if bitcoinErr == nil {
      util.PostMessageOnSlack(webhookUrl, 
        fmt.Sprintf("*Current Bitcoing price (in USD):* %s", strconv.FormatFloat(bitcoinPrice, 'f', 2, 64)))
    } else {
      fmt.Printf("Bitcoing price couldn't be fetched | %s \n", bitcoinErr)
    }
  }()

  wg.Wait()
}
