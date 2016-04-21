package service

import (
	"fmt"
	"jarvis/config"
	"jarvis/data"
	"jarvis/log"
	"jarvis/util"
)

type Slack struct{}

const (
	SLACK_CACHE_UN_FROM_UID_PREFIX = "slack-username-"
	SLACK_CACHE_UID_FROM_UN_PREFIX = "slack-userid-"
	SLACK_CACHE_IM_FROM_UID_PREFIX = "slack-im-channel-"
)

func (s Slack) UserNameFromUserId(userId string) string {
	in, un := data.Get(SLACK_CACHE_UN_FROM_UID_PREFIX + userId)
	if !in {
		log.Trace("Converting userId %v with slack api call", userId)
		url := fmt.Sprintf("https://slack.com/api/users.info?token=%v&user=%v", config.SlackAuthToken(), userId)
		slackData, err := util.HttpGet(url)
		util.Check(err)
		un = slackData["user"].(map[string]interface{})["name"].(string)
		data.Set(SLACK_CACHE_UN_FROM_UID_PREFIX+userId, un)
	}
	return un
}

func (s Slack) UserIdFromUserName(username string) string {
	if username[0] == '@' {
		username = username[1:]
	}
	in, uid := data.Get(SLACK_CACHE_UID_FROM_UN_PREFIX + username)
	if in {
		return uid
	} else {
		return ""
	}
}

func (s Slack) IMChannelFromUserId(userId string) (string, error) {
	if data.JarvisUserId() == userId {
		return "", fmt.Errorf("You cannot open a channel to yourself")
	}
	in, ch := data.Get(SLACK_CACHE_IM_FROM_UID_PREFIX + userId)
	if !in {
		log.Trace("Getting IM channel for user %v", userId)
		url := fmt.Sprintf("https://slack.com/api/im.open?token=%v&user=%v", config.SlackAuthToken(), userId)
		slackData, err := util.HttpGet(url)
		util.Check(err)
		ch = slackData["channel"].(map[string]interface{})["id"].(string)
		data.Set(SLACK_CACHE_IM_FROM_UID_PREFIX+userId, ch)
	}
	return ch, nil
}
