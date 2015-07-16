// Root handler/dispatcher for all "commands" from slack.
// Commands are anything of the form
//    jarvis {ACTION} {arguments}
// Examples
//    jarvis PING www.google.com
//    jarvis LOVE mikehock
//    jarvis GIF cats
//    jarvis GOOGLE weird fetish porn
//    jarvis YOUTUBE Rick Astley

package handlers

import (
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
  "strings"
)

var ch = make(chan util.IncomingSlackMessage)

func InitCommands() {
  log.Info("Initing command listener")
  ws.SubscribeToMessages(ch)
  go BeginCommandLoop()
}

func BeginCommandLoop() {
  for {
    msg := <-ch
    if !IsCommand(msg.Text) {
      continue
    }
  }
}

func IsCommand(text string) bool {
  if !strings.HasPrefix(text, "jarvis") and !strings.HasPrefix(text, "Jarvis") {
    return false
  }

}
