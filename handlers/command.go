
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
  if strings.Contains(text, "jarvis") {
    return true
  }
  if strings.Contains(text, "Jarvis") {
    return true
  }
  return false
}

func MatchCommand(msg util.IncomingSlackMessage) util.Command {
  if strings.Contains(msg.Text, "help") {
    HelpHandler(msg)
    return nil
  }
  for _, command := range CommandManifest {
    for _, regex := range command.Matches() {
      if regex.MatchString(msg.Text) {
        return command
      }
    }
  }
  return nil
}

func HelpHandler(msg util.IncomingSlackMessage) {
  i := strings.Index(msg.Text, "help")
  if i + len("help ") > len(msg.Text) {
    ws.SendMessage(GeneralHelp(), msg.Channel)
  } else {
    cmd := CommandManifest[msg.Text[i + len("help "):]]
    if cmd == nil {
      ws.SendMessage("I don't recognize the command you're asking for help on.", msg.Channel)
    } else {
      cmd.Help(msg)
    }
  }
}

func GeneralHelp() string {
  resp := "Hi. I'm Jarvis, your friendly neighborhood slackbot.\n"
  resp += "You can access additional information about a given command using the syntax `jarvis help (command)`\n"
  resp += "The commands I have installed include:\n"
  for key, _ := range CommandManifest {
    resp += "`" + key + "`, "
  }
  return resp[:len(resp)-2]
}
