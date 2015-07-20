
package util

import (
  "github.com/mhoc/jarvis/log"
)

func Check(e error) {
  if e != nil {
    log.Fatal(e.Error())
  }
}
