
package util

import (
  "regexp"
)

type Command interface {
  Name() string
  Matches() []*regexp.Regexp

  // For documentation purposes
  Description() string
  Format() string
  Examples() []string
  OtherDocs() []HelpTopic

  // Behavior
  Execute(IncomingSlackMessage)

}
