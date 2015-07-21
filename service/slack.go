
package service

import (
  "fmt"
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
)

type SlackService struct {
  cachedUserIds map[string]string
}
var GlobalSlackService *SlackService

func Slack() *SlackService {
  if GlobalSlackService == nil {
    GlobalSlackService = &SlackService{cachedUserIds: make(map[string]string)}
  }
  return GlobalSlackService
}

func (s SlackService) UserNameFromUserId(userId string) string {
  if userName, in := s.cachedUserIds[userId]; in {
    log.Trace("Converting cached userId %v", userId)
    return userName
  }
  log.Trace("Converting userId %v with slack api call", userId)
  url := fmt.Sprintf("https://slack.com/api/users.info?token=%v&user=%v", config.SlackAuthToken(), userId)
  data := util.HttpGet(url)
  s.cachedUserIds[userId] = data["user"].(map[string]interface{})["name"].(string)
  return s.cachedUserIds[userId]
}
