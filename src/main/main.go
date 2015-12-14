package main

import (
  "jarvis/commands"
  "jarvis/config"
  "jarvis/data"
  "jarvis/handlers"
  "jarvis/log"
  "jarvis/ws"
  "runtime"
)

func main() {
  log.Info("Starting Jarvis")
  runtime.GOMAXPROCS(runtime.NumCPU())
  config.LoadYaml()
  data.CheckRedisConn()
  data.InitTeamData()
  handlers.Init()
  commands.StartReminderLoop()
  ws.Init()
  handlers.Announce()
  log.Info("Jarvis is live and receiving messages")
  select {}
}
