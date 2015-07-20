
package ws

import (
  "fmt"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
  "golang.org/x/net/websocket"
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
  websocket.JSON.Send(wsConnection, msg)
}
