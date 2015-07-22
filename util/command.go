
package util

import (
  "regexp"
)

type Command interface {
  Matches() []*regexp.Regexp
  Description() string
  Execute(IncomingSlackMessage)
}
