
import sys

name = sys.argv[1]
f = open("./commands/" + name.lower() + ".go", "w+")

content = '''

package commands

import (
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
)

type %s struct {}

func (c %s) Name() string {
  return "%s"
}

func (c %s) Matches() []util.Regex {
  return []util.Regex{
    util.NewRegex("%s"),
  }
}

func (c %s) Description() string {
  return ""
}

func (c %s) Examples() []string {
  return []string{}
}

func (c %s) OtherDocs() []util.HelpTopic {
  return []util.HelpTopic{}
}

func (c %s) Execute(m util.IncomingSlackMessage) {
  ws.SendMessage("Nice!", m.Channel)
}

''' % (
    name, name, name.lower(), name, name.lower(), name, name, name, name
)

f.write(content)
print "Command created!"
print "Be sure to install the command to handlers/command.go when you're finished with it."
