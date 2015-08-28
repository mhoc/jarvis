
// Recall is the retrieval interface into datum storage

package commands

import (
  "fmt"
  "github.com/mhoc/jarvis/data"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
  "strings"
)

type Recall struct {}

func NewRecall() Recall {
  return Recall{}
}

func (r Recall) Name() string {
  return "recall"
}

func (r Recall) Description() string {
  return "instructs jarvis to recall some piece of information which he has already stored."
}

func (r Recall) Examples() []string {
  return []string{"jarvis recall my zip code", "jarvis get my zip code", "jarvis what is my birthday"}
}

func (r Recall) OtherDocs() []util.HelpTopic {
  var keys string
  for _, datum := range data.RegisteredDatums {
    keys += datum.Aliases[0] + "\n"
  }
  keys = keys[:len(keys)-1]
  return []util.HelpTopic{
    util.HelpTopic{
      Name: "data keys",
      Body: keys,
    },
  }
}

func (r Recall) SubCommands() []util.SubCommand {
  return []util.SubCommand{
    util.NewSubCommand("^jarvis recall (?P<key>.+)$", r.Get),
    util.NewSubCommand("^jarvis get (?P<key>.+)$", r.Get),
    util.NewSubCommand("^jarvis what is (?P<key>.+)$", r.Get),
  }
}

func (r Recall) Get(m util.IncomingSlackMessage, reg util.Regex) {
  key := reg.SubExpression(m.Text, 0)
  in, data := data.GetDatum(key, m.User)
  key = strings.Replace(key, "my", "your", -1)
  if !in || data == "" {
    ws.SendMessage(fmt.Sprintf("I can't seem to remember what %v is.", key), m.Channel)
    return
  }
  ws.SendMessage(fmt.Sprintf("Looks like %v is %v.", key, data), m.Channel)
}
