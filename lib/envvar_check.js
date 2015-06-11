
var log = require('tablog')

module.exports = function() {
  log.trace('Scanning expected envvars.')

  if (process.env.MACHINE == null) {
    log.warn("$MACHINE is not set.")
  }

  if (process.env.DARK_SKY_API_TOKEN == null) {
    log.warn("$DARK_SKY_API_TOKEN is not set.")
  }

  if (process.env.ZIP_CODE_API_TOKEN == null) {
    log.warn("$ZIP_CODE_API_TOKEN is not set.")
  }

}
