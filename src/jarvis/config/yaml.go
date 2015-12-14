
package config

import (
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "jarvis/log"
  "time"
)

var ConfigFile struct {
  Redis string `yaml:"redis"`
  Tokens struct {
    DarkSky string `yaml:"darksky"`
    Slack string `yaml:"slack"`
    ZipCode string `yaml:"zipcode"`
    Aws struct {
      Access string `yaml:"access"`
      Secret string `yaml:"secret"`
    } `yaml:"aws"`
  }
  Admins []string `yaml:"admins"`
  Location string `yaml:"location"`
  ChannelWhitelist []string `yaml:"channel_whitelist"`
  ChannelBlacklist []string `yaml:"channel_blacklist"`
  Announcements []string `yaml:"announcements"`
  UserBlacklist []string `yaml:"user_blacklist"`
  Static []map[string]interface{} `yaml:"static"`
}

const ConfigLocation = "config.yaml"

func LoadYaml() {
  log.Info("Setting uptime start time")
  StartedAt = time.Now()
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
  if ConfigFile.Redis == "" {
    log.Fatal("Must provide a redis uri inside config.yaml under 'redis'")
  }
  if ConfigFile.Tokens.Slack == "" {
    log.Fatal("Must provide a slack auth token inside config.yaml under tokens.slack")
  }
  if ConfigFile.Tokens.DarkSky == "" {
    log.Fatal("Must provide a darksky weather auth token inside config.yaml under tokens.darksky")
  }
  if ConfigFile.Tokens.ZipCode == "" {
    log.Fatal("Must provide a zip code api auth token inside config.yaml under tokens.zipcode")
  }
  if ConfigFile.Admins == nil {
    ConfigFile.Admins = make([]string, 0, 0)
  }
  if len(ConfigFile.Admins) == 0 {
    log.Warn("Admin level commands will be unavailable if no admins are provided in config.yaml")
  }
  if ConfigFile.Location == "" {
    log.Warn("No human readable location name provided under `location` in `config.yaml`")
    ConfigFile.Location = "a very secret place"
  }
  if ConfigFile.Static == nil {
    log.Warn("You can provide static data under the `static` key in config.yaml if you like")
  }
}

func RedisURI() string {
  return ConfigFile.Redis
}

func SlackAuthToken() string {
  return ConfigFile.Tokens.Slack
}

func DarkSkyAPIToken() string {
  return ConfigFile.Tokens.DarkSky
}

func ZipCodeAPIToken() string {
  return ConfigFile.Tokens.ZipCode
}

func AwsCredentials() (string, string) {
  return ConfigFile.Tokens.Aws.Access, ConfigFile.Tokens.Aws.Secret
}

func Admins() []string {
  return ConfigFile.Admins
}

func AnnounceChannels() []string {
  return ConfigFile.Announcements
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

func HasWhitelist() bool {
  return len(ConfigFile.ChannelWhitelist) != 0
}

func ChannelIsWhitelisted(ch string) bool {
  for _, i := range ConfigFile.ChannelWhitelist {
    if ch == i {
      return true
    }
  }
  return false
}

func ChannelIsBlacklisted(ch string) bool {
  for _, i := range ConfigFile.ChannelBlacklist {
    if ch == i {
      return true
    }
  }
  return false
}

func UserIsBlacklisted(uid string) bool {
  for _, i := range ConfigFile.UserBlacklist {
    if i == uid {
      return true
    }
  }
  return false
}

func Static() []map[string]interface{} {
  return ConfigFile.Static
}
