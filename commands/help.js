
var logger = require('log4js').getLogger();

module.exports = {

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

}
