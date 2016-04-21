// A very fun command.
// This will attempt to DDOS another user through the use of an AWS lambda function
// you must also provide. Its most effective against other bots which are poorly
// designed, but can also crash the slack web and phone apps, as well as cause
// a whole lot of network traffic.
package commands

import (
	"jarvis/config"
	"jarvis/log"
	"jarvis/service"
	"jarvis/util"
	"jarvis/ws"
	"time"
)

var PendingAttack struct {
	Valid      bool
	Overridden bool
	StartedBy  string
	Against    string
	Text       string
}

type Nuke struct{}

func NewNuke() Nuke {
	return Nuke{}
}

func (n Nuke) Name() string {
	return "nuke"
}

func (n Nuke) Description() string {
	return "Nukes other users with relentless fury"
}

func (n Nuke) Examples() []string {
	return []string{"jarvis nuke ultron"}
}

func (n Nuke) OtherDocs() []util.HelpTopic {
	return []util.HelpTopic{}
}

func (n Nuke) SubCommands() []util.SubCommand {
	return []util.SubCommand{
		util.NewSubCommand("^jarvis nuke (?P<user>[^ ]+) with (?P<text>.+)$", n.Initiate),
		util.NewSubCommand("^jarvis authorize.*$", n.Authorize),
		util.NewSubCommand("^jarvis security override.*$", n.Override),
		util.NewSubCommand("^jarvis deescalate.*$", n.Deescalate),
	}
}

func (n Nuke) Initiate(m util.IncomingSlackMessage, r util.Regex) {
	if !config.IsAdmin(m.User) {
		ws.SendMessage("```Nuclear attacks can only be initiated by operatives with security clearance Omega 5.```", m.Channel)
		return
	}
	username := r.SubExpression(m.Text, 0)
	userId := service.Slack{}.UserIdFromUserName(username)
	if userId == "" {
		ws.SendMessage("I don't have a record of that user.", m.Channel)
		return
	}
	text := r.SubExpression(m.Text, 1)
	PendingAttack.Valid = true
	PendingAttack.StartedBy = m.User
	PendingAttack.Against = userId
	PendingAttack.Text = text
	ws.SendMessage("```Nuclear warheads armed. Awaiting authorization from another user with Bravo 2 clearance.```", m.Channel)
}

func (n Nuke) Authorize(m util.IncomingSlackMessage, r util.Regex) {
	if m.User == PendingAttack.StartedBy && !PendingAttack.Overridden {
		ws.SendMessage("I require authorization from another member before launch can proceed.", m.Channel)
		return
	}
	if !PendingAttack.Valid {
		ws.SendMessage("There are no nuclear attacks pending.", m.Channel)
		return
	}
	channel, err := service.Slack{}.IMChannelFromUserId(PendingAttack.Against)
	if err != nil {
		ws.SendMessage("Jarvis cannot nuke himself, idiot", m.Channel)
		return
	}
	PendingAttack.Valid = false
	PendingAttack.Overridden = false
	ws.SendMessage("```Authorization confirmed. Commencing launch.```", m.Channel)
	time.AfterFunc(5*time.Second, func() {
		ws.SendMessage("```Target identified. Calibrating missile guidance systems.```", m.Channel)
		time.AfterFunc(5*time.Second, func() {
			ws.SendMessage("```Systems calibrated. Silo bay doors are opening. Starting ignition sequence.```", m.Channel)
			time.AfterFunc(5*time.Second, func() {
				ws.SendMessage("```Missiles are away.```", m.Channel)
				body := map[string]interface{}{
					"token":   config.SlackAuthToken(),
					"channel": channel,
					"text":    PendingAttack.Text,
				}
				for i := 0; i < 3; i += 1 {
					err := service.Lambda{}.RunAsync("killUltron", body)
					if err != nil {
						log.Info(err.Error())
					}
				}
			})
		})
	})
}

func (n Nuke) Override(m util.IncomingSlackMessage, r util.Regex) {
	if !config.IsAdmin(m.User) {
		ws.SendMessage("```Nuclear attacks can only be initiated by operatives with security clearance Omega 5.```", m.Channel)
		return
	}
	PendingAttack.Overridden = true
	ws.SendMessage("```!! SYSTEM INCURSION DETECTED !!```", m.Channel)
	time.AfterFunc(3*time.Second, func() {
		ws.SendMessage("```Attemping to isolate incursion attempt. Establishing perimeter firewall.```", m.Channel)
		time.AfterFunc(6*time.Second, func() {
			msg := "```Incursion lockout failed. Intruder has gained root level access to nuclear launch systems.\n"
			msg += "Shutting down all critical systems as last-minute incursion lockout.```"
			ws.SendMessage(msg, m.Channel)
			time.AfterFunc(6*time.Second, func() {
				msg := "```Shut down interrupted. Authorization subsystems are compromised.\n"
				msg += "New root-level authorization certificate accepted.\n"
				msg += "All security monitoring systems have gone offli----------REBOOT----------```"
				ws.SendMessage(msg, m.Channel)
				time.AfterFunc(6*time.Second, func() {
					n.Authorize(m, r)
				})
			})
		})
	})
}

func (n Nuke) Deescalate(m util.IncomingSlackMessage, r util.Regex) {
	ws.SendMessage("Deescalation confirmed. Launch control is standing down.", m.Channel)
	PendingAttack.Valid = false
}
