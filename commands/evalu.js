
var config = require('../config')
var log = require('tablog')

module.exports = {

  description: "executes arbitrary javascript. This isn't a security nightmare at all, right?",

  match: [
    /jarvis evaluate (.+)/
  ],

  run: function(msg, respond) {
    log.trace('Running evaluate command')
    if (msg._matchResult[1].indexOf("exit") > -1) {
      msg.text = "I'm not a fan of letting you execute anything with the word 'exit' in it. Sorry."
      respond(msg)
      return
    }
    try {
      msg.text = "" + eval(msg._matchResult[1])
    } catch (err) {
      msg.text = "Nice try crashing me fucktard."
    }
    respond(msg)
  }

}
