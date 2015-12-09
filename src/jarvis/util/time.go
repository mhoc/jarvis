
// Various functions for manipulating time objects
package util

import (
  "errors"
  "fmt"
  "github.com/jinzhu/now"
  "math"
  "strings"
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

func StringToDuration(durStr string) (time.Duration, error) {
  // Run the duration string through a bunch of processing to get it into a time.Duration format that can be parsed by go
  durStr = strings.Replace(durStr, " seconds", "s", -1)
  durStr = strings.Replace(durStr, " second", "s", -1)
  durStr = strings.Replace(durStr, " minutes", "m", -1)
  durStr = strings.Replace(durStr, " minute", "m", -1)
  durStr = strings.Replace(durStr, " hours", "h", -1)
  durStr = strings.Replace(durStr, " hour", "h", -1)
  durStr = strings.Replace(durStr, " ", "", -1)
  d, err := time.ParseDuration(durStr)
  if err != nil {
    return d, errors.New("Apologies, but I can't seem to parse your duration string.")
  }
  if d.Hours() < 0 || d.Minutes() < 0 || d.Seconds() < 0 {
    return d, errors.New("Apologies, but my functionality does not include the recognition of negative time.")
  }
  return d, nil
}

func StringToTime(ts string) (time.Time, error) {
  stockErr := errors.New("Apologies, but I can't seem to read the time you gave me.")
  t, err := now.Parse(ts)
  fmt.Printf("%v\n", t.String())
  if err != nil {
    return t, stockErr
  }
  return t, nil
}
