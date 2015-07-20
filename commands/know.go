
package commands

import (
  "fmt"
  "github.com/jbrukh/bayesian"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
  "strings"
)

type Know struct {}
const KnowClass bayesian.Class = "know"

func (k Know) Class() bayesian.Class {
  return KnowClass
}

func (k Know) TrainingStrings() []string {
  return []string{
    "know",
  }
}

func (k Know) Description() string {
  return "Tells jarvis some very helpful information about yourself."
}

func (k Know) Execute(m util.IncomingSlackMessage) {
  knowSplit := strings.Split(m.Text, "is ")
  fmt.Printf("%v\n", knowSplit)
  if len(knowSplit) != 3 {
    ws.SendMessage("Please put what you want me to know after the word 'is' in your request.", m.Channel)
    return
  }

  knowWhat := knowSplit[2]
  var knowThat string

  if strings.Contains(m.Text, "zip code") || strings.Contains(m.Text, "zipcode") {
    knowThat = "zip code"
  } else {
    ws.SendMessage("I'm not totally sure what you're trying to tell me to know.", m.Channel)
    return
  }
  ws.SendMessage(fmt.Sprintf("No problem. I now know that your %v is %v.", knowThat, knowWhat), m.Channel)

}
