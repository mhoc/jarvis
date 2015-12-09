
// Various functions for manipulating time objects
package util

import (
  "errors"
  "fmt"
  "github.com/jinzhu/now"
  "jarvis/log"
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

// This is a horribly complex function to convert an "absolute time" into a go time
// It uses a combination of jinzhu's NOW library and some custom parsing code to
// give the best user experience possible.
func StringToTime(ts string) (time.Time, error) {
  defaultErr := errors.New("Apologies, but I can't seem to read the time you gave me.")

  // Case 1: "Remind me at 8 to do X"
  // NOW will parse this to always mean "8AM of the current day", but I want this to actually mean
  //  - "8PM Today if it is after 8am today but before 8pm today"
  //  - "8AM Tomorrow if it is after 8AM today and after 8pm today"
  works, t, err := parseAbsTimeLoneNumber(ts, defaultErr)
  if err != nil {
    return t, err
  }
  if works {
    return t, nil
  }

  // At this point we pass the timestamp over to NOW to parse
  t, err = now.Parse(ts)
  fmt.Printf("%v\n", t.String())
  if err != nil {
    return t, defaultErr
  }
  return t, nil

}

func parseAbsTimeLoneNumber(ts string, defaultErr error) (bool, time.Time, error) {
  if NewRegex("^[0-9]{1,2}$").Matches(ts) || NewRegex("^[0-9]{1,2}:[0-9]{2}$").Matches(ts) {
    log.Trace("Parsing absolute time assuming lone number")
    t, err := now.Parse(ts)
    if err != nil {
      log.Trace("Error: %v\n", err.Error())
      return false, t, defaultErr
    }
    if t.Before(time.Now()) {
      t = t.Add(12 * time.Hour)
    }
    if t.Before(time.Now()) {
      t = t.Add(12 * time.Hour)
    }
    return true, t, nil
  } else {
    return false, time.Now(), nil
  }
}
