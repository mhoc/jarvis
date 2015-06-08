
// Any deterministic replies where jarvis doesn't need external data.
// Essentially stupid stuff that's just for fun, or help, or links to things, etc etc.
var logger = require('log4js').getLogger()
var request = require('request')
var slack = require('../util/slack')

module.exports = [
  {
    description: "provides a link to a glossary of commands I support.",

    match: [
      /jarvis help/,
      /jarvis help me/,
      /jarvis what commands do you support/,
      /jarvis what can you do/,
    ],

    run: function(msg, respond) {
      logger.info('Running help command')
      msg.text = 'Jarvis, at your service.\nYou can find a full documentation of my capabilities at http://github.com/mhoc/jarvis'
      respond(msg)
    }
  },
  {
    description: "provides information about SIGAPP and links to SIGAPP resources.",

    match: [
      /jarvis sigapp/,
      /jarvis sigapp info/,
      /jarvis give me information about sigapp/
    ],

    run: function(msg, respond) {
      logger.info('Running sigapp info command')
      msg.text =  'SIGAPP meets every Tuesday and Thursday at 7pm, and every Saturday at 4pm.\n'
      msg.text += 'You can find our code at http://github.com/purdue-acm-sigapp\n'
      msg.text += 'and our wiki with in-depth information about everything we do at http://github.com/purdue-acm-sigapp/wiki/wiki.'
      respond(msg)
    }
  },
  {
    description: "tells everyone what I really think about amazon.",

    match: [
      /jarvis what do you think about [Aa]mazon/
    ],

    run: function(msg, respond) {
      logger.info('Running amazon command')
      slack.is_user_online("U02GY0X1R", function(err, isOnline) {
        if (isOnline) {
          msg.text = "Oh its the best company of all time, no doubt about it guys."
        } else {
          msg.text = "Eric isn't here is he? I'm gonna be honest, I think he's a bit in love with that place."
        }
        respond(msg)
      })
    }

  }
]
