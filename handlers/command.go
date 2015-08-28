
package handlers

import (
  "github.com/mhoc/jarvis/commands"
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
  "strings"
)

var commandRegex = util.NewRegex("^[Jj]arvis [A-Za-z0-9 ]+")
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
    if !IsCommand(&msg) {
      continue
    }
    if config.ChannelIsBlacklisted(msg.Channel) {
      log.Trace("Ignoring message sent on blacklisted channel %v", msg.Channel)
      continue
    }
    if config.HasWhitelist() && !config.ChannelIsWhitelisted(msg.Channel) {
      log.Trace("Running with whitelist. Ignoring message not sent on whitelisted channel %v", msg.Channel)
      continue
    }
    FormatCommand(&msg)
    MatchCommand(msg)
  }
}

func IsCommand(msg *util.IncomingSlackMessage) bool {
  if strings.Contains(msg.Text, "help") {
    return false
  }
  if commandRegex.Matches(msg.Text) {
    return true
  }
  return false
}

func FormatCommand(msg *util.IncomingSlackMessage) {
  msg.Text = strings.Replace(msg.Text, "Jarvis", "jarvis", -1)
}

func MatchCommand(msg util.IncomingSlackMessage) {
  for _, command := range CommandManifest {
    for _, subcommand := range command.SubCommands() {
      if subcommand.Pattern.Matches(msg.Text) {
        go subcommand.Exec(msg, subcommand.Pattern)
      }
    }
  }
}
