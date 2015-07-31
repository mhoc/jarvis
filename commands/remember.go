
// Remember is the command interface into datum storage

package commands

import (
  "fmt"
  "github.com/mhoc/jarvis/data"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
  "regexp"
  "strings"
)

type Remember struct {}

func (r Remember) Name() string {
  return "remember"
}

func (r Remember) Matches() []*regexp.Regexp {
  return []*regexp.Regexp {
    regexp.MustCompile("remember"),
    regexp.MustCompile("know"),
  }
}

func (r Remember) Description() string {
  return "instructs jarvis to commit some piece of information to memory.\ninformation has categories which the commit has to reference.\nthus you cant have jarvis remember arbitrary data. rather, only data which jarvis is configured to remember."
}

func (r Remember) Format() string {
  return "jarvis (match) that (data key) is (data value)"
}

func (r Remember) Examples() []string {
  return []string{"jarvis remember that my zip code is 46723", "jarvis know that my birthday is march 11 1993"}
}

func (r Remember) OtherDocs() []util.HelpTopic {
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

func (r Remember) Execute(m util.IncomingSlackMessage) {
  regex := util.NewRegex("that (.+) is (.+)")
  if !regex.Matches(m.Text) {
    ws.SendMessage("I can't parse your query. I'm sorry I can't live up to your expectations daddy.", m.Channel)
    return
  }
  key, value := regex.SubExpression(m.Text, 0), regex.SubExpression(m.Text, 1)
  success := data.StoreDatum(key, value, m.User)
  if !success {
    ws.SendMessage("I don't recognize the type of data you're asking me to remember.", m.Channel)
    return
  }
  key = strings.Replace(key, "my", "your", -1)
  ws.SendMessage(fmt.Sprintf("Alright. I'll remember that %v is %v.", key, value), m.Channel)
}
