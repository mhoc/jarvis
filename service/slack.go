
package service

import (
  "fmt"
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/data"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
)

type Slack struct {}

const (
  SLACK_CACHE_UN_FROM_UID_PREFIX = "slack-username-"
)

func (s Slack) UserNameFromUserId(userId string) string {
  in, un := data.GetCache(SLACK_CACHE_UN_FROM_UID_PREFIX + userId)
  if !in {
    log.Trace("Converting userId %v with slack api call", userId)
    url := fmt.Sprintf("https://slack.com/api/users.info?token=%v&user=%v", config.SlackAuthToken(), userId)
    slackData := util.HttpGet(url)
    un = slackData["user"].(map[string]interface{})["name"].(string)
    data.Cache(SLACK_CACHE_UN_FROM_UID_PREFIX + userId, un)
  }
  return un
}
