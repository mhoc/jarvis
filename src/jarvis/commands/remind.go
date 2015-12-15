package commands

import (
	"encoding/json"
	"fmt"
	"jarvis/data"
	"jarvis/log"
	"jarvis/service"
	"jarvis/util"
	"jarvis/ws"
	"time"
)

type Reminder struct {
	Id        string
	TargetId  string
	ToDo      string
	OnChannel string
	At        time.Time
}

func NewReminder(targetId string, todo string, onchannel string, at time.Time) Reminder {
	return Reminder{
		Id:        util.NewId(),
		TargetId:  targetId,
		ToDo:      todo,
		OnChannel: onchannel,
		At:        at,
	}
}

func ReminderFromJSON(js string) Reminder {
	var r Reminder
	json.Unmarshal([]byte(js), &r)
	return r
}

func (r Reminder) RedisKey() string {
	return fmt.Sprintf("reminder-%v-%v", r.TargetId, r.Id)
}

func (r Reminder) ToJSON() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r Reminder) Send() {
	name := service.Slack{}.UserNameFromUserId(r.TargetId)
	ws.SendMessage(name+", you asked me to remind you to "+r.ToDo+".", r.OnChannel)
	ws.SendPrivateMessage("Hey there: Don't forget to "+r.ToDo+".", r.TargetId)
}

func (r Reminder) String() string {
	name := service.Slack{}.UserNameFromUserId(r.TargetId)
	s := fmt.Sprintf("At %v on %v, %v will be reminded to %v", r.At.Format("15:04:15"), r.At.Format("Jan 2"), name, r.ToDo)
	return s
}

type Remind struct{}

func NewRemind() Remind {
	return Remind{}
}

func (c Remind) Name() string {
	return "remind"
}

func (c Remind) Description() string {
	return `allows you to set and view reminders
setting reminders by duration can be done in units of hours, minutes, and seconds.
reminders can be expected to be sent out within 10 seconds of the scheduled time.
one user setting more than 25 reminders will cause jarvis to nuke all of them.`
}

func (c Remind) Examples() []string {
	return []string{
		"jarvis remind me in 10 minutes to take out the garbage",
		"jarvis remind me at 8am to wake up",
	}
}

func (c Remind) OtherDocs() []util.HelpTopic {
	return []util.HelpTopic{}
}

func (c Remind) SubCommands() []util.SubCommand {
	return []util.SubCommand{
		util.NewSubCommand("^jarvis remind me in (?P<duration>.+) to (?P<note>.+)$", c.SetDurationReminder),
		util.NewSubCommand("^jarvis remind me at (?P<time>.+) to (?P<note>.+)$", c.SetAbsoluteReminder),
		util.NewSubCommand("^jarvis list reminders$", c.ListReminders),
		util.NewSubCommand("^jarvis what reminders are you tracking\\??", c.ListReminders),
	}
}

func (c Remind) SetDurationReminder(m util.IncomingSlackMessage, r util.Regex) {
	// Parse the user's duration string
	durStr, note := r.SubExpression(m.Text, 0), r.SubExpression(m.Text, 1)
	actDur, err := service.Time{}.StringToDuration(durStr)
	if err != nil {
		log.Trace("Incorrect duration string")
		ws.SendMessage(err.Error(), m.Channel)
		return
	}

	// Create and put the reminder in redis
	rem := NewReminder(m.User, note, m.Channel, time.Now().Add(actDur))
	data.Set(rem.RedisKey(), rem.ToJSON())
	ws.SendMessage("Alright. I'll remind you in "+service.Time{}.DurationToString(actDur)+" to "+note, m.Channel)
}

func (c Remind) SetAbsoluteReminder(m util.IncomingSlackMessage, r util.Regex) {
	// Parse absolute time string
	absTimeString, note := r.SubExpression(m.Text, 0), r.SubExpression(m.Text, 1)
	t, err := service.Time{}.StringToTime(absTimeString, m.User)
	if err != nil {
		log.Trace("Incorrect absolute time string")
		ws.SendMessage(err.Error(), m.Channel)
		return
	}

	// Check to make sure the time they entered is after the current time
	if time.Now().After(t) {
		ws.SendMessage("Unfortunately the time you entered exists in the past and I can't travel through time.", m.Channel)
		return
	}

	// Send it
	rem := NewReminder(m.User, note, m.Channel, t)
	data.Set(rem.RedisKey(), rem.ToJSON())
	ws.SendMessage("Alright. I'll remind you in "+service.Time{}.DurationToString(t.Sub(time.Now()))+" to "+note, m.Channel)
}

func (c Remind) ListReminders(m util.IncomingSlackMessage, r util.Regex) {
	resp := "I'm tracking the following reminders for you:\n"
	reminderKeys := data.Keys(fmt.Sprintf("reminder-%v-*", m.User))
	for _, remKey := range reminderKeys {
		_, jsonString := data.Get(remKey)
		resp += ReminderFromJSON(jsonString).String() + "\n"
	}
	if len(resp) == 0 {
		resp += "I'm not currently tracking any reminders."
	}
	ws.SendMessage(resp, m.Channel)
}

// This loop reads all the reminders from Redis and
func StartReminderLoop() {
	go func() {
		log.Info("Starting redis read loop on reminders")
		ticker := time.Tick(10 * time.Second)
		for range ticker {
			reminderKeys := data.Keys("reminder-*")
			for _, remKey := range reminderKeys {
				_, jsonString := data.Get(remKey)
				reminder := ReminderFromJSON(jsonString)
				if time.Now().After(reminder.At) {
					reminder.Send()
					data.Remove(reminder.RedisKey())
				}
			}
		}
	}()
	go func() {
		log.Info("Starting redis rate-limit loop on reminders")
		ticker := time.Tick(60 * time.Second)
		for range ticker {
			userCount := make(map[string][]string)
			reminderKeys := data.Keys("reminder-*")
			for _, remKey := range reminderKeys {
				_, jsonString := data.Get(remKey)
				reminder := ReminderFromJSON(jsonString)
				if _, in := userCount[reminder.TargetId]; in {
					userCount[reminder.TargetId] = append(userCount[reminder.TargetId], reminder.RedisKey())
				} else {
					userCount[reminder.TargetId] = []string{reminder.RedisKey()}
				}
			}
			for user, keys := range userCount {
				if len(keys) > 25 {
					msg := "Sorry to bother you, but it looks like you've scheduled an inordinate number of reminders.\n"
					msg += "Remember: Jarvis is a shared resource and trying to kill me is the act of a murderous psychopath.\n"
					msg += "Your reminders have now been nuked. #sorrynotsorry"
					ws.SendMessage(msg, service.Slack{}.IMChannelFromUserId(user))
					for _, key := range keys {
						data.Remove(key)
					}
				}
			}
		}
	}()
}
