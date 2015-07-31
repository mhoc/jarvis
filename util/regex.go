// Wrapper functions around the regexp package to decrease boilerplate code

package util

import (
  "regexp"
)

type JarvisRegex struct {
  goRegex *regexp.Regexp
}

func NewRegex(pattern string) JarvisRegex {
  goReg, err := regexp.Compile(pattern)
  Check(err)
  return JarvisRegex{
    goRegex: goReg,
  }
}

func (jr JarvisRegex) String() string {
  return jr.goRegex.String()
}

func (jr JarvisRegex) Matches(test string) bool {
  return jr.goRegex.MatchString(test)
}

func (jr JarvisRegex) SubExpression(test string, i int) string {
  r := jr.goRegex.FindAllStringSubmatch(test, -1)
  return r[0][i+1]
}
