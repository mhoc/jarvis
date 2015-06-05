
var logger = require('log4js').getLogger()

logger.trace("Compiling list of installed commands")

path = require('path').join(__dirname, '../commands')
commands = []
require('fs').readdirSync(path).forEach(function(file) {
  cmd = require('../commands/' + file)
  if (Object.prototype.toString.call(cmd) === '[object Array]') {
    commands = commands.concat(cmd)
  } else {
    commands.push(cmd)
  }
})

module.exports = commands
