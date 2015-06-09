
var logger = require('log4js').getLogger()

module.exports = [

  {
    description: "orders jarvis to kill the node server he is running on.",

    match: [
      /jarvis kill yourself/
    ],

    run: function(msg, respond) {
      
    }

  }

]
