
// Various functions for manipulating time objects
package util

import (
  "fmt"
  "math"
  "time"
)

// Replaces the normal duration.String() function with one which formats the
// data in a much more human readable way.
func DurationToString(d time.Duration) string {
  s := ""
  hours := math.Floor(d.Hours())
  if hours > 0 {
    s += fmt.Sprintf("%v hour", hours)
  }
  if hours > 1 {
    s += "s"
  }
  minutes := int(d.Minutes()) % 60
  if minutes > 0 {
    if hours > 0 {
      s += " "
    }
    s += fmt.Sprintf("%v minute", minutes)
    if minutes > 1 {
      s += "s"
    }
  }
  seconds := int(d.Seconds()) % 60
  if seconds > 0 {
    if hours > 0 || minutes > 0 {
      s += " "
    }
    s += fmt.Sprintf("%v second", seconds)
    if seconds > 1 {
      s += "s"
    }
  }
  return s
}
