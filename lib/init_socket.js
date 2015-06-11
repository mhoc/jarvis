
var bridge = require('./cmd_matcher')
var log = require('tablog')
var wsClient = require('websocket').client
var connection = null

// Writes to the socket
var sendMessage = function(msg) {
  connection.sendUTF(JSON.stringify({
    id: 1,
    type: "message",
    channel: msg.channel,
    text: msg.text
  }))
}

// Initializes a web socket client to be connected somewhere else
// @returns the web socket client
module.exports = function() {

  log.info('Creating websocket client')
  client = new wsClient()

  client.on('connectFailed', function() {
    log.fatal('Failed to connect to slack websocke')
  })

  client.on('connect', function(conn) {
    log.info('Web socket client has connected')
    connection = conn

    conn.on('error', function(err) {
      log.error('Connected websocket raised an error')
      log.error(err.toString())
      process.exit(1)
    })

    conn.on('close', function() {
      log.info('Server has closed websocket')
    })

    conn.on('message', function(msg) {
      slackMsg = JSON.parse(msg.utf8Data)
      if (slackMsg.type === "message" && slackMsg.text) {
        log.info("Received '" + slackMsg.text + "'")
        bridge(slackMsg, sendMessage)
      } else {
        log.info("Received " + slackMsg.type)
      }
    })

  })

  return client

}
