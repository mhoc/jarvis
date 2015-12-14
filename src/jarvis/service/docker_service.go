
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
  cmd := exec.Command("docker", "run", "python", "python", "-c", command)
  go func() {
    out, err := cmd.Output()
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
    cmd.Process.Kill()
    return "", fmt.Errorf("Your command took longer than %v to run and has thus been killed.", Time{}.DurationToString(timeout))
  }
}
