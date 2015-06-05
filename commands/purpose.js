
var logger = require('log4js').getLogger();
var storedPurpose = "To serve the citizens of the glorious SIGAPP nation."

module.exports = [
  {
    description: "gives me purpose in life.",

    match: [
      /jarvis your purpose is (.+)/,
      /jarvis your purpose in life is (.+)/,
      /jarvis your new purpose is (.+)/,
      /jarvis your new purpose in life is (.+)/
    ],

    run: function(msg, respond) {
      logger.info('Running set purpose command')
      storedPurpose = msg._matchResult[1]
      msg.text = 'Ok, my new purpose in life is ' + storedPurpose
      respond(msg)
    }
  },
  {
    description: "shares my purpose in life with you.",

    match: [
      /jarvis what is your purpose/
    ],

    run: function(msg, respond) {
      logger.info('Running get purpose command')
      msg.text = storedPurpose
      respond(msg)
    }

  }
]
