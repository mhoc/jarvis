
package handlers

import (
  "fmt"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/service"
  "github.com/mhoc/jarvis/ws"
)

var printCh = make(chan map[string]interface{})

func InitPrinter() {
  log.Info("Registering message printing receiver")
  ws.SubscribeToAll(printCh)
  go BeginPrintLoop()
}

func BeginPrintLoop() {
  for {
    msg := <-printCh
    switch msg["type"] {
    case "message":
      PrintMessage(msg)
    }
  }
}

func PrintMessage(msg map[string]interface{}) {
  userName := service.Slack{}.UserNameFromUserId(msg["user"].(string))
  log.Info(fmt.Sprintf("%v: %v", userName, msg["text"]))
}
