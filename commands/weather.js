
var logger = require('log4js').getLogger()

module.exports = [
  {
    description: "gives the weather for a specific location you request.",

    match: [
      /jarvis weather (.*)/,
      /jarvis give me the weather for (.*)/
    ],

    run: function(slackMsg, respond) {

    }

  },
  {

    description: "gives the weather for your current location.",

    match: [
      /jarvis my weather/,
      /jarvis whats the weather like outside/
    ],

    run: function(slackMsg, respond) {

    }

  }
]
