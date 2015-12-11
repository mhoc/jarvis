
// These are commands specific to my SIGAPP club.
// Other users of jarvis would probably want to remove these commands as
// they wont be all that useful
package commands

import (
  "jarvis/config"
  "jarvis/log"
  "jarvis/service"
  "jarvis/ws"
  "jarvis/util"
  "math/rand"
)

const (
  UltronId = "U08013DQ8"
)

var (
  UltronPurgeMessages = []string{
    "With Pleasure.",
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
    util.NewSubCommand("^jarvis,? make ultron sorry he was ever alive$", s.UltronPurge),
  }
}

func (s SIGAPP) UltronPurge(m util.IncomingSlackMessage, r util.Regex) {
  if !config.IsAdmin(m.User) {
    ws.SendMessage("Only The Vision has the power to destroy Ultron.", m.Channel)
    return
  }

  channel := service.Slack{}.IMChannelFromUserId(UltronId)
  // channel := service.Slack{}.IMChannelFromUserId("U02GYT2PN")
  body := map[string]interface{}{
    "token": config.SlackAuthToken(),
    "channel": channel,
    "text": "ultron status",
  }

  for i := 0; i < 10; i += 1 {
    err := service.Lambda{}.RunAsync("killUltron", body)
    if err != nil {
      log.Info(err.Error())
    }
  }

  msg := UltronPurgeMessages[rand.Intn(len(UltronPurgeMessages))]
  ws.SendMessage(msg, m.Channel)

}
