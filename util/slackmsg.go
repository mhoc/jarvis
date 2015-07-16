
package util

type IncomingSlackMessage struct {
  Type string `json:"type"`
  Channel string `json:"channel"`
  User string `json:"user"`
  Text string `json:"text"`
  Timestamp string `json:"ts"`
}

type OutgoingSlackMessage struct {
  Channel string `json:"channel"`
  Text string `json:"text"`
}
