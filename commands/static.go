

package commands

import (
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
)

type Static struct {
  Topics map[string]string
  Regexes []util.Regex
}

func NewStatic() Static {
  rawTopics := config.Static()
  actTopics := make(map[string]string)
  for _, t := range rawTopics {
    actTopics[t["key"].(string)] = t["data"].(string)
  }
  regexes := make([]util.Regex, 0, 0)
  for key, _ := range actTopics {
    regexes = append(regexes, util.NewRegex(key))
  }
  return Static{
    Topics: actTopics,
    Regexes: regexes,
  }
}

func (c Static) Name() string {
  return "static"
}

func (c Static) Matches() []util.Regex {
  return c.Regexes
}

func (c Static) Description() string {
  return "provides the ability for jarvis admins to provide easily accessible static data to users based on certain keywords"
}

func (c Static) Format() string {
  return "jarvis (keyword)"
}

func (c Static) Examples() []string {
  if len(c.Regexes) > 0 {
    return []string{
      "jarvis " + c.Regexes[0].String(),
    }
  }
  return []string{}
}

func (c Static) OtherDocs() []util.HelpTopic {
  return []util.HelpTopic{}
}

func (c Static) Execute(m util.IncomingSlackMessage) {
  for _, regex := range c.Regexes {
    if regex.Matches(m.Text) {
      ws.SendMessage(c.Topics[regex.String()], m.Channel)
      return
    }
  }
}
