
var logger = require('log4js').getLogger();

logger.trace('Creating help command');

module.exports = {

  description: "Provides a link to a glossary of commands I support.",

  match: [
    /[Jj]arvis help/,
    /[Jj]arvis help me/,
    /[Jj]arvis what commands do you support/,
    /[Jj]arvis what can you do/,
  ],

  run: function(msg, respond) {
    logger.info('Running help command')
    msg.text = 'Jarvis, at your service.\nYou can find a full documentation of my capabilities at http://github.com/mhoc/jarvis'
    respond(msg)
  }

}
