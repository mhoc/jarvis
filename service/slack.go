
package service

import (
  "fmt"
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
)

type Slack struct {}

func (s Slack) UserNameFromUserId(userId string) string {
  log.Trace("Converting userId %v with slack api call", userId)
  url := fmt.Sprintf("https://slack.com/api/users.info?token=%v&user=%v", config.SlackAuthToken(), userId)
  data := util.HttpGet(url)
  return data["user"].(map[string]interface{})["name"].(string)
}
