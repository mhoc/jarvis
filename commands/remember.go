
package commands

import (
  "fmt"
  "github.com/mhoc/jarvis/data"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
  "regexp"
  "strings"
)

type (
  Remember struct {}
  RememberKey struct {
    Key string
    Aliases []string
  }
)

var (
  RememberKeys = []RememberKey{
    RememberKey{
      Key: "user-zipcode-",
      Aliases: []string{"zipcode", "zip code", "zip"},
    },
  }
)

func (r Remember) Matches() []*regexp.Regexp {
  return []*regexp.Regexp {
    regexp.MustCompile("remember"),
  }
}

func (r Remember) Help(m util.IncomingSlackMessage) {
  message := util.HelpGenerator{
    CommandName: "remember",
    Description: "instructs jarvis to commit some piece of information to memory.\ninformation has categories which the commit has to reference.\nthus you cant have jarvis remember arbitrary data. rather, only data which jarvis is configured to remember.",
    RegexMatches: r.Matches(),
    Format: "jarvis (match) .* (data key) is (data value)",
    Examples: []string{"jarvis remember that my zip code is 46723", "jarvis know that my birthday is march 11 1993"},
    OtherTopics: []util.HelpGeneratorTopic{
      util.HelpGeneratorTopic{
        Name: "data keys",
        Body: "zipcode\nbirthday",
      },
    },
  }.Generate()
  ws.SendMessage(message, m.Channel)
}

func (r Remember) Execute(m util.IncomingSlackMessage) {
  key, datum := r.GetRememberKeyDatum(m.Text)
  if key == nil {
    ws.SendMessage("I don't recognize the type of data you're asking me to remember.", m.Channel)
    return
  }
  data.Cache(key.Key + m.User, datum)
  message := fmt.Sprintf("Alright. I'll remember that your %v is '%v'.", key.Aliases[0], datum)
  ws.SendMessage(message, m.Channel)
}

func (r Remember) GetRememberKeyDatum(text string) (*RememberKey, string) {
  for _, key := range RememberKeys {
    for _, alias := range key.Aliases {
      alias += " is "
      index := strings.Index(text, alias)
      if index > -1 {
        text = text[index+len(alias):]
        return &key, text
      }
    }
  }
  return nil, ""
}
