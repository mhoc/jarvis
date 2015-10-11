
package handlers

import (
  "jarvis/config"
  "jarvis/log"
  "jarvis/util"
  "jarvis/ws"
  "strings"
)

var helpRegex = util.NewRegex("^[Jj]arvis help$")
var helpCommandRegex = util.NewRegex("^[Jj]arvis help ([A-Za-z]+)$")
var helpCh = make(chan util.IncomingSlackMessage)

func InitHelp() {
  log.Info("Initing help handler")
  ws.SubscribeToMessages(helpCh)
  go BeginHelpLoop()
}

func BeginHelpLoop() {
  for {
    msg := <-helpCh
    if config.ChannelIsBlacklisted(msg.Channel) {
      continue
    }
    if config.HasWhitelist() && !config.ChannelIsWhitelisted(msg.Channel) {
      continue
    }
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
    ws.SendMessage(helpGenerate(cmd), msg.Channel)
  } else {
    ws.SendMessage("I don't seem to have a record of that command.", msg.Channel)
  }
}

func helpGenerate(c util.Command) string {
  help := "```\n"
  help += c.Name() + "\n"
  help += "  " + strings.Replace(c.Description(), "\n", "\n  ", -1) + "\n\n"
  help += "matches on\n"
  for _, match := range c.SubCommands() {
    help += "  " + match.Pattern.String() + "\n"
  }
  help += "\nexamples\n"
  for _, ex := range c.Examples() {
    help += "  " + ex + "\n"
  }
  for _, topic := range c.OtherDocs() {
    help += "\n" + topic.Name + "\n"
    help += "  " + strings.Replace(topic.Body, "\n", "\n  ", -1) + "\n"
  }
  help += "```"
  return help
}
