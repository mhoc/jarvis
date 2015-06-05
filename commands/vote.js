
var id = require('../util/id')
var logger = require('log4js').getLogger()

// We only allow one vote per slack channel at any time
var activeVotes = {}

module.exports = [
  {
    description: "starts a new poll.",

    match: [
      /jarvis start a poll/,
      /jarvis start a new pool/,
      /jarvis start a vote/
    ],

    run: function(msg, respond) {

      // Check if there is already a vote going on for that channel
      if (activeVotes[msg.channel] != null) {
        msg.text = "There's already a poll in this channel, wait just a bit for it to finish."
        respond(msg)
        return
      }

      activeVotes[msg.channel] = {}
      msg.text  = "I've created a poll for you.\n"
      msg.text += "You can respond by saying 'jarvis vote (your vote)'.\n"
      msg.text += "The poll will end in 5 minutes."

      respond(msg)
    }
  },
  {
    description: "registers a vote on a poll.",

    match: [
      /jarvis vote (.*)/
    ],

    run: function(msg, respond) {
      if (activeV)
      activeVotes[msg._matchResult[1]
    }
  }

]
