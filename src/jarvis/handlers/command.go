
package handlers

import (
  "jarvis/commands"
  "jarvis/config"
  "jarvis/data"
  "jarvis/log"
  "jarvis/util"
  "jarvis/ws"
  "strings"
  "time"
)

var (
  commandRegex = util.NewRegex("^[Jj]arvis [A-Za-z0-9 ]+")
  CommandManifest map[string]util.Command
  cmdCh = make(chan util.IncomingSlackMessage)
  ratelimitMap = make(map[string]<-chan time.Time)
)

func InitCommands() {
  log.Info("Initing command listener")
  // This is a super awkward way of filling this map, but it ensures that the
  // data for command names is stored in a single location. I'll look for ways of
  // improving it eventually.
  CommandManifest = map[string]util.Command{
    commands.Debug{}.Name(): commands.NewDebug(),
    commands.OnionHoroscope{}.Name(): commands.OnionHoroscope{},
    commands.Recall{}.Name(): commands.NewRecall(),
    commands.Remember{}.Name(): commands.NewRemember(),
    commands.Remind{}.Name(): commands.NewRemind(),
    commands.Static{}.Name(): commands.NewStatic(),
    commands.Status{}.Name(): commands.NewStatus(),
    commands.Weather{}.Name(): commands.NewWeather(),
    commands.SIGAPP{}.Name(): commands.NewSIGAPP(),
  }
  ws.SubscribeToMessages(cmdCh)
  go BeginCommandLoop()
}

func BeginCommandLoop() {
  shitlist := data.GetShitlist()
  lastShitlistUpdate := time.Now()
  for {
    msg := <-cmdCh
    send := true
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
    if time.Now().Sub(lastShitlistUpdate).Minutes() > 10 {
      shitlist = data.GetShitlist()
      lastShitlistUpdate = time.Now()
    }
    for _, user := range shitlist {
      if msg.User == user {
        log.Trace("Shitlisted user trying to send another message, nuke it.")
        send = false
        break
      }
    }
    if send {
      go func() {
        RatelimitUser(msg)
        FormatCommand(&msg)
        MatchCommand(msg)
      }()
    }
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

// This is a basic ratelimit with burst. I'd like to eventually use the
// shitlist capability exposed in the data package but right now this
// should work quite well.
func RatelimitUser(msg util.IncomingSlackMessage) {
  if _, in := ratelimitMap[msg.User]; !in {
    ratelimitMap[msg.User] = time.Tick(1 * time.Second)
  }
  <-ratelimitMap[msg.User]
}

func FormatCommand(msg *util.IncomingSlackMessage) {
  msg.Text = strings.Replace(msg.Text, "Jarvis", "jarvis", -1)
}

func MatchCommand(msg util.IncomingSlackMessage) {
  for _, command := range CommandManifest {
    for _, subcommand := range command.SubCommands() {
      if subcommand.Pattern.Matches(msg.Text) {
        subcommand.Exec(msg, subcommand.Pattern)
        return
      }
    }
  }
}
