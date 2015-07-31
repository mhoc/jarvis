
package data

import (
  "github.com/mhoc/jarvis/log"
  "time"
)

const (
  CACHE_PREFIX = "cache-"
)

func Cache(key string, val string) {
  log.Trace("Caching %v -> %v", key, val)
  set(CACHE_PREFIX + key, val)
}

func CacheTimeout(key string, val string, timeout time.Duration) {
  log.Trace("Caching %v -> %v with timeout %vs", key, val, timeout.Seconds())
  setTimeout(CACHE_PREFIX + key, val, timeout)
}

func GetCache(key string) (bool, string) {
  log.Trace("Getting cache entry %v", key)
  return get(CACHE_PREFIX + key)
}
