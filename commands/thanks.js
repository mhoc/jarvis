
var log = require('tablog')

var possibleResponses = [
  "You got it, buddy.",
  "No problem man.",
  "Its my pleasure."
]

module.exports = {
  description: "thanks people for thanking jarvis!",

  match: [
    /thanks [Jj]arvis/,
    /thank you [Jj]arvis/
  ],

  run: function(msg, respond) {
    log.trace("Running thanks function")
    msg.text = possibleResponses[Math.floor(Math.random() * possibleResponses.length)]
    respond(msg)
  }

}
