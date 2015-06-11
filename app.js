
var log = require('tablog')
log.trace('Starting')

require('./lib/get_slack_ws')(function(err, wsurl) {
  require('./lib/init_socket')().connect(wsurl)
})
