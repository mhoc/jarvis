// Provides an API and error checking for environment configuration variables,
// which are mostly API keys

package config

import (
  "os"
  "github.com/mhoc/jarvis/log"
)

func SlackAuthToken() string {
  auth := os.Getenv("SLACK_AUTH_TOKEN")
  if auth == "" {
    log.Fatal("Must provide a slack auth token under the envvar SLACK_AUTH_TOKEN")
  }
  return auth
}
