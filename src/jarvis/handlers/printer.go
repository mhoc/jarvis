package handlers

import (
	"fmt"
	"jarvis/log"
	"jarvis/service"
	"jarvis/ws"
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
	if name, in := msg["user"]; in {
		userName := service.Slack{}.UserNameFromUserId(name.(string))
		log.Info(fmt.Sprintf("%v: %v", userName, msg["text"]))
	}
}
