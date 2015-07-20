
package service

import (
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
  "os/exec"
  "strings"
)

type GitService struct{}
var GlobalGitService *GitService

func Git() *GitService {
  if GlobalGitService == nil {
    GlobalGitService = &GitService{}
  }
  return GlobalGitService
}

func (g GitService) LastCommitId() string {
  log.Trace("Getting last commit id for status")
  out, err := exec.Command("git", "rev-parse", "HEAD").Output()
  util.Check(err)
  return strings.Replace(string(out), "\n", "", -1)
}

func (g GitService) LastTag() string {
  log.Trace("Getting last tag name for status")
  out, err := exec.Command("git", "describe", "--tags").Output()
  util.Check(err)
  return strings.Replace(string(out), "\n", "", -1)
}
