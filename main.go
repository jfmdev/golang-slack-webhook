package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
  "sync"
  "github.com/joho/godotenv"
  "git.mills.io/prologic/bitcask"
)

import "main/util"

const BITCOIN_ENV_KEY = "BITCOIN_THRESHOLD"
const JOKE_ENV_KEY = "JOKE_THRESHOLD"
const WEBHOOK_ENV_KEY = "WEBHOOK_URL"

type EnvData struct {
  BitboinPriceThreshold float64
  BitcoinTimeThreshold int
  JokeTimeThreshold int
  WebhookUrl string
}

func initialize() (EnvData, error) {
  var data EnvData

  envErr := godotenv.Load(".env")
  if envErr != nil {
    return data, envErr
  }

  bitcoinRaw := strings.Split(os.Getenv(BITCOIN_ENV_KEY), ",")
  bitcoinPrice, _ := strconv.ParseFloat(bitcoinRaw[1], 64)
  bitcoinTime, _ := strconv.Atoi(bitcoinRaw[0])
  jokeTime, _ := strconv.Atoi(os.Getenv(JOKE_ENV_KEY))

  data = EnvData{
    BitboinPriceThreshold: bitcoinPrice,
    BitcoinTimeThreshold: bitcoinTime,
    JokeTimeThreshold: jokeTime,
    WebhookUrl: os.Getenv(WEBHOOK_ENV_KEY)}
  return data, nil
}

func isCronEnabled() bool {
  // CAVEAT: We assume that the "--cron" flag will always be the first parameter.
  return len(os.Args) >= 2 && os.Args[1] == "--cron"
}

func main() {
  // --- Initialization --- //

  envData, initErr := initialize()
  if initErr != nil {
    fmt.Printf("Application couldn't be initialized | %s", initErr)
    return
  }

  // CAVEAT: It's assumed that the "--cron" flag will always be the first parameter.
  isCronEnabled := len(os.Args) >= 2 && os.Args[1] == "--cron"

  myDb, _ := bitcask.Open("/tmp/db")
  defer myDb.Close()

  // --- Execution --- //

  var wg sync.WaitGroup

  wg.Add(1)
  go func() {
    defer wg.Done()

    if(isCronEnabled && !util.ShouldSendJoke(myDb, envData.JokeTimeThreshold)) {
      return
    }

    jokeRes, jokeErr := util.GetJoke()
    if jokeErr == nil {
      util.PostMessageOnSlack(envData.WebhookUrl, fmt.Sprintf("*Joke of the day:* %s", jokeRes.Joke))
      util.UpdateJokeData(myDb, jokeRes.Joke)
    } else {
      fmt.Printf("Joke couldn't be fetched | %s \n", jokeErr)
    }
  }()

  wg.Add(1)
  go func() {
    defer wg.Done()

    bitcoinPrice, bitcoinErr := util.GetBitcoinPrice()
    if bitcoinErr == nil {
      if(isCronEnabled && !util.ShouldSendBitcoinPrice(myDb, 
        envData.BitcoinTimeThreshold, 
        bitcoinPrice, 
        envData.BitboinPriceThreshold)) {
        return
      }

      util.PostMessageOnSlack(envData.WebhookUrl, 
        fmt.Sprintf("*Current Bitcoing price (in USD):* %s", strconv.FormatFloat(bitcoinPrice, 'f', 2, 64)))
      util.UpdateBitcoinData(myDb, bitcoinPrice)
    } else {
      fmt.Printf("Bitcoing price couldn't be fetched | %s \n", bitcoinErr)
    }
  }()

  wg.Wait()
}
