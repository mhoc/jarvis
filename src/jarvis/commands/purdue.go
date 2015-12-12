
// Commands specific to purdue university.
// Other users of jarvis would probably want to disable these commands by removing
// the initialization of the Purdue{} object from handlers/commands.go
package commands

import (
  "jarvis/util"
  // "jarvis/ws"
  // "strings"
)

type Purdue struct {}

func NewPurdue() Purdue {
  return Purdue{}
}

func (r Purdue) Name() string {
  return "purdue"
}

func (r Purdue) Description() string {
  return "commands specific to purdue university."
}

func (r Purdue) Examples() []string {
  return []string{"jarvis dining menu at earhart today"}
}

func (r Purdue) OtherDocs() []util.HelpTopic {
  return []util.HelpTopic{}
}

func (r Purdue) SubCommands() []util.SubCommand {
  return []util.SubCommand{
    util.NewSubCommand("^jarvis dining menus? at (?P<location>[^ ]+) ?(today|tomorrow)?$", r.DiningMenu),
  }
}

func (p Purdue) DiningMenu(m util.IncomingSlackMessage, r util.Regex) {

}
