
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

func (r Recall) Name() string {
  return "recall"
}

func (r Recall) Matches() []util.Regex {
  return []util.Regex{
    util.NewRegex("recall"),
    util.NewRegex("get"),
    util.NewRegex("what is"),
  }
}

func (r Recall) Description() string {
  return "instructs jarvis to recall some piece of information which he has already stored."
}

func (r Recall) Format() string {
  return "jarvis (match) (data key)"
}

func (r Recall) Examples() []string {
  return []string{"jarvis recall my zip code", "jarvis get my zip code", "jarvis what is my birthday"}
}

func (r Recall) OtherDocs() []util.HelpTopic {
  var keys string
  for _, datum := range data.RegisteredDatums {
    keys += "  " + datum.Aliases[0] + "\n"
  }
  keys = keys[:len(keys)-1]
  return []util.HelpTopic{
    util.HelpTopic{
      Name: "data keys",
      Body: keys,
    },
  }
}

func (r Recall) Execute(m util.IncomingSlackMessage) {
  regex := util.NewRegex("jarvis (recall|get|what is) ([A-Za-z0-9 ]+)")
  if !regex.Matches(m.Text) {
    ws.SendMessage("My appologies but I can't seem to parse your query.", m.Channel)
    return
  }
  key := regex.SubExpression(m.Text, 1)
  in, data := data.GetDatum(key, m.User)
  key = strings.Replace(key, "my", "your", -1)
  if !in || data == "" {
    ws.SendMessage(fmt.Sprintf("I can't seem to remember what %v is.", key), m.Channel)
    return
  }
  ws.SendMessage(fmt.Sprintf("Looks like %v is %v.", key, data), m.Channel)
}
