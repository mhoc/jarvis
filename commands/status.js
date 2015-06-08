
var exec = require('child_process').exec
var logger = require('log4js').getLogger()

module.exports = {
  description: "provides a cool status printout.",

  match: [
    /jarvis status/,
    /jarvis ping/
  ],

  run: function(msg, respond) {
    logger.info("Running status command")
    msg.text = "Don't worry, I'm alive.\n"

    // Exec git status
    exec('git rev-parse HEAD', function(err, stdout, stderr) {
      if (err) {
        logger.warn('Error getting latest git commit hash')
        respond(msg)
        return
      }
      msg.text += "I'm currently running jarvis version " + stdout.substring(0, 6) + "."
      respond(msg)
    })

  }
}
