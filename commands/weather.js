
var currentWeather = require('../util/current_weather')
var logger = require('log4js').getLogger()
var zipToLoc = require('../util/zip_to_loc')

module.exports = [
  {
    description: "gives the weather for a specific location you request.",

    match: [
      /jarvis weather (.*)/,
      /jarvis give me the weather for (.*)/
    ],

    run: function(slackMsg, respond) {
      logger.info('Running weather command')
      location = slackMsg._matchResult[1]

      // We will assume its a zip code for now
      zipToLoc(location, function(err, locData) {
        currentWeather(locData.lat, locData.lng, function(err, weatherData) {
          resStr  = weatherData.minutely.summary + "\n"
          resStr += "The current temperature is " + Math.floor(weatherData.currently.temperature) + "°F."
          slackMsg.text = resStr
          respond(slackMsg)
        })
      })

    }

  },
  {

    description: "gives the weather for your current location.",

    match: [
      /jarvis my weather/,
      /jarvis whats the weather like outside/
    ],

    run: function(slackMsg, respond) {

    }

  }
]
