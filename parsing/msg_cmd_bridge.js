
var commands = require('../util/cmd_dir')
var logger = require('log4js').getLogger()

// Glue logic between a raw slack message and a format that the commands
// expect. This function reads every message and attempts to find a command
// which it can match against. If it finds a match, it invokes that function.
module.exports = function(slackMsg, writeback) {
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
