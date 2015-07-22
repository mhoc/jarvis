
package service

import (
  "fmt"
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
)

type ZipCodeService struct {}
var GlobalZipCodeService *ZipCodeService

func ZipCode() *ZipCodeService {
  if GlobalZipCodeService == nil {
    GlobalZipCodeService = &ZipCodeService{}
  }
  return GlobalZipCodeService
}

func (z ZipCodeService) ToLatLng(zipcode string) (float64, float64) {
  log.Trace("Converting %v to lat/lng", zipcode)
  auth := config.ZipCodeAPIToken()
  url := fmt.Sprintf("https://www.zipcodeapi.com/rest/%v/info.json/%v/degrees", auth, zipcode)
  data := util.HttpGet(url)
  return data["lat"].(float64), data["lng"].(float64)
}
