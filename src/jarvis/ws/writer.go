
package ws

import (
  "encoding/json"
  "fmt"
  "github.com/gorilla/websocket"
  "jarvis/log"
  "jarvis/service"
  "jarvis/util"
)

func SendMessage(message string, channelId string) {
  if len(message) > 25 {
    log.Trace(fmt.Sprintf("Writing message '%v...' to channel '%v'", message[:24], channelId))
  } else {
    log.Trace(fmt.Sprintf("Writing message '%v' to channel '%v'", message, channelId))
  }
  msg := util.OutgoingSlackMessage{
    Channel: channelId,
    Text: message,
    Type: "message",
    Id: 1,
  }
  json, err := json.Marshal(msg)
  util.Check(err)
  wsConnection.WriteMessage(websocket.TextMessage, json)
}

func SendPrivateMessage(message string, user string) {
  ch := service.Slack{}.IMChannelFromUserId(user)
  SendMessage(message, ch)
}
