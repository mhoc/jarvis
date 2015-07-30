
package handlers

import (
  "github.com/mhoc/jarvis/commands"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
  "strings"
)

var commandManifest = map[string]util.Command{
  "remember": commands.Remember{},
  "status": commands.Status{},
  "weather": commands.Weather{},
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
  for _, command := range commandManifest {
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
    cmd := commandManifest[msg.Text[i + len("help "):]]
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
  for key, _ := range commandManifest {
    resp += "`" + key + "`, "
  }
  return resp[:len(resp)-2]
}
