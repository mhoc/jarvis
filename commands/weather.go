
package commands

import (
  "github.com/mhoc/jarvis/service"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
  "regexp"
)

type Weather struct {}

func (w Weather) Matches() []*regexp.Regexp {
  return []*regexp.Regexp{
    regexp.MustCompile("weather"),
    regexp.MustCompile("rain"),
  }
}

func (w Weather) Description() string {
  return "Your local weatherman."
}

func (w Weather) Execute(m util.IncomingSlackMessage) {
  zipCodeRegex := regexp.MustCompile("[0-9]{5}")
  zipCode := string(zipCodeRegex.Find([]byte(m.Text)))
  if zipCode == "" {
    ws.SendMessage("You should probably provide a zipcode", m.Channel)
  } else {
    weather := service.Weather{}.ForecastFriendly(zipCode)
    ws.SendMessage(weather, m.Channel)
  }
}
