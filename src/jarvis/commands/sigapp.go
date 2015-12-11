
// These are commands specific to my SIGAPP club.
// Other users of jarvis would probably want to remove these commands as
// they wont be all that useful
package commands

import (
  "jarvis/ws"
  "jarvis/util"
  "math/rand"
)

const (
  UltronId = "U08013DQ8"
)

var (
  UltronPurgeMessages = []string{
    "I wish I could.",
  }
)

type SIGAPP struct {}

func NewSIGAPP() SIGAPP {
  return SIGAPP{}
}

func (s SIGAPP) Name() string {
  return "sigapp"
}

func (s SIGAPP) Description() string {
  return "commands to provide functionality specific to Purdue ACM SIGAPP"
}

func (s SIGAPP) Examples() []string {
  return []string{"jarvis purge ultron"}
}

func (s SIGAPP) OtherDocs() []util.HelpTopic {
  return []util.HelpTopic{}
}

func (s SIGAPP) SubCommands() []util.SubCommand {
  return []util.SubCommand{
    util.NewSubCommand("^jarvis purge ultron( from the net)?$", s.UltronPurge),
  }
}

func (s SIGAPP) UltronPurge(m util.IncomingSlackMessage, r util.Regex) {
  msg := UltronPurgeMessages[rand.Intn(len(UltronPurgeMessages))]
  ws.SendMessage(msg, m.Channel)
}
