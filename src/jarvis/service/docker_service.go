
package service

import (
  "fmt"
  "os/exec"
  "strings"
  "time"
)

type CommandResult struct {
  Text string
  Error error
}

type Docker struct {}

func (d Docker) RunCommandInContainer(image string, command string, timeout time.Duration) (string, error) {
  resultCh := make(chan CommandResult)
  go func() {
    sp := []string{"run", image}
    for _, cmd := range strings.Split(command, " ") {
      sp = append(sp, cmd)
    }
    out, err := exec.Command("docker", sp...).Output()
    resultCh <- CommandResult{
      Text: string(out),
      Error: err,
    }
  }()
  select {
  case res := <-resultCh:
    return res.Text, res.Error
  case <-time.After(timeout):
    return "", fmt.Errorf("Your command took longer than %v to run.", Time{}.DurationToString(timeout))
  }
}
