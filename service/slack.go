
package service

import (
  "encoding/json"
  "fmt"
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
  "io/ioutil"
  "net/http"
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
    return userName
  }
  log.Trace(fmt.Sprintf("Contacting slack to convert %v to username", userId))
  url := fmt.Sprintf("https://slack.com/api/users.info?token=%v&user=%v", config.SlackAuthToken(), userId)
  res, err := http.Get(url)
  util.Check(err)
  resb, err := ioutil.ReadAll(res.Body)
  util.Check(err)
  var data map[string]interface{}
  err = json.Unmarshal(resb, &data)
  util.Check(err)
  s.cachedUserIds[userId] = data["user"].(map[string]interface{})["name"].(string)
  return s.cachedUserIds[userId]
}
