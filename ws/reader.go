
package ws

import (
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
  "golang.org/x/net/websocket"
)

var allReceivers = make([]chan map[string]interface{}, 0)
var msgReceivers = make([]chan util.IncomingSlackMessage, 0)

func StartReading() {
  log.Info("Beginning read loop on websocket")
  for {
    frame := make(map[string]interface{})
    websocket.JSON.Receive(wsConnection, &frame)
    if len(frame) == 0 {
      continue
    }
    if sender, in := frame["user"]; in && sender == config.JarvisUserId() {
      log.Trace("Ignoring message sent by jarvis")
      continue
    }
    go Dispatch(frame)
  }
}

func Dispatch(msg map[string]interface{}) {
  DispatchAll(msg)
  if msg["type"] == "message" {
    msgStruct := util.IncomingSlackMessage{
      Type: msg["type"].(string),
      Channel: msg["channel"].(string),
      User: msg["user"].(string),
      Text: msg["text"].(string),
      Timestamp: msg["ts"].(string),
    }
    DispatchMessage(msgStruct)
  }
}

func DispatchAll(msg map[string]interface{}) {
  for _, receiver := range allReceivers {
    select {
    case receiver <- msg:
    default:
    }
  }
}

func DispatchMessage(msg util.IncomingSlackMessage) {
  for _, receiver := range msgReceivers {
    select {
    case receiver <- msg:
    default:
    }
  }
}

func SubscribeToAll(c chan map[string]interface{}) {
  allReceivers = append(allReceivers, c)
}

func SubscribeToMessages(c chan util.IncomingSlackMessage) {
  msgReceivers = append(msgReceivers, c)
}
