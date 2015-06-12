// Checks configuration parameters from config.json and env vars
// If you're missing a parameters, you will get yelled at when you npm start
// so you can fix it right away.

var log = require('tablog')

var expectedEnvars = [
  "SLACK_AUTH_TOKEN",
  "DARK_SKY_API_TOKEN",
  "ZIP_CODE_API_TOKEN"
]

var expectedConfigKeys = {
  "admins": "[object Array]",
  "machine_name": "[object String]"
}

module.exports = function() {
  log.trace('Scanning expected envvars.')

  expectedEnvars.forEach(function(envv) {
    if (process.env[envv] == null) {
      log.warn("$" + envv + " is not set.")
    }
  })

  for (var key in expectedConfigKeys) {
    if (expectedConfigKeys.hasOwnProperty(key)) {
      if (require('../config')[key] == null) {
        log.warn("config.json -> " + key + " is not set")
      }
      if (Object.prototype.toString.call(require('../config')[key]) != expectedConfigKeys[key]) {
        log.warn("config.json -> " + key + " should be of type " + expectedConfigKeys[key])
      }
    }
  }

  // Put special config checks here you want to run
  if (require('../config').admins.length == 0) {
    log.warn('config.json -> Providing no admins will cause admin level functions to be unusable')
  }

}
