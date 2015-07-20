
package config

import (

)

var OtherConf struct {
  JarvisUserId string
}

func JarvisUserId() string {
  return OtherConf.JarvisUserId
}
