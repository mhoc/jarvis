
package service

import (
  "fmt"
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
)

type Weather struct {}

func (w Weather) ForecastFriendly(lat float64, lng float64) (string, error) {
  log.Trace("Getting friendly weather report for %v %v", lat, lng)
  url := fmt.Sprintf("https://api.forecast.io/forecast/%v/%v,%v", config.DarkSkyAPIToken(), lat, lng)
  data, err := util.HttpGet(url)
  if err == nil {
    return data["hourly"].(map[string]interface{})["summary"].(string), nil
  } else {
    return "", err
  }
}
