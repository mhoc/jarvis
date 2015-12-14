
package commands

import (
  "jarvis/util"
)

type MyCommand struct {}

func NewMyCommand() MyCommand {
  return MyCommand{}
}

func (c MyCommand) Name() string {
  return "mycommand"
}

func (c MyCommand) Description() string {
  return ""
}

func (c MyCommand) Examples() []string {
  return []string{"jarvis mycommand"}
}

func (c MyCommand) OtherDocs() []util.HelpTopic {
  return []util.HelpTopic{}
}

func (c MyCommand) SubCommands() []util.SubCommand {
  return []util.SubCommand{
    util.NewSubCommand("^jarvis mycommand$", c.Thing),
  }
}

func (c MyCommand) Thing(m util.IncomingSlackMessage, r util.Regex) {
  
}
