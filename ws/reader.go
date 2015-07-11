
package ws

import (
  log "github.com/Sirupsen/logrus"
  "golang.org/x/net/websocket"
)

var receivers []chan map[string]interface{}

func StartReading(ws *websocket.Conn) {
  log.Info("Beginning read loop on websocket")
  for {
    frame := make(map[string]interface{})
    err := websocket.JSON.Receive(ws, &frame)
    Check(err)
    for _, receiver := range receivers {
      receiver <- frame
    }
  }
}

func RegisterMsgReceiver(c chan map[string]interface{}) {
  receivers = append(receivers, c)
}
