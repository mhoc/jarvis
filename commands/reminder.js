
var log = require('tablog')
var scheduler = require('node-schedule')

module.exports = [

  {
    description: "allows you to schedule reminders.",

    match: [
      /jarvis remind me in ([0-9]+) (seconds?|minutes?|hours?) to (.*)/
    ],

    run: function(msg, respond) {
      log.trace('Running set reminder command')

      // Parse the input
      var addSeconds = 0;
      if (msg._matchResult[2] === "second" || msg._matchResult[2] === "seconds") {
        addSeconds = parseInt(msg._matchResult[1])
      } else if (msg._matchResult[2] === "minute" || msg._matchResult[2] === "minutes") {
        addSeconds = parseInt(msg._matchResult[1]) * 60
      } else if (msg._matchResult[2] === "hour" || msg._matchResult[2] === "hours") {
        addSeconds = parseInt(msg._matchResult[1]) * 3600
      }

      // Generate the date
      var current = new Date()
      var target = new Date(
        current.getFullYear(),
        current.getMonth(),
        current.getDate(),
        current.getHours(),
        current.getMinutes(),
        current.getSeconds() + addSeconds,
        current.getMilliseconds()
      )

      // Alert node scheduler
      scheduler.scheduleJob(target, function(m) {
        m.text = "<@" + m.user + ">: You told me to remind you to " + m._matchResult[3]
        respond(m)
      }.bind(null, msg))

      msg.text = "Ok, I'll be sure to remind you."
      respond(msg)

    }

  }

]
