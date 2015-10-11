

package commands

import (
  "jarvis/config"
  "jarvis/data"
  "jarvis/util"
  "jarvis/ws"
  "os"
)

type Debug struct {}

func NewDebug() Debug {
  return Debug{}
}

func (c Debug) Name() string {
  return "debug"
}

func (c Debug) Description() string {
  return "provides admin level debug functionality into the jarvis and slack internals."
}

func (c Debug) Examples() []string {
  return []string{"jarvis debug suicide"}
}

func (c Debug) OtherDocs() []util.HelpTopic {
  return []util.HelpTopic{}
}

func (c Debug) SubCommands() []util.SubCommand {
  return []util.SubCommand{
    util.NewSubCommand("^jarvis debug dump redis$", c.DumpRedis),
    util.NewSubCommand("^jarvis debug suicide$", c.Suicide),
    util.NewSubCommand("^jarvis debug info$", c.Info),
  }
}

func (c Debug) IsAdmin(m util.IncomingSlackMessage) bool {
  if !config.IsAdmin(m.User) {
    ws.SendMessage("*ACCESS DENIED* :police_car: *THIS ATTEMPT HAS BEEN REPORTED* :police_car: *ACCESS DENIED*", m.Channel)
    return false
  }
  return true
}

func (c Debug) DumpRedis(m util.IncomingSlackMessage, r util.Regex) {
  if !c.IsAdmin(m) {
    return
  }
  keys := data.Keys("*")
  resp := "I'm currently tracking these keys in storage:\n"
  for _, key := range keys {
    _, v := data.Get(key)
    resp += key + ": " + v + "\n"
  }
  ws.SendMessage(resp, m.Channel)
}

func (c Debug) Suicide(m util.IncomingSlackMessage, r util.Regex) {
  if !c.IsAdmin(m) {
    return
  }
  ws.SendMessage("You got it boss, nighty night.", m.Channel)
  os.Exit(0)
}

func (c Debug) Info(m util.IncomingSlackMessage, r util.Regex) {
  if !c.IsAdmin(m) {
    return
  }
  resp := "You are user " + m.User + "\n"
  resp += "We are in channel " + m.Channel
  ws.SendMessage(resp, m.Channel)
}
