
package ws

import (
  "crypto/tls"
  "github.com/gorilla/websocket"
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
  "net/http"
  "net/url"
)

var (
  wsConnection *websocket.Conn
  slackUrl = "https://slack.com/api/rtm.start?token="
)

func Init() {
  log.Info("Initializing websocket connection to slack")
  u, err := url.Parse(GetSlackWsUrl())
  util.Check(err)
  wsConnection = CreateWebsocket(u)
  go StartReading()
}

func GetSlackWsUrl() string {
  log.Trace("Getting slack websocket url")
  slackAuth := config.SlackAuthToken()
  slackUrl += slackAuth
  data := util.HttpGet(slackUrl)
  StoreJarvisUserId(data)
  return data["url"].(string)
}

func CreateWebsocket(u *url.URL) *websocket.Conn {
  log.Trace("Creating websocket")
  rawConn, err := tls.Dial("tcp", u.Host + ":443", nil)
  util.Check(err)
  wsHeaders := http.Header{
    "Origin": { "http://localhost/" },
  }
  wsConnection, _, err := websocket.NewClient(rawConn, u, wsHeaders, 16384, 16384)
  util.Check(err)
  return wsConnection
}

func StoreJarvisUserId(data map[string]interface{}) {
  jarvisId := data["self"].(map[string]interface{})["id"].(string)
  config.OtherConf.JarvisUserId = jarvisId
}
