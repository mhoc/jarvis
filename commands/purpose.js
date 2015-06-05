
var logger = require('log4js').getLogger();

logger.trace('Creating purpose commands');

module.exports = [
  {
    description: "Gives me purpose in life.",

    match: [
      /jarvis your purpose is to (.*)/
    ],

    run: function(msg, respond) {
      logger.info('Running set purpose command')
    }
  },
  {
    description: "Shares my purpose in life with you.",

    match: [
      /jarvis what is your purpose/
    ],

    run: function(msg, respond) {
      logger.info('Running get purpose command')
    }

  }
]
