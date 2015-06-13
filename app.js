
var log = require('tablog')
log.trace('Starting setup process')

// Check envvars to see if they're all set up
require('./lib/config_check')()

require('./lib/get_slack_ws')(function(err, wsurl) {
  require('./lib/init_socket')().connect(wsurl)
})
