
package commands

import (
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
  "regexp"
)

type Help struct {}

func (h Help) Matches() []*regexp.Regexp {
  return []*regexp.Regexp {
    regexp.MustCompile("help"),
  }
}

func (h Help) Description() string {
  return "Prints some very helpful help text."
}

func (h Help) Execute(m util.IncomingSlackMessage) {
  response := "My help functionality is a bit underdeveloped at the moment.\n"
  response += "Check out github.com/mhoc/jarvis for more information."
  ws.SendMessage(response, m.Channel)
}
