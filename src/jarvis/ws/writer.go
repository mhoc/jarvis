package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"jarvis/log"
	"jarvis/service"
	"jarvis/util"
	"time"
)

var (
	MasterWriteLimiter = time.Tick(1 * time.Second)
)

func SendMessage(message string, channelId string) {
	<-MasterWriteLimiter
	if len(message) > 25 {
		log.Trace(fmt.Sprintf("Writing message '%v...' to channel '%v'", message[:24], channelId))
	} else {
		log.Trace(fmt.Sprintf("Writing message '%v' to channel '%v'", message, channelId))
	}
	if len(message) > 3000 {
		log.Trace("Attempting to send message that is far too large")
		message = "The message I just attempted to send you exceeds Slack's maximum message length."
	}
	msg := util.OutgoingSlackMessage{
		Channel: channelId,
		Text:    message,
		Type:    "message",
		Id:      1,
	}
	json, err := json.Marshal(msg)
	util.Check(err)
	wsConnection.WriteMessage(websocket.TextMessage, json)
}

func SendPrivateMessage(message string, user string) {
	ch, err := service.Slack{}.IMChannelFromUserId(user)
	if err != nil {
		return
	}
	SendMessage(message, ch)
}
