
package handlers

import (
  "fmt"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/ws"
)

var ch = make(chan map[string]interface{})

func Init() {
  log.Info("Registering message logging receiver")
  ws.RegisterMsgReceiver(ch)
  go ReadMessage()
}

func ReadMessage() {
  for {
    msg := <-ch
    switch msg["type"] {
    case "message":
      PrintMessage(msg)
    }
  }
}

func PrintMessage(msg map[string]interface{}) {
  log.Info(fmt.Sprintf("%v: %v\n", msg["user"], msg["text"]))
}
