
package util

import (
  "regexp"
)

type Command interface {
  Matches() []*regexp.Regexp
  Help(IncomingSlackMessage)
  Execute(IncomingSlackMessage)
}
