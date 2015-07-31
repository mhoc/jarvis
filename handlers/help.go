
package handlers

import (
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
)

var helpRegex = util.NewRegex("^jarvis help$")
var helpCommandRegex = util.NewRegex("jarvis help ([a-z]+)")
var helpCh = make(chan util.IncomingSlackMessage)

func InitHelp() {
  log.Info("Initing help handler")
  ws.SubscribeToMessages(helpCh)
  go BeginHelpLoop()
}

func BeginHelpLoop() {
  for {
    msg := <-helpCh
    if helpRegex.Matches(msg.Text) {
      generalHelp(msg)
    } else if helpCommandRegex.Matches(msg.Text) {
      commandHelp(msg)
    }
  }
}

func generalHelp(msg util.IncomingSlackMessage) {
  resp := "Hi. I'm Jarvis, your friendly neighborhood slackbot.\n"
  resp += "You can access additional information about a given command using the syntax `jarvis help (command)`\n"
  resp += "The commands I have installed include:\n"
  for key, _ := range CommandManifest {
    resp += "`" + key + "`, "
  }
  ws.SendMessage(resp[:len(resp)-2], msg.Channel)
}

func commandHelp(msg util.IncomingSlackMessage) {
  cmdName := helpCommandRegex.SubExpression(msg.Text, 0)
  if cmd, in := CommandManifest[cmdName]; in {
    cmd.Help(msg)
  } else {
    ws.SendMessage("I don't seem to have a record of that command.", msg.Channel)
  }
}
