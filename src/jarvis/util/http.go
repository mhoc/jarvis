
package util

import (
  "encoding/json"
  "jarvis/log"
  "io/ioutil"
  "net/http"
)

func HttpGet(url string) (map[string]interface{}, error) {
  log.Trace("Getting url " + url)
  res, err := http.Get(url)
  if err != nil {
    log.Warn("Error on http request, returning to client")
    return nil, err
  }
  resB, err := ioutil.ReadAll(res.Body)
  Check(err)
  var data map[string]interface{}
  err = json.Unmarshal(resB, &data)
  Check(err)
  return data, nil
}

func HttpGetArr(url string) ([]map[string]interface{}, error) {
  log.Trace("Getting url " + url)
  res, err := http.Get(url)
  if err != nil {
    log.Warn("Error on http request, returning to client")
    return nil, err
  }
  resB, err := ioutil.ReadAll(res.Body)
  Check(err)
  var data []map[string]interface{}
  err = json.Unmarshal(resB, &data)
  Check(err)
  return data, nil
}
