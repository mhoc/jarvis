
var config = require('../config')
var log = require('tablog')
var request = require('request')

module.exports = {

  description: "displays an excellent quote of my chosing. Quotes are great!",

  match: [
    /jarvis quote/
  ],

  run: function(msg, respond) {
    log.trace('Running quote command')
    request(config.quote_url, function(err, res, body) {
      quotes = body.split("\n")
      msg.text = quotes[Math.floor(Math.random() * quotes.length)]
      respond(msg)
    })
  }

}
