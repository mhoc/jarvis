
package ws

import (
  "encoding/json"
  "github.com/mhoc/jarvis/config"
  log "github.com/Sirupsen/logrus"
  "golang.org/x/net/websocket"
  "io/ioutil"
  "net/http"
)

var slackUrl = "https://slack.com/api/rtm.start?token="

func Init() {
  log.Info("Initializing websocket connection to slack")
  url := GetSlackWsUrl()
  ws := CreateWebsocket(url)
  go StartReading(ws)
}

func GetSlackWsUrl() string {
  log.Info("Getting slack websocket url")
  slackAuth := config.SlackAuthToken()
  slackUrl += slackAuth
  resIo, err := http.Get(slackUrl)
  Check(err)
  resb, err := ioutil.ReadAll(resIo.Body)
  Check(err)
  var data map[string]interface{}
  err = json.Unmarshal(resb, &data)
  Check(err)
  return data["url"].(string)
}

func CreateWebsocket(url string) *websocket.Conn {
  ws, err := websocket.Dial(url, "", "http://localhost/")
  Check(err)
  return ws
}

func Check(e error) {
  if e != nil {
    log.Fatal(e)
  }
}
