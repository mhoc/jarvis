
package commands

import (
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/service"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
  "regexp"
)

type Status struct {}

func (s Status) Matches() []*regexp.Regexp {
  return []*regexp.Regexp{
    regexp.MustCompile("status"),
    regexp.MustCompile("alive"),
  }
}

func (s Status) Help(m util.IncomingSlackMessage) {
  message := util.HelpGenerator{
    CommandName: "status",
    Description: "prints status information about the jarvis runtime, including the running version and location",
    RegexMatches: s.Matches(),
    Format: "jarvis (match)",
    Examples: []string{"jarvis status"},
    OtherTopics: []util.HelpGeneratorTopic{},
  }.Generate()
  ws.SendMessage(message, m.Channel)
}

func (s Status) Execute(m util.IncomingSlackMessage) {
  response := "Jarvis, at your service.\n"
  version := service.Git{}.LastTag()
  response += "I'm running version " + version
  location := config.Location()
  response += " on " + location + "."
  ws.SendMessage(response, m.Channel)
}
