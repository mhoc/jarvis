
package service

import (
  "fmt"
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
)

type Weather struct {}

func (w Weather) ForecastFriendly(zipcode string) string {
  log.Trace("Getting friendly weather report for %v", zipcode)
  lat, lng := ZipCode{}.ToLatLng(zipcode)
  auth := config.DarkSkyAPIToken()
  url := fmt.Sprintf("https://api.forecast.io/forecast/%v/%v,%v", auth, lat, lng)
  data := util.HttpGet(url)
  // fmt.Printf("%v", data)
  return data["hourly"].(map[string]interface{})["summary"].(string)
}
