
// Jarvis' uptime really isn't a configuration setting, but this is the single
// package that doesn't import any other jarvis/ packages, and it makes the
// most sense to store that information here.
package config

import (
  "time"
)

var (
  StartedAt time.Time
)

func Uptime() time.Duration {
  return time.Since(StartedAt)
}
