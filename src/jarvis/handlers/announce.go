
package handlers

import (
  "jarvis/config"
  "jarvis/ws"
)

// Announces that jarvis is online
// This isn't really a handler but that seemed like the best place to put it
func Announce() {
  announceChannels := config.AnnounceChannels()
  for _, ch := range announceChannels {
    msg := "Jarvis has just been restarted and is now ready to receive messages."
    ws.SendMessage(msg, ch)
  }
}
