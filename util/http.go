
package util

import (
  "encoding/json"
  "github.com/mhoc/jarvis/log"
  "io/ioutil"
  "net/http"
)

func HttpGet(url string) map[string]interface{} {
  log.Trace("Getting url " + url)
  res, err := http.Get(url)
  Check(err)
  resB, err := ioutil.ReadAll(res.Body)
  Check(err)
  var data map[string]interface{}
  err = json.Unmarshal(resB, &data)
  Check(err)
  // log.Trace("%v\n", data)
  return data
}
