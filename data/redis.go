
package data

import (
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
  "github.com/xuyu/goredis"
  "time"
)

func CheckRedisConn() {
  _, err := goredis.Dial(&goredis.DialConfig{Address: config.RedisURI()})
  if err != nil {
    log.Warn("Redis must be running on jarvis' machine")
    util.Check(err)
  }
}

func redisConn() *goredis.Redis {
  conn, err := goredis.Dial(&goredis.DialConfig{Address: config.RedisURI()})
  util.Check(err)
  return conn
}

func set(key string, value string) {
  conn := redisConn()
  err := conn.Set(key, value, 0, 0, false, false)
  util.Check(err)
}

func setTimeout(key string, value string, timeout time.Duration) {
  conn := redisConn()
  err := conn.Set(key, value, int(timeout.Seconds()), 0, false, false)
  util.Check(err)
}

func get(key string) (bool, string) {
  conn := redisConn()
  resp, err := conn.Get(key)
  util.Check(err)
  if resp == nil {
    return false, ""
  } else {
    return true, string(resp)
  }
}
