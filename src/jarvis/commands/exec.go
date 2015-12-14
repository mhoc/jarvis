
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
  if err != nil {
    ws.SendMessage(err.Error(), m.Channel)
  } else {
    msg := "Result:\n```"
    msg += result + "\n```"
    ws.SendMessage(msg, m.Channel)
  }
}
