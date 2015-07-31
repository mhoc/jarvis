
package config

import (
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "github.com/mhoc/jarvis/log"
)

var ConfigFile struct {
  Admins []string `json:"admins"`
  Location string `json:"location"`
}
const ConfigLocation = "config.yaml"

func LoadYaml() {
  log.Info("Loading configuration file")
  ba, err := ioutil.ReadFile(ConfigLocation)
  if err != nil {
    log.Fatal(err.Error())
  }
  err = yaml.Unmarshal(ba, &ConfigFile)
  if err != nil {
    log.Fatal(err.Error())
  }
  ValidateYaml()
}

func ValidateYaml() {
  log.Info("Validating configuration file")
  if ConfigFile.Admins == nil {
    log.Fatal("Must provide a list of admin jarvis users in config.yaml")
  }
  if len(ConfigFile.Admins) == 0 {
    log.Warn("Admin level commands will be unavailable if no admins are provided in config.yaml")
  }
  if ConfigFile.Location == "" {
    log.Fatal("Must provide a human readable location name for where jarvis is running in config.yaml")
  }
}

func Admins() []string {
  return ConfigFile.Admins
}

func IsAdmin(userid string) bool {
  for _, admin := range Admins() {
    if admin == userid {
      return true
    }
  }
  return false
}

func Location() string {
  return ConfigFile.Location
}
