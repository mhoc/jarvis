
package service

import (
  "fmt"
  "jarvis/log"
  "os/exec"
  // "strings"
  "time"
)

type CommandResult struct {
  Text string
  Error error
}

type Docker struct {}

func (d Docker) RunPythonInContainer(command string, timeout time.Duration) (string, error) {
  log.Info("Executing command '%v' in container 'python'", command)
  resultCh := make(chan CommandResult)
  go func() {
    out, err := exec.Command("docker", "run", "python", "python", "-c", command).Output()
    log.Trace("Result: %v", string(out))
    if err != nil {
      log.Trace("Error: %v", err.Error())
    } else {
      log.Trace("No error")
    }
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
