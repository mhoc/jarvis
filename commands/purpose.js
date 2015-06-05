
var logger = require('log4js').getLogger();

var storedPurpose = "to serve you."

logger.trace('Creating purpose commands');
module.exports = [
  {
    description: "Gives me purpose in life.",

    match: [
      /jarvis your purpose is (.+)/,
      /jarvis your purpose in life is (.+)/,
      /jarvis your new purpose is (.+)/,
      /jarvis your new purpose in life is (.+)/
    ],

    run: function(msg, respond) {
      logger.info('Running set purpose command')
      storedPurpose = msg._matchResult[1]
      respond('Ok, my new purpose in life is ' + storedPurpose, msg.channel)
    }
  },
  {
    description: "Shares my purpose in life with you.",

    match: [
      /jarvis what is your purpose/
    ],

    run: function(msg, respond) {
      logger.info('Running get purpose command')
      respond(storedPurpose, msg.channel)
    }

  }
]
