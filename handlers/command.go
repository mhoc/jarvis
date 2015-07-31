
package handlers

import (
  "github.com/mhoc/jarvis/commands"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
  "strings"
)

// This is a super awkward way of filling this map, but it ensures that the
// data for command names is stored in a single location. I'll look for ways of
// improving it eventually.
var CommandManifest = map[string]util.Command{
  commands.Remember{}.Name(): commands.Remember{},
  commands.Status{}.Name(): commands.Status{},
  commands.Weather{}.Name(): commands.Weather{},
}

var cmdCh = make(chan util.IncomingSlackMessage)

func InitCommands() {
  log.Info("Initing command listener")
  ws.SubscribeToMessages(cmdCh)
  go BeginCommandLoop()
}

func BeginCommandLoop() {
  for {
    msg := <-cmdCh
    if !IsCommand(msg.Text) {
      continue
    }
    cmd := MatchCommand(msg)
    if cmd != nil {
      go cmd.Execute(msg)
    }
  }
}

func IsCommand(text string) bool {
  if strings.Contains(text, "help") {
    return false
  }
  if strings.Contains(text, "jarvis") {
    return true
  }
  if strings.Contains(text, "Jarvis") {
    return true
  }
  return false
}

func MatchCommand(msg util.IncomingSlackMessage) util.Command {
  for _, command := range CommandManifest {
    for _, regex := range command.Matches() {
      if regex.MatchString(msg.Text) {
        return command
      }
    }
  }
  return nil
}
