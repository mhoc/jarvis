
package ws

import (
  "github.com/mhoc/jarvis/log"
  "golang.org/x/net/websocket"
)

var receivers = make([]chan map[string]interface{}, 0)

func StartReading(ws *websocket.Conn) {
  log.Info("Beginning read loop on websocket")
  for {
    frame := make(map[string]interface{})
    websocket.JSON.Receive(ws, &frame)
    if len(frame) == 0 {
      continue
    }
    for _, receiver := range receivers {
      select {
        case receiver <- frame:
        default:
      }
    }
  }
}

func RegisterMsgReceiver(c chan map[string]interface{}) {
  receivers = append(receivers, c)
}
