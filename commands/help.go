
package commands

import (
  "github.com/jbrukh/bayesian"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
)

type Help struct {}
const HelpClass bayesian.Class = "help"

func (h Help) Class() bayesian.Class {
  return HelpClass
}

func (h Help) TrainingStrings() []string {
  return []string{
    "help",
  }
}

func (h Help) Description() string {
  return "Prints some very helpful help text."
}

func (h Help) Execute(m util.IncomingSlackMessage) {
  ws.SendMessage("This is help", m.Channel)
}
