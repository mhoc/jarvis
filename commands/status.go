
package commands

import (
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/service"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
)

type Status struct {}

func NewStatus() Status {
  return Status{}
}

func (s Status) Name() string {
  return "status"
}

func (s Status) Description() string {
  return "prints status information about the jarvis runtime, including the running version and location"
}

func (s Status) Examples() []string {
  return []string{"jarvis status"}
}

func (s Status) OtherDocs() []util.HelpTopic {
  return []util.HelpTopic{}
}

func (s Status) SubCommands() []util.SubCommand {
  return []util.SubCommand{
    util.NewSubCommand("^jarvis status$", s.Report),
    util.NewSubCommand("^jarvis are you alive$", s.Report),
    util.NewSubCommand("^jarvis are you awake$", s.Report),
  }
}

func (s Status) Report(m util.IncomingSlackMessage, r util.Regex) {
  response := "Jarvis, at your service.\n"
  version := service.Git{}.LastTag()
  response += "I'm running version " + version
  location := config.Location()
  response += " on " + location + "."
  ws.SendMessage(response, m.Channel)
}
