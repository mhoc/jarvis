
package commands

import (
  "github.com/jbrukh/bayesian"
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/service"
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
  response := "Jarvis, at your service.\n"
  version := service.Git().LastTag()
  response += "I'm running version " + version
  location := config.Location()
  response += " on " + location + "."
  ws.SendMessage(response, m.Channel)
}
