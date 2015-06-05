
require('./socket/get_slack_ws')(function(err, wsurl) {
  require('./socket/init')().connect(wsurl)
})
