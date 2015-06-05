
var logger = require('log4js').getLogger()

// Generates a new id of a given length for use in whatever a command
// might need it to be used in.
module.exports = function(length) {
  id = ""
  for (var i = 0; i < length; i++) {
    id += String.fromCharCode(Math.random() * (122-97) + 97)
  }
  return id
}
