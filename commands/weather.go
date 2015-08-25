
package commands

import (
  "github.com/mhoc/jarvis/data"
  "github.com/mhoc/jarvis/service"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
)

type Weather struct {}

func NewWeather() Weather {
  return Weather{}
}

func (w Weather) Name() string {
  return "weather"
}

func (w Weather) Matches() []util.Regex {
  return []util.Regex{
    util.NewRegex("weather"),
  }
}

func (w Weather) Description() string {
  return "provides current weather and weather forcasts through the darksky weather api.\n you can use the 'remember' command to give jarvis your zipcode and it will use it if you don't provide a zipcode in the command."
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
  // Different types of requests to the weather api match different regexes
  noZipCodeProvided := util.NewRegex("^jarvis weather$")
  zipCodeProvided := util.NewRegex("^jarvis weather ([0-9]{5})$")
  if noZipCodeProvided.Matches(m.Text) {
    w.noZipCodeProvided(m)
  } else if zipCodeProvided.Matches(m.Text) {
    w.zipCodeProvided(zipCodeProvided.SubExpression(m.Text, 0), m)
  } else {
    ws.SendMessage("I don't recognize that type of weather request. Sorry :(", m.Channel)
  }
}

func (w Weather) noZipCodeProvided(m util.IncomingSlackMessage) {
  in, storedZipCode := data.GetDatum("my zip", m.User)
  if !in {
    ws.SendMessage("You didn't provide a location and I don't have one on file for you. Bummer.", m.Channel)
    return
  }
  weather, err := service.Weather{}.ForecastFriendly(storedZipCode)
  switch err.(type) {
  case service.BadZipCodeError:
    ws.SendMessage("The zip code I have on file for you doesn't appear to be valid. Might want to fix that.", m.Channel)
  case error:
    ws.SendMessage("A fatal error has occured. Don't worry I'm fine.", m.Channel)
  default:
    ws.SendMessage("Here's my forecast for " + storedZipCode + ":\n" + weather, m.Channel)
  }
}

func (w Weather) zipCodeProvided(zipcode string, m util.IncomingSlackMessage) {
  weather, err := service.Weather{}.ForecastFriendly(zipcode)
  switch err.(type) {
  case service.BadZipCodeError:
    ws.SendMessage("It doesn't look like the zipcode you provided is valid. Bummer.", m.Channel)
  case error:
    ws.SendMessage("A really bad error occured during my attempt to predict the future. Of weather. Sorry.", m.Channel)
  default:
    ws.SendMessage(weather, m.Channel)
  }
}
