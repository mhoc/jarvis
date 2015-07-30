
package data

import (
  "github.com/mhoc/jarvis/config"
  "github.com/mhoc/jarvis/util"
  "github.com/xuyu/goredis"
)

func RedisConn() *goredis.Redis {
  conn, err := goredis.Dial(&goredis.DialConfig{Address: config.RedisURI()})
  util.Check(err)
  return conn
}
