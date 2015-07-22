
package service

import (
  "fmt"
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
)

type WeatherService struct {}
var GlobalWeatherService *WeatherService

func Weather() *WeatherService {
  if GlobalWeatherService == nil {
    GlobalWeatherService = &WeatherService{}
  }
  return GlobalWeatherService
}

func (w WeatherService) CurrentFriendly(zipcode string) string {
  log.Trace("Getting friendly weather report for %v", zipcode)
  lat, lng := ZipCode().ToLatLng(zipcode)
  auth := config.DarkSkyAPIToken()
  url := fmt.Sprintf("https://api.forecast.io/forecast/%v/%v,%v", auth, lat, lng)
  data := util.HttpGet(url)
  return data["currently"].(map[string]interface{})["summary"].(string)
}
