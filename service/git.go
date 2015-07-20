
package service

import (
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
  "os/exec"
  "strings"
)

type Git struct{}

func (g Git) LastCommitId() string {
  log.Trace("Getting last commit id for status")
  out, err := exec.Command("git", "rev-parse", "HEAD").Output()
  util.Check(err)
  return strings.Replace(string(out), "\n", "", -1)
}
