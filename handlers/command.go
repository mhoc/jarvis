
package handlers

import (
  "github.com/jbrukh/bayesian"
  "github.com/mhoc/jarvis/commands"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
  "strings"
)

var commandManifest = []util.Command{
  commands.Help{},
  commands.Status{},
}

var cmdCh = make(chan util.IncomingSlackMessage)
var Classifier *bayesian.Classifier

func InitCommands() {
  log.Info("Initing command listener")
  ws.SubscribeToMessages(cmdCh)
  TrainCommandClassifier()
  go BeginCommandLoop()
}

func TrainCommandClassifier() {
  classes := []bayesian.Class{}
  for _, command := range commandManifest {
    classes = append(classes, command.Class())
  }
  Classifier = bayesian.NewClassifier(classes...)
  for _, command := range commandManifest{
    Classifier.Learn(command.TrainingStrings(), command.Class())
  }
}

func BeginCommandLoop() {
  for {
    msg := <-cmdCh
    if !IsCommand(msg.Text) {
      continue
    }
    cmd := MatchCommand(msg.Text)
    if cmd != nil {
      go cmd.Execute(msg)
    }
  }
}

func IsCommand(text string) bool {
  if strings.Contains(text, "jarvis") {
    return true
  }
  if strings.Contains(text, "Jarvis") {
    return true
  }
  return false
}

func MatchCommand(text string) util.Command {
  _, likely, _ := Classifier.ProbScores(strings.Split(text, " "))
  return commandManifest[likely]
}
