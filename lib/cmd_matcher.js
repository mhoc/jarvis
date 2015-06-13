
var commands = require('./load_cmd_dir')
var config = require('../config')
var log = require('tablog')

// Glue logic between a raw slack message and a format that the commands
// expect. This function reads every message and attempts to find a command
// which it can match against. If it finds a match, it invokes that function.
module.exports = function(slackMsg, writeback) {

  // Check to make sure the message is not from jarvis
  if (slackMsg.user === require('../config')._jarvisUserId) {
    return
  }

  // If we have channels whitelisted, decline the message if it isnt from that channel
  if (config.whitelist_channels) {
    var inn = false
    config.whitelist_channels.forEach(function(channel) {
      if (channel === slackMsg.channel) {
        inn = true
      }
    })
    if (!inn) {
      log.info("Received message on non-whitelisted channel, not responding.")
      return
    }
  }

  // If we have channels blacklisted, ignore the message from that channel
  if (config.blacklist_channels) {
    config.blacklist_channels.forEach(function(channel) {
      if (channel === slackMsg.channel) {
        log.info("Received message on blacklisted channel, not responding.")
        return
      }
    })
  }

  // If the first word of the message is jarvis, lowercase it
  // This is just a convenience thing so we dont have to write
  // matching regex for both upper and lower case just to support
  // the autocapitalization feature on most smartphones
  if (slackMsg.text.split(' ')[0] === "Jarvis") {
    slackMsg.text = 'j' + slackMsg.text.slice(1)
  }

  commands.forEach(function(command) {
    command.match.forEach(function(matchPhrase) {
      result = slackMsg.text.match(matchPhrase)
      if (result != null) {
        slackMsg._matchResult = result
        command.run(slackMsg, writeback)
        return
      }
    })
  })
}
