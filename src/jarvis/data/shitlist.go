
// The Jarvis Shitlist
// If jarvis detects someone is trying to shit all over him, jarvis will
// temporarily shitlist them and not respond to any messages they send.
package data

import (
  "fmt"
  "time"
)

var (
  ShitlistTime = 10 * time.Minute
)

// This is implemented with keys instead of a set because I want redis
// to handle cache expiration
func Shitlist(user string) {
  SetTimeout(fmt.Sprintf("shitlist-%v", user), user, ShitlistTime)
}

func GetShitlist() []string {
  u := []string{}
  for _, key := range Keys("shitlist-*") {
    _, user := Get(key)
    u = append(u, user)
  }
  return u
}
