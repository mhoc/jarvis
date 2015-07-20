
package util

import (
  "github.com/jbrukh/bayesian"
)

type Command interface {
  Class() bayesian.Class
  TrainingStrings() []string
  Description() string
  Execute(IncomingSlackMessage)
}
