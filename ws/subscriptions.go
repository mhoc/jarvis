
package ws

import (
  "github.com/mhoc/jarvis/util"
)

var universalSubs = make([]chan map[string]interface{}, 0)
var messageSubs = make([]chan util.IncomingSlackMessage, 0)
var keywordSubs = make(map[string][]chan util.IncomingSlackMessage)

func Dispatch(frame map[string]interface{}) {
  DispatchUniverally(frame)
  DispatchMessage(frame)
}

func DispatchUniverally(frame map[string]interface) {
  for _, subscription := range universalSubs {
    subscription <- frame
  }
}

func DispatchMessage(frame map[string]interface) {
  if t, in := frame["type"]; in && t == "message" {
    msgStruct := util.IncomingSlackMessage{
      Type: msg["type"].(string),
      Channel: msg["channel"].(string),
      User: msg["user"].(string),
      Text: msg["text"].(string),
      Timestamp: msg["ts"].(string),
    }
    for _, sub := range messageSubs {
      sub <- msgStruct
    }
    DispatchKeywords(msgStruct)
  }
}

func DispatchKeywords(message util.IncomingSlackMessage) {
  for keyword, subList := range keywordSubs {

  }
}
