
package commands

import (
  "jarvis/service"
  "jarvis/util"
  "jarvis/ws"
  "time"
)

type Exec struct {}

func NewExec() Exec {
  return Exec{}
}

func (c Exec) Name() string {
  return "exec"
}

func (c Exec) Description() string {
  return "provides a containerized, 'safe' environment to execute arbitrary code."
}

func (c Exec) Examples() []string {
  return []string{
    "jarvis exec python print(1+1)",
    "jarvis exec js console.log('hi there')",
  }
}

func (c Exec) OtherDocs() []util.HelpTopic {
  return []util.HelpTopic{}
}

func (c Exec) SubCommands() []util.SubCommand {
  return []util.SubCommand{
    util.NewSubCommand("^jarvis exec python (?P<command>.+)$", c.Python),
  }
}

func (c Exec) Python(m util.IncomingSlackMessage, r util.Regex) {
  command := r.SubExpression(m.Text, 0)
  result, err := service.Docker{}.RunPythonInContainer(command, 10 * time.Second)
  if len(result) > 0 {
    if result[len(result)-1] == '\n' {
      result = result[:len(result)-1]
    }
  }
  var msg string
  if result != "" {
    msg = "Result: \n```"
    msg += result + "\n```"
    if err != nil {
      msg += "\n"
    }
  }
  if err != nil {
    msg = "Your command resulted in an error:\n```"
    msg += err.Error() + "\n```"
  }
  if len(msg) == 0 {
    msg = "Your command generated no output."
  }
  ws.SendMessage(msg, m.Channel)
}
