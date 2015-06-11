
var log = require('tablog')

module.exports = [

  {
    description: "orders jarvis to kill the node server he is running on.",

    match: [
      /jarvis kill yourself/
    ],

    run: function(msg, respond) {
      log.trace('Running admin kill yourself command')
    }

  }

]
