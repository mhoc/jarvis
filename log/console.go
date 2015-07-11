
package log

import (
  "fmt"
  "github.com/mhoc/golor"
  "os"
  "time"
)

func Prefix() string {
  t := time.Now()
  pre := fmt.Sprintf("%v:%v:%v %v:%v:%v| ", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
  return pre
}

func Info(s string) {
  fmt.Printf("%v%v\n", golor.Green(Prefix()), s)
}

func Fatal(s string) {
  fmt.Printf("%v%v\n", golor.Red(Prefix()), s)
  os.Exit(1)
}
