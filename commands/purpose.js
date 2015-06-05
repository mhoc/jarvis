
var logger = require('log4js').getLogger();

var storedPurpose = "To serve the citizens of the glorious SIGAPP nation."

logger.trace('Creating purpose commands');
module.exports = [
  {
    description: "Gives me purpose in life.",

    match: [
      /[Jj]arvis your purpose is (.+)/,
      /[Jj]arvis your purpose in life is (.+)/,
      /[Jj]arvis your new purpose is (.+)/,
      /[Jj]arvis your new purpose in life is (.+)/
    ],

    run: function(msg, respond) {
      logger.info('Running set purpose command')
      storedPurpose = msg._matchResult[1]
      msg.text = 'Ok, my new purpose in life is ' + storedPurpose
      respond(msg)
    }
  },
  {
    description: "Shares my purpose in life with you.",

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
