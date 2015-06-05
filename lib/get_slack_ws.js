
var logger = require('log4js').getLogger()
var request = require('request')

// Returns the slack websocket url through the callback
// @param callback function(err, wsURL)
module.exports = function(callback) {
  logger.trace('Getting websocket url from slack api')

  // Read the token from the envvars
  var auth = process.env.SLACK_AUTH_TOKEN
  if (auth == null || auth === "") {
    logger.error("You must provide a slack auth token under the envvar SLACK_AUTH_TOKEN.")
    process.exit(1)
  }

  // Make the request asynchronously spelling oh my god
  request({url: "https://slack.com/api/rtm.start?token=" + auth, json: true}, function(err, res, body) {
    if (err || res.statusCode >= 300) {
      log.error("Error requesting websocket url from slack")
    }

    // We store the userid of jarvis himself so we can ignore messages he sends
    // Otherwise it might be possible to enter an infinite loop (though in practice it doesnt seem to)
    require('../config').jarvisUserId = body.self.id
    callback(err, body.url)

  })

}
