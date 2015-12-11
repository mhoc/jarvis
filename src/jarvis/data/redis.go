
package data

import (
  "gopkg.in/redis.v3"
  "jarvis/config"
  "jarvis/log"
  "jarvis/util"
  "strings"
  "time"
)

func CheckRedisConn() {
  client := redis.NewClient(&redis.Options{
    Addr: config.RedisURI(),
  })
  _, err := client.Ping().Result()
  if err != nil {
    log.Warn("Redis must be running at the URl specificed in config.yaml")
    util.Check(err)
  }
}

func redisConn() *redis.Client {
  return redis.NewClient(&redis.Options{
    Addr: config.RedisURI(),
  })
}

func Set(key string, value string) {
  conn := redisConn()
  err := conn.Set(key, value, 0).Err()
  util.Check(err)
}

func SetTimeout(key string, value string, timeout time.Duration) {
  conn := redisConn()
  err := conn.Set(key, value, timeout).Err()
  util.Check(err)
}

func Get(key string) (bool, string) {
  conn := redisConn()
  resp, err := conn.Get(key).Result()
  if err != nil && strings.Contains(err.Error(), "WRONGTYPE") {
    return false, ""
  }
  if err != nil && strings.Contains(err.Error(), "nil") {
    return false, ""
  }
  util.Check(err)
  return true, resp
}

func Remove(key string) {
  conn := redisConn()
  conn.Del(key).Result()
}

func Keys(match string) []string {
  conn := redisConn()
  resp, err := conn.Keys(match).Result()
  util.Check(err)
  return resp
}

func SetAdd(setname string, value string) {
  conn := redisConn()
  err := conn.SAdd(setname, value).Err()
  util.Check(err)
}

func SetGet(setname string) []string {
  conn := redisConn()
  members, err := conn.SMembers(setname).Result()
  if err != nil {
    log.Info(err.Error())
    return []string{}
  } else {
    return members
  }
}
