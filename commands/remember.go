
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

func (r Remember) Description() string {
  return "Tells me to remember something. Forever."
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
