
package commands

import (
  "github.com/jbrukh/bayesian"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
)

type Know struct {}
const KnowClass bayesian.Class = "know"

func (k Know) Class() bayesian.Class {
  return KnowClass
}

func (k Know) TrainingStrings() []string {
  return []string{
    "know that",
  }
}

func (k Know) Description() string {
  return "Tells jarvis some very helpful information about yourself."
}

func (k Know) Execute(m util.IncomingSlackMessage) {
  response := "My help functionality is a bit underdeveloped at the moment.\n"
  response += "Check out github.com/mhoc/jarvis for more information."
  ws.SendMessage(response, m.Channel)
}
