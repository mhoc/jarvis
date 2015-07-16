
package commands

import (
  "github.com/jbrukh/bayesian"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
)

type Status struct {}
const StatusClass bayesian.Class = "status"

func (s Status) Class() bayesian.Class {
  return StatusClass
}

func (s Status) TrainingStrings() []string {
  return []string{
    "status",
    "alive",
  }
}

func (s Status) Description() string {
  return "Prints my status, with information about my version and where I am hosted."
}

func (s Status) Execute(m util.IncomingSlackMessage) {
  ws.SendMessage("Im here", m.Channel)
}
