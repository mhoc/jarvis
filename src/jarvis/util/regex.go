// Command matching functionality and
// Wrapper functions around the regexp package to decrease boilerplate code

package util

import (
  "regexp"
)

type Regex struct {
  goRegex *regexp.Regexp
}

func NewRegex(pattern string) Regex {
  goReg, err := regexp.Compile(pattern)
  Check(err)
  return Regex{
    goRegex: goReg,
  }
}

func (r Regex) String() string {
  return r.goRegex.String()
}

func (r Regex) Matches(test string) bool {
  return r.goRegex.MatchString(test)
}

func (r Regex) NSubExpressions(test string) int {
  res := r.goRegex.FindAllStringSubmatch(test, -1)
  return len(res[0])
}

func (r Regex) SubExpression(test string, i int) string {
  res := r.goRegex.FindAllStringSubmatch(test, -1)
  return res[0][i+1]
}
