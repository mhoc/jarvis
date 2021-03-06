package service

import (
	"jarvis/log"
	"jarvis/util"
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

func (g Git) LastTag() string {
	log.Trace("Getting last tag name for status")
	out, err := exec.Command("git", "describe", "--tags").Output()
	util.Check(err)
	return strings.Replace(string(out), "\n", "", -1)
}
