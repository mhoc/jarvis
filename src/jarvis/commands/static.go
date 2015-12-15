package commands

import (
	"jarvis/config"
	"jarvis/util"
	"jarvis/ws"
)

type Static struct {
	Topics       map[string]string
	UserCommands []util.SubCommand
}

func NewStatic() Static {
	s := Static{}
	rawTopics := config.Static()
	actTopics := make(map[string]string)
	for _, t := range rawTopics {
		actTopics[t["key"].(string)] = t["value"].(string)
	}
	userCmds := make([]util.SubCommand, 0, 0)
	for key, _ := range actTopics {
		userCmds = append(userCmds, util.NewSubCommand("^jarvis "+key+"$", s.Exec))
	}
	s.Topics = actTopics
	s.UserCommands = userCmds
	return s
}

func (c Static) Name() string {
	return "static"
}

func (c Static) Description() string {
	return "provides the ability for jarvis admins to provide easily accessible static data to users based on certain keywords"
}

func (c Static) Examples() []string {
	if len(c.UserCommands) > 0 {
		return []string{
			"see 'matches on' above",
		}
	}
	return []string{}
}

func (c Static) OtherDocs() []util.HelpTopic {
	return []util.HelpTopic{}
}

func (c Static) SubCommands() []util.SubCommand {
	return c.UserCommands
}

func (c Static) Exec(m util.IncomingSlackMessage, r util.Regex) {
	ws.SendMessage(c.Topics[r.String()], m.Channel)
}
