
var bridge = require('../parsing/msg_cmd_bridge')
var logger = require('log4js').getLogger()
var wsClient = require('websocket').client
var connection = null

// Writes to the socket
var sendMessage = function(message, chan) {
  connection.sendUTF(JSON.stringify({
    id: 1,
    type: "message",
    channel: chan,
    text: message
  }))
}

// Initializes a web socket client to be connected somewhere else
// @returns the web socket client
module.exports = function() {

  logger.trace('Creating websocket client')
  client = new wsClient()

  client.on('connectFailed', function() {
    logger.error('Failed to connect to slack websocke')
    process.exit(1)
  })

  client.on('connect', function(conn) {
    logger.trace('Web socket client has connected')
    connection = conn

    conn.on('error', function(err) {
      logger.error('Connected websocket raised an error')
      logger.error(err.toString())
      process.exit(1)
    })

    conn.on('close', function() {
      logger.info('Server has closed websocket')
    })

    conn.on('message', function(msg) {
      slackMsg = JSON.parse(msg.utf8Data)
      if (slackMsg.type == "message") {
        logger.info("Received '" + slackMsg.text + "'")
        bridge(slackMsg, sendMessage)
      } else {
        logger.info("Received " + slackMsg.type)
      }
    })

  })

  return client

}
