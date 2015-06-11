
var log = require('tablog')
var request = require('request')

// Returns the slack websocket url through the callback
// @param callback function(err, wsURL)
module.exports = function(callback) {
  log.info('Getting websocket url from slack api')

  // Read the token from the envvars
  var auth = process.env.SLACK_AUTH_TOKEN
  if (auth == null || auth === "") {
    log.fatal("You must provide a slack auth token under the envvar SLACK_AUTH_TOKEN.")
  }

  // Make the request asynchronously spelling oh my god
  request({url: "https://slack.com/api/rtm.start?token=" + auth, json: true}, function(err, res, body) {
    if (err || res.statusCode >= 300) {
      log.fatal("Error requesting websocket url from slack")
    }

    // We store the userid of jarvis himself so we can ignore messages he sends
    // Otherwise it might be possible to enter an infinite loop (though in practice it doesnt seem to)
    require('../config')._jarvisUserId = body.self.id

    // And we store the channels jarvis is in incase he needs to make a global announcement
    channels = []
    body.channels.forEach(function(channel) {
      channels.push(channel.id)
    })
    body.groups.forEach(function(group) {
      channels.push(group.id)
    })
    require('../config')._jarvisChannels = channels

    callback(err, body.url)

  })

}
