package commands

import (
	"jarvis/config"
	"jarvis/service"
	"jarvis/util"
	"jarvis/ws"
)

type Status struct{}

func NewStatus() Status {
	return Status{}
}

func (s Status) Name() string {
	return "status"
}

func (s Status) Description() string {
	return "prints status information about the jarvis runtime, including the running version and location"
}

func (s Status) Examples() []string {
	return []string{"jarvis status"}
}

func (s Status) OtherDocs() []util.HelpTopic {
	return []util.HelpTopic{}
}

func (s Status) SubCommands() []util.SubCommand {
	return []util.SubCommand{
		util.NewSubCommand("^jarvis status$", s.Report),
		util.NewSubCommand("^jarvis are you alive\\??$", s.Confirm),
		util.NewSubCommand("^jarvis are you awake\\??$", s.Confirm),
		util.NewSubCommand("^jarvis are you there\\??$", s.Confirm),
		util.NewSubCommand("^jarvis are you dead\\?$", s.Deny),
	}
}

func (s Status) Report(m util.IncomingSlackMessage, r util.Regex) {
	response := "Jarvis, at your service.\n"
	version := service.Git{}.LastTag()
	response += "I'm running version " + version
	location := config.Location()
	response += " on " + location + ".\n"
	response += "I have been alive for " + service.Time{}.DurationToString(config.Uptime()) + "."
	ws.SendMessage(response, m.Channel)
}

func (s Status) Confirm(m util.IncomingSlackMessage, r util.Regex) {
	response := "Absolutely.\n"
	response += "I have been alive for " + service.Time{}.DurationToString(config.Uptime()) + "."
	ws.SendMessage(response, m.Channel)
}

func (s Status) Deny(m util.IncomingSlackMessage, r util.Regex) {
	response := "Of course not.\n"
	response += "I have been alive for " + service.Time{}.DurationToString(config.Uptime()) + "."
	ws.SendMessage(response, m.Channel)
}
