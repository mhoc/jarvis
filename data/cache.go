
package data

import (
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
  "time"
)

func Cache(key string, val string) {
  log.Trace("Caching %v -> %v", key, val)
  conn := RedisConn()
  err := conn.Set(key, val, 0, 0, false, false)
  util.Check(err)
}

func CacheTimeout(key string, val string, timeout time.Duration) {
  log.Trace("Caching %v -> %v with timeout %vs", key, val, timeout.Seconds())
  conn := RedisConn()
  err := conn.Set(key, val, int(timeout.Seconds()), 0, false, false)
  util.Check(err)
}

func GetCache(key string) (bool, string) {
  log.Trace("Getting cache entry %v", key)
  conn := RedisConn()
  resp, err := conn.Get(key)
  util.Check(err)
  if resp == nil {
    return false, ""
  } else {
    return true, string(resp)
  }
}
