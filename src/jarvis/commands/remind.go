
package commands

import (
  "fmt"
  "jarvis/data"
  "jarvis/log"
  "jarvis/service"
  "jarvis/util"
  "jarvis/ws"
  "strings"
  "time"
)

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
however, 'list reminders' might still incorrectly list them as pending, as the list of all pending reminders is cached to persistent storage until they expire.
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
    util.NewSubCommand("^jarvis list reminders$", c.ListReminders),
  }
}

func (c Remind) SetDurationReminder(m util.IncomingSlackMessage, r util.Regex) {
  username := service.Slack{}.UserNameFromUserId(m.User)
  durStr, note := r.SubExpression(m.Text, 0), r.SubExpression(m.Text, 1)
  actDur, err := c.ParseDuration(durStr)
  if err != nil {
    log.Trace("Incorrect duration string: %v", err.Error())
    ws.SendMessage("I can't seem to parse your duration string.", m.Channel)
    return
  }
  time.AfterFunc(actDur, func() {
    ws.SendMessage(username + ", you asked me to remind you to " + note, m.Channel)
  })
  data.SetTimeout(fmt.Sprintf("remind-entry-%v-%v", m.User, time.Now().String()), fmt.Sprintf("Remind %v to %v in %v.", username, note, durStr), actDur)
  ws.SendMessage("Alright. I'll remind you in " + durStr + " to " + note, m.Channel)
}

func (c Remind) ParseDuration(durStr string) (time.Duration, error) {
  // Run the duration string through a bunch of processing to get it into a time.Duration format that can be parsed by go
  durStr = strings.Replace(durStr, " seconds", "s", -1)
  durStr = strings.Replace(durStr, " second", "s", -1)
  durStr = strings.Replace(durStr, " minutes", "m", -1)
  durStr = strings.Replace(durStr, " minute", "m", -1)
  durStr = strings.Replace(durStr, " hours", "h", -1)
  durStr = strings.Replace(durStr, " hour", "h", -1)
  durStr = strings.Replace(durStr, " ", "", -1)
  return time.ParseDuration(durStr)
}

func (c Remind) ListReminders(m util.IncomingSlackMessage, r util.Regex) {
  resp := ""
  reminderKeys := data.Keys("remind-entry-*")
  if len(reminderKeys) == 0 {
    ws.SendMessage("I'm not currently tracking any reminders.", m.Channel)
  } else {
    resp += "I'm currently tracking the following reminders:\n"
    for _, key := range reminderKeys {
      in, val := data.Get(key)
      if in {
        resp += val + "\n"
      }
    }
    ws.SendMessage(resp, m.Channel)
  }
}
