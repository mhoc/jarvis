
var log = require('tablog')
var request = require('request')

// Exports an object which contains a bunch of functions related to slack
module.exports = {

  // @arg userId the slack userid of the user
  // @callback function(err, bool)
  is_user_online: function(userId, callback) {
    log.trace('Determining if user ' + userId + ' is online right now')

    // Get slack token
    var token = process.env.SLACK_AUTH_TOKEN
    if (token == null || token === "") {
      log.fatal('No slack auth token set in SLACK_AUTH_TOKEN')
    }

    request({
      url: "https://slack.com/api/users.getPresence?token=" + token + "&user=" + userId,
      json: true
    }, function(err, res, body) {
      if (err != null) {
        log.error('Error contacting slack api for user status')
      }
      callback(err, body.presence === "active")
    })

  }

}
