// Provides a getter and setter interface for dynamic configuration variables
// which are set while jarivs is starting.

package config

var DynamicConfig struct {
  JarvisUserId string
  JarvisActiveGroups []string
}

func JarvisUserId() string {
  return DynamicConfig.JarvisUserId
}

func JarvisActiveGroups() []string {
  return DynamicConfig.JarvisActiveGroups
}
