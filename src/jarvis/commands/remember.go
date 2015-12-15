// Remember is the command interface into datum storage

package commands

import (
	"fmt"
	"jarvis/data"
	"jarvis/util"
	"jarvis/ws"
	"strings"
)

type Remember struct{}

func NewRemember() Remember {
	return Remember{}
}

func (r Remember) Name() string {
	return "remember"
}

func (r Remember) Description() string {
	return "instructs jarvis to commit some piece of information to memory.\ninformation has categories which the commit has to reference.\nthus you cant have jarvis remember arbitrary data. rather, only data which jarvis is configured to remember."
}

func (r Remember) Examples() []string {
	return []string{"jarvis remember that my zip code is 46723", "jarvis know that my birthday is march 11 1993"}
}

func (r Remember) OtherDocs() []util.HelpTopic {
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

func (r Remember) SubCommands() []util.SubCommand {
	return []util.SubCommand{
		util.NewSubCommand("^jarvis remember (that)? (?P<key>.+) is (?P<value>.+)$", r.Save),
		util.NewSubCommand("^jarvis know that (?P<key>.+) is (?P<value>.+)$", r.Save),
	}
}

func (r Remember) Save(m util.IncomingSlackMessage, regex util.Regex) {
	var key, value string
	if regex.NSubExpressions(m.Text) == 2 {
		key, value = regex.SubExpression(m.Text, 0), regex.SubExpression(m.Text, 1)
	} else {
		key, value = regex.SubExpression(m.Text, 1), regex.SubExpression(m.Text, 2)
	}
	err := data.StoreDatum(key, value, m.User)
	if err != nil {
		ws.SendMessage(err.Error(), m.Channel)
		return
	}
	key = strings.Replace(key, "my", "your", -1)
	ws.SendMessage(fmt.Sprintf("Alright. I'll remember that %v is %v.", key, value), m.Channel)
}
