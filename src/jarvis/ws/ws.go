
package ws

import (
  "crypto/tls"
  "github.com/gorilla/websocket"
  "jarvis/config"
  "jarvis/data"
  "jarvis/log"
  "jarvis/util"
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
  data, err := util.HttpGet(slackUrl)
  util.Check(err)
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

func StoreJarvisUserId(d map[string]interface{}) {
  jarvisId := d["self"].(map[string]interface{})["id"].(string)
  data.Set("jarvis-user-id", jarvisId)
}
