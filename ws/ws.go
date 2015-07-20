
package ws

import (
  "encoding/json"
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
  "golang.org/x/net/websocket"
  "io/ioutil"
  "net/http"
)

var slackUrl = "https://slack.com/api/rtm.start?token="
var wsConnection *websocket.Conn

func Init() {
  log.Info("Initializing websocket connection to slack")
  url := GetSlackWsUrl()
  wsConnection = CreateWebsocket(url)
  go StartReading()
}

func GetSlackWsUrl() string {
  log.Info("Getting slack websocket url")
  slackAuth := config.SlackAuthToken()
  slackUrl += slackAuth
  resIo, err := http.Get(slackUrl)
  util.Check(err)
  resb, err := ioutil.ReadAll(resIo.Body)
  util.Check(err)
  var data map[string]interface{}
  err = json.Unmarshal(resb, &data)
  util.Check(err)
  StoreJarvisUserId(data)
  return data["url"].(string)
}

func CreateWebsocket(url string) *websocket.Conn {
  ws, err := websocket.Dial(url, "", "http://localhost/")
  util.Check(err)
  return ws
}

func StoreJarvisUserId(data map[string]interface{}) {
  jarvisId := data["self"].(map[string]interface{})["id"].(string)
  config.OtherConf.JarvisUserId = jarvisId
}
