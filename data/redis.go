
package data

import (
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
  "github.com/xuyu/goredis"
)

func CheckRedisConn() {
  _, err := goredis.Dial(&goredis.DialConfig{Address: config.RedisURI()})
  if err != nil {
    log.Warn("Redis must be running on jarvis' machine")
    util.Check(err)
  }
}

func RedisConn() *goredis.Redis {
  conn, err := goredis.Dial(&goredis.DialConfig{Address: config.RedisURI()})
  util.Check(err)
  return conn
}
