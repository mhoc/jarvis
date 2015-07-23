
package handlers

import (
  "github.com/mhoc/jarvis/commands"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
  "strings"
)

var commandManifest = []util.Command{
  commands.Help{},
  commands.Status{},
  commands.Weather{},
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
    cmd := MatchCommand(msg.Text)
    if cmd != nil {
      go cmd.Execute(msg)
    }
  }
}

func IsCommand(text string) bool {
  if strings.Contains(text, "jarvis") {
    return true
  }
  if strings.Contains(text, "Jarvis") {
    return true
  }
  return false
}

func MatchCommand(text string) util.Command {
  for _, command := range commandManifest {
    for _, match := range command.Matches() {
      if match.MatchString(text) {
        return command
      }
    }
  }
  return nil
}
