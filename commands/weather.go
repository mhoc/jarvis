
package commands

import (
  "github.com/mhoc/jarvis/service"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
  "regexp"
)

type Weather struct {}

func (w Weather) Name() string {
  return "weather"
}

func (w Weather) Matches() []util.Regex {
  return []util.Regex{
    util.NewRegex("weather"),
  }
}

func (w Weather) Description() string {
  return "provides current weather and weather forcasts through the darksky weather api"
}

func (w Weather) Format() string {
  return "jarvis (match) (zipcode)"
}

func (w Weather) Examples() []string {
  return []string{"jarvis weather 46723"}
}

func (w Weather) OtherDocs() []util.HelpTopic {
  return []util.HelpTopic{}
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
