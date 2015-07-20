// Provides an API and error checking for environment configuration variables,
// which are mostly API keys

package config

import (
  "os"
  "github.com/mhoc/jarvis/log"
)

var ExpectedEnvs = []string{
  "SLACK_AUTH_TOKEN",
  "DARK_SKY_API_TOKEN",
  "ZIP_CODE_API_TOKEN",
}

func VerifyEnvs() {
  for _, env := range ExpectedEnvs {
    s := os.Getenv(env)
    if s == "" {
        log.Fatal("Expected environment variable under $" + env)
    }
  }
}

func SlackAuthToken() string {
  return os.Getenv("SLACK_AUTH_TOKEN")
}
func DarkSkyAPIToken() string {
  return os.Getenv("DARK_SKY_API_TOKEN")
}
func ZipCodeAPIToken() string {
  return os.Getenv("ZIP_CODE_API_TOKEN")
}
