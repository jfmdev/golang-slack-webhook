package util

import (
  "encoding/json"
  "math"
  "time"
  "git.mills.io/prologic/bitcask"
)

type BitcoinData struct {
  Price float64 `json:"price"`
  Timestamp time.Time `json:"timestamp"`
}

type JokeData struct {
  Text string `json:"text"`
  Timestamp time.Time `json:"timestamp"`
}

var BITCOIN_KEY = []byte("last_price")
var JOKE_KEY = []byte("last_joke")

// Price should be send if the previous price was sent less than 'timeThreshold' minutes or if the price changed more than
// a 'changeThreshold' percentage.
func ShouldSendBitcoinPrice(myDb *bitcask.Bitcask, timeThreshold int, currentPrice float64, changeThreshold float64) bool {
  var lastData BitcoinData

  lastDataRaw, dbErr := myDb.Get(BITCOIN_KEY)
  if(dbErr != nil) {
    return true
  }

  jsonErr := json.Unmarshal(lastDataRaw, &lastData)
  if(jsonErr != nil) {
    return true
  }

  now := time.Now()
  duration := now.Sub(lastData.Timestamp)
  if(duration.Minutes() > float64(timeThreshold)) {
    return true
  }

  return math.Abs((currentPrice - lastData.Price) / lastData.Price) > changeThreshold
}

// Joke should be send if the previous joke was sent less than 'timeThreshold' minutes.
func ShouldSendJoke(myDb *bitcask.Bitcask, timeThreshold int) bool {
  var lastData JokeData

  lastDataRaw, dbErr := myDb.Get(JOKE_KEY)
  if(dbErr != nil) {
    return true
  }

  jsonErr := json.Unmarshal(lastDataRaw, &lastData)
  if(jsonErr != nil) {
    return true
  }

  now := time.Now()

  duration := now.Sub(lastData.Timestamp)
  return  duration.Minutes() > float64(timeThreshold)
}

func UpdateBitcoinData(myDb *bitcask.Bitcask, price float64) {
  data := &BitcoinData{
    Price: price,
    Timestamp: time.Now()}

  jsonData, _ := json.Marshal(data)
  byteData := []byte(string(jsonData))

  myDb.Put(BITCOIN_KEY, byteData)
}

func UpdateJokeData(myDb *bitcask.Bitcask, joke string) {
  data := &JokeData{
    Text: joke,
    Timestamp: time.Now()}

  jsonData, _ := json.Marshal(data)
  byteData := []byte(string(jsonData))

  myDb.Put(JOKE_KEY, byteData)
}
