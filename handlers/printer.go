
package handlers

import (
  "fmt"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/ws"
)

var ch = make(chan map[string]interface{})

func InitPrinter() {
  log.Info("Registering message logging receiver")
  ws.SubscribeToAll(ch)
  go BeginPrintLoop()
}

func BeginPrintLoop() {
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
