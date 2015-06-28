
// Function which logs the content of a message from slack to the command line
// through tablog. Aims to provide a read-only

var tablog = require('tablog')

module.exports = function(message) {

    if (message.type === "message" && message.text) {
      tablog.info(message.user + ": " + message.text)
    }
    else {
      tablog.info("Received " + message.type)
    }

}
