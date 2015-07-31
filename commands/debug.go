

package commands

import (
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/data"
  "github.com/mhoc/jarvis/service"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
  "os"
)

type Debug struct {}

var DebugAccessAttempts = make([]string, 0)

func (c Debug) Name() string {
  return "debug"
}

func (c Debug) Matches() []util.Regex {
  return []util.Regex{
    util.NewRegex("debug"),
  }
}

func (c Debug) Description() string {
  return "provides various debug utilities for inspecting jarvis behavior"
}

func (c Debug) Format() string {
  return "jarvis debug (debug command)"
}

func (c Debug) Examples() []string {
  return []string{"if you should use this command you already know how to use it"}
}

func (c Debug) OtherDocs() []util.HelpTopic {
  return []util.HelpTopic{}
}

func (c Debug) Execute(m util.IncomingSlackMessage) {
  if !config.IsAdmin(m.User) {
    ws.SendMessage("*ACCESS DENIED* :police_car: *THIS ATTEMPT HAS BEEN REPORTED* :police_car: *ACCESS DENIED*", m.Channel)
    name := service.Slack{}.UserNameFromUserId(m.User)
    DebugAccessAttempts = append(DebugAccessAttempts, "*" + name + "*: " + m.Text)
    return
  }
  c.debugSuicide(m)
  c.debugAttempts(m)
  c.debugDumpRedisKeys(m)
}

func (c Debug) debugSuicide(m util.IncomingSlackMessage) {
  suicide := util.NewRegex("jarvis debug suicide")
  if suicide.Matches(m.Text) {
    ws.SendMessage("You got it boss, nighty night.", m.Channel)
    os.Exit(0)
  }
}

func (c Debug) debugAttempts(m util.IncomingSlackMessage) {
  attempts := util.NewRegex("jarvis debug attempts")
  if attempts.Matches(m.Text) {
    resp := "I've recorded these attempts at debug access since I was last started:"
    for _, attempt := range DebugAccessAttempts {
      resp += "\n" + attempt
    }
    ws.SendMessage(resp, m.Channel)
  }
}

func (c Debug) debugDumpRedisKeys(m util.IncomingSlackMessage) {
  reg := util.NewRegex("jarvis debug dump redis")
  if reg.Matches(m.Text) {
    keys := data.Keys("*")
    resp := "I'm currently tracking these keys in storage:\n"
    for _, key := range keys {
      _, v := data.Get(key)
      resp += key + ": " + v + "\n"
    }
    ws.SendMessage(resp, m.Channel)
  }
}
