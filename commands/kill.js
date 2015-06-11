
var log = require('tablog')

module.exports = {

  description: "orders jarvis to kill the node server he is running on.",

  match: [
    /jarvis kill yourself/
  ],

  run: function(msg, respond) {
    log.trace('Running admin kill yourself command')

    // Check to make sure the user is authenticated to do this
    authed = false
    require('../config').admins.forEach(function(admin) {
      if (admin === msg.user) {
        authed = true
      }
    })

    if (authed) {
      msg.text =  "If you insist, sir. Taking myself offline now."
      respond(msg)
      process.exit(0)
    } else {
      msg.text = "I appologize, but I'm afraid you don't have the proper privledges to run that command."
      respond(msg)
    }

  }

}
