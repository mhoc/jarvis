
package config

import (
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "github.com/mhoc/jarvis/log"
)

var ConfigFile struct {
  Admins []string `json:"admins"`
  BoltDBPath string `json:"boltdb_path"`
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
  if ConfigFile.BoltDBPath == "" {
    log.Trace("Config.yaml 'boltdb_path' not set, defaulting to ./bolt.db")
    ConfigFile.BoltDBPath = "bolt.db"
  }
}

func Admins() []string {
  return ConfigFile.Admins
}

func BoltDBPath() string {
  return ConfigFile.BoltDBPath
}
