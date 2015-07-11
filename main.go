
package main

import (
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/ws"
  log "github.com/Sirupsen/logrus"
  "runtime"
)

func main() {
  log.Info("Starting Jarvis")
  runtime.GOMAXPROCS(runtime.NumCPU())
  config.Load()
  ws.Init()
  for {}
}
