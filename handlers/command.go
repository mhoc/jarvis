
package handlers

import (
  "github.com/mhoc/jarvis/commands"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
  "strings"
)
var CommandManifest map[string]util.Command
var cmdCh = make(chan util.IncomingSlackMessage)

func InitCommands() {
  log.Info("Initing command listener")
  // This is a super awkward way of filling this map, but it ensures that the
  // data for command names is stored in a single location. I'll look for ways of
  // improving it eventually.
  CommandManifest = map[string]util.Command{
    commands.Debug{}.Name(): commands.NewDebug(),
    commands.Recall{}.Name(): commands.NewRecall(),
    commands.Remember{}.Name(): commands.NewRemember(),
    commands.Static{}.Name(): commands.NewStatic(),
    commands.Status{}.Name(): commands.NewStatus(),
    commands.Weather{}.Name(): commands.NewWeather(),
  }
  ws.SubscribeToMessages(cmdCh)
  go BeginCommandLoop()
}

func BeginCommandLoop() {
  for {
    msg := <-cmdCh
    if !IsCommand(msg) {
      continue
    }
    cmd := MatchCommand(msg)
    if cmd != nil {
      go cmd.Execute(msg)
    }
  }
}

func IsCommand(msg util.IncomingSlackMessage) bool {
  if strings.Contains(msg.Text, "help") {
    return false
  }
  if strings.Split(msg.Text, " ")[0] == "jarvis" {
    return true
  }
  if strings.Split(msg.Text, " ")[0] == "Jarvis" {
    return true
  }
  if strings.Contains(msg.Text, "jarivs") {
    ws.SendMessage("Dude, you can't even spell my name right? Whatever.", msg.Channel)
    return true
  }
  return false
}

func MatchCommand(msg util.IncomingSlackMessage) util.Command {
  for _, command := range CommandManifest {
    for _, regex := range command.Matches() {
      if regex.Matches(msg.Text) {
        return command
      }
    }
  }
  return nil
}
