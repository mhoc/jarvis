
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

type Slack struct {}
var slackCachedUserIds = make(map[string]string)

func (s Slack) UserNameFromUserId(userId string) string {
  if userName, in := slackCachedUserIds[userId]; in {
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
  slackCachedUserIds[userId] = data["user"].(map[string]interface{})["name"].(string)
  return slackCachedUserIds[userId]
}
