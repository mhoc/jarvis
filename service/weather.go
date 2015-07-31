
package service

import (
  "fmt"
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
)

type Weather struct {}

func (w Weather) ForecastFriendly(zipcode string) (string, error) {
  log.Trace("Getting friendly weather report for %v", zipcode)
  lat, lng, err := ZipCode{}.ToLatLng(zipcode)
  if err != nil {
    return "", err
  }
  auth := config.DarkSkyAPIToken()
  url := fmt.Sprintf("https://api.forecast.io/forecast/%v/%v,%v", auth, lat, lng)
  data, err := util.HttpGet(url)
  if err == nil {
    return data["hourly"].(map[string]interface{})["summary"].(string), nil
  } else {
    return "", err
  }
}
