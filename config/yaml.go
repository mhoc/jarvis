
package config

import (
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "github.com/mhoc/jarvis/log"
)

var ConfigFile struct {
  Admins []string `json:"admins"`
}
const ConfigLocation = "config.yaml"

func Load() {
  log.Info("Loading configuration file")
  ba, err := ioutil.ReadFile(ConfigLocation)
  if err != nil {
    log.Fatal(err.Error())
  }
  err = yaml.Unmarshal(ba, &ConfigFile)
  if err != nil {
    log.Fatal(err.Error())
  }
}
