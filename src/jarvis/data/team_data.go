
package data

import (
  "fmt"
  "jarvis/config"
  "jarvis/log"
  "jarvis/util"
)

func InitTeamData() {
  log.Info("Initializing team data into redis")
  url := fmt.Sprintf("https://slack.com/api/users.list?token=%v", config.SlackAuthToken())
  resp, _ := util.HttpGet(url)
  for _, member := range resp["members"].([]interface{}) {
    m := member.(map[string]interface{})
    id := m["id"].(string)
    name := m["name"].(string)
    Set(fmt.Sprintf("slack-username-%v", id), name)
    Set(fmt.Sprintf("slack-userid-%v", name), id)
  }
}
