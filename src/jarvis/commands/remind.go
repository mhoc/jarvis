
package commands

import (
  "fmt"
  "jarvis/log"
  "jarvis/service"
  "jarvis/util"
  "jarvis/ws"
  "time"
)

var (
  PendingReminders = make([]Reminder, 0)
)

type Reminder struct {
  Id string
  TargetId string
  TargetName string
  ToDo string
  OnChannel string
  At time.Time
}

func NewReminder(targetId string, targetname string, todo string, onchannel string, at time.Time) Reminder {
  return Reminder{
    Id: util.NewId(),
    TargetId: targetId,
    TargetName: targetname,
    ToDo: todo,
    OnChannel: onchannel,
    At: at,
  }
}

func (r Reminder) Start() {
  time.AfterFunc(r.At.Sub(time.Now()), r.Send)
}

func (r Reminder) Send() {
  deleteIndex := -1
  for i, reminder := range PendingReminders {
    if reminder.Id == r.Id {
      deleteIndex = i
      break
    }
  }
  // TODO: This is probably a race condition
  PendingReminders = append(PendingReminders[:deleteIndex], PendingReminders[deleteIndex+1:]...)
  ws.SendMessage(r.TargetName + ", you asked me to remind you to " + r.ToDo + ".", r.OnChannel)
  ws.SendPrivateMessage("Hey there: Don't forget to " + r.ToDo + ".", r.TargetId)
}

func (r Reminder) String() string {
  s := fmt.Sprintf("At %v on %v, %v will be reminded to %v", r.At.Format("15:04:15"), r.At.Format("Jan 2"), r.TargetName, r.ToDo)
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
if jarvis is reset, all pending reminders will be lost
setting reminders by duration can be done in units of hours, minutes, and seconds.`
}

func (c Remind) Examples() []string {
  return []string{"jarvis remind me in 10 minutes to take out the garbage"}
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
  username := service.Slack{}.UserNameFromUserId(m.User)
  durStr, note := r.SubExpression(m.Text, 0), r.SubExpression(m.Text, 1)
  actDur, err := service.Time{}.StringToDuration(durStr)
  if err != nil {
    log.Trace("Incorrect duration string")
    ws.SendMessage(err.Error(), m.Channel)
    return
  }

  // Create and start the reminder
  rem := NewReminder(m.User, username, note, m.Channel, time.Now().Add(actDur))
  rem.Start()

  // Cache the reminder in our list of pending reminders
  PendingReminders = append(PendingReminders, rem)
  ws.SendMessage("Alright. I'll remind you in " + service.Time{}.DurationToString(actDur) + " to " + note, m.Channel)
}

func (c Remind) SetAbsoluteReminder(m util.IncomingSlackMessage, r util.Regex) {
  // Parse absolute time string
  username := service.Slack{}.UserNameFromUserId(m.User)
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
  rem := NewReminder(m.User, username, note, m.Channel, t)
  rem.Start()

  // Cache the reminder
  PendingReminders = append(PendingReminders, rem)
  ws.SendMessage("Alright. I'll remind you in " + service.Time{}.DurationToString(t.Sub(time.Now())) + " to " + note, m.Channel)
}


func (c Remind) ListReminders(m util.IncomingSlackMessage, r util.Regex) {
  resp := ""
  for _, reminder := range PendingReminders {
    resp += reminder.String() + "\n"
  }
  if len(resp) == 0 {
    resp += "I'm not currently tracking any reminders."
  }
  ws.SendMessage(resp, m.Channel)
}
