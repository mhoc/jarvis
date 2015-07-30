
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

func (w Weather) Help(m util.IncomingSlackMessage) {
  message := util.HelpGenerator{
    CommandName: "weather",
    Description: "provides current weather and weather forcasts through the darksky weather api",
    RegexMatches: w.Matches(),
    Format: "jarvis (match) (zipcode)",
    Examples: []string{"jarvis weather 46723"},
    OtherTopics: []util.HelpGeneratorTopic{},
  }.Generate()
  ws.SendMessage(message, m.Channel)
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
