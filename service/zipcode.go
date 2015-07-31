
package service

import (
  "fmt"
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
)

type ZipCode struct {}
type BadZipCodeError struct {}

// Converts a zipcode to a pair of latitude/longitude coordinates.
// Returns a generic error if there are any problems during the http get
// to the zip code api. Also has the potential of returning a BadZipCodeError
// if the zip code provided isn't valid.
func (z ZipCode) ToLatLng(zipcode string) (float64, float64, error) {
  log.Trace("Converting %v to lat/lng", zipcode)
  auth := config.ZipCodeAPIToken()
  url := fmt.Sprintf("https://www.zipcodeapi.com/rest/%v/info.json/%v/degrees", auth, zipcode)
  data, err := util.HttpGet(url)
  if err != nil {
    return 0, 0, err
  }
  if util.MapHasElements(data, "lat", "lng") {
    return data["lat"].(float64), data["lng"].(float64), nil
  } else {
    return 0, 0, BadZipCodeError{}
  }
}

func (z BadZipCodeError) Error() string {
  return "The zip code provided doesn't seem to be valid."
}
