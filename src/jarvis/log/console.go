package log

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Prefix() string {
	t := time.Now()
	pre := fmt.Sprintf("[%4v:%02v:%02v %02v:%02v:%02v] ", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
	return pre
}

func Trace(s string, args ...interface{}) {
	s = strings.Replace(s, "\n", " ", -1)
	fmt.Printf(fmt.Sprintf("%v\n", FormatColor(Prefix()+s, BOLD_GRAY)), args...)
}

func Info(s string, args ...interface{}) {
	s = strings.Replace(s, "\n", " ", -1)
	fmt.Printf(fmt.Sprintf("%v%v\n", FormatColor(Prefix(), BOLD_GREEN), s), args...)
}

func Warn(s string, args ...interface{}) {
	s = strings.Replace(s, "\n", " ", -1)
	fmt.Printf(fmt.Sprintf("%v%v\n", FormatColor(Prefix(), YELLOW), s), args...)
}

func Error(e error) bool {
	if e != nil {
		s := e.Error()
		s = strings.Replace(s, "\n", " ", -1)
		fmt.Printf("%v\n", s)
		return true
	}
	return false
}

func Fatal(s string, args ...interface{}) {
	s = strings.Replace(s, "\n", " ", -1)
	fmt.Printf(fmt.Sprintf("%v\n", FormatColor(Prefix()+s, BOLD_RED)), args...)
	os.Exit(1)
}
