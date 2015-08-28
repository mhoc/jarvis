
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

func (w Weather) Description() string {
  return "provides current weather and weather forcasts through the darksky weather api.\n you can use the 'remember' command to give jarvis your zipcode and it will use it if you don't provide a zipcode in the command."
}

func (w Weather) Examples() []string {
  return []string{"jarvis weather 46723"}
}

func (w Weather) OtherDocs() []util.HelpTopic {
  return []util.HelpTopic{}
}

func (w Weather) SubCommands() []util.SubCommand {
  return []util.SubCommand{
    util.NewSubCommand("^jarvis weather$", w.NoZipCode),
    util.NewSubCommand("^jarvis weather (?P<zipcode>[0-9]{5})$", w.WithZipCode),
  }
}

func (w Weather) NoZipCode(m util.IncomingSlackMessage, r util.Regex) {
  in, storedZipCode := data.GetDatum("my zip", m.User)
  if !in {
    ws.SendMessage("If you give me your zipcode to remember I can fetch the weather for you easier. Try `jarvis help remember`.", m.Channel)
    return
  }
  lat, lng, err := service.ZipCode{}.ToLatLng(storedZipCode)
  switch err.(type) {
  case service.BadZipCodeError:
    ws.SendMessage("The zip code I have on file for you doesn't appear to be valid.", m.Channel)
  case error:
    ws.SendMessage("I've encountered an error while converting your zipcode into latitude and longitude coordinates.", m.Channel)
  }
  weather, err := service.Weather{}.ForecastFriendly(lat, lng)
  if err == nil {
    ws.SendMessage("Here's my forecast for " + storedZipCode + ":\n" + weather, m.Channel)
  } else {
    ws.SendMessage("My weather source returned an error when I tried to get your forecast.", m.Channel)
  }
}

func (w Weather) WithZipCode(m util.IncomingSlackMessage, r util.Regex) {
  zipCode := r.SubExpression(m.Text, 0)
  lat, lng, err := service.ZipCode{}.ToLatLng(zipCode)
  switch err.(type) {
  case service.BadZipCodeError:
    ws.SendMessage("The zip code you gave me doesn't appear to be valid.", m.Channel)
  case error:
    ws.SendMessage("I've encountered an error while converting your zipcode into latitude and longitude coordinates.", m.Channel)
  }
  weather, err := service.Weather{}.ForecastFriendly(lat, lng)
  if err == nil {
    ws.SendMessage(weather, m.Channel)
  } else {
    ws.SendMessage("My weather source returned an error when I tried to get your forecast.", m.Channel)
  }
}
