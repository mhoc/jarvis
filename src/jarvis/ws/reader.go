package ws

import (
	"encoding/json"
	"jarvis/data"
	"jarvis/log"
	"jarvis/util"
)

var allReceivers = make([]chan map[string]interface{}, 0)
var msgReceivers = make([]chan util.IncomingSlackMessage, 0)

func StartReading() {
	log.Info("Beginning read loop on websocket")
	for {
		_, p, err := wsConnection.ReadMessage()
		util.Check(err)
		var frame map[string]interface{}
		json.Unmarshal(p, &frame)
		if err != nil {
			log.Warn("Websocket read threw an error.")
			log.Warn(err.Error())
		}
		if len(frame) == 0 {
			continue
		}
		_, jarvisUserId := data.Get("jarvis-user-id")
		if sender, in := frame["user"]; in && sender == jarvisUserId {
			log.Trace("Ignoring message sent by jarvis")
			continue
		}
		go Dispatch(frame)
	}
}

func Dispatch(msg map[string]interface{}) {
	DispatchAll(msg)
	if msg["type"] == "message" {
		if !util.MapHasElements(msg, "type", "channel", "user", "text", "ts") {
			return
		}
		msgStruct := util.IncomingSlackMessage{
			Type:      msg["type"].(string),
			Channel:   msg["channel"].(string),
			User:      msg["user"].(string),
			Text:      msg["text"].(string),
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
