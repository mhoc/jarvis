
package log

import (
  "fmt"
  "os"
  "time"
)

func Prefix() string {
  t := time.Now()
  pre := fmt.Sprintf("[%4v:%02v:%02v %02v:%02v:%02v] ", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
  return pre
}

func Trace(s string) {
  fmt.Printf("%v\n", FormatColor(Prefix() + s, BOLD_GRAY))
}

func Info(s string) {
  fmt.Printf("%v%v\n", FormatColor(Prefix(), BOLD_GREEN), s)
}

func Fatal(s string) {
  fmt.Printf("%v%v\n", FormatColor(Prefix(), BOLD_RED), s)
  os.Exit(1)
}
