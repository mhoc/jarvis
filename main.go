
package main

import (
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/handlers"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/ws"
  "runtime"
)

func main() {
  log.Info("Starting Jarvis")
  runtime.GOMAXPROCS(runtime.NumCPU())
  config.Load()
  ws.Init()
  handlers.Init()
  log.Info("Jarvis is live and receiving messages")
  for {}
}
