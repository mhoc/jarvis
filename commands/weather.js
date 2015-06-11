
var currentWeather = require('../util/current_weather')
var logger = require('log4js').getLogger()
var zipToLoc = require('../util/zip_to_loc')

module.exports = [
  {
    description: "gives the weather for a specific location you request.",

    match: [
      /jarvis weather (.*)/,
      /jarvis weather for (.*)/,
      /jarvis weather at (.*)/,
      /jarvis give me the weather for (.*)/
    ],

    run: function(slackMsg, respond) {
      logger.info('Running weather command')
      location = slackMsg._matchResult[1]

      // We will assume its a zip code for now
      zipToLoc(location, function(err, locData) {
        currentWeather(locData.lat, locData.lng, function(err, weatherData) {
          if (err || !weatherData.minutely || !weatherData.currently) {
            slackMsg.text = "I'm unable to correctly contact the service I use for weather data."
            respond(slackMsg)
            return
          }
          resStr  = weatherData.minutely.summary + "\n"
          resStr += "The current temperature is " + Math.floor(weatherData.currently.temperature) + "°F."
          slackMsg.text = resStr
          respond(slackMsg)
        })
      })

    }
  }
]
