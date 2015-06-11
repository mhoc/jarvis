
var log = require('tablog')
var request = require('request')

// Returns a weather information object for a given latitude and longitude
// @param lat float
// @param lng float
// @param callback function(err, obj)
module.exports = function(lat, lng, callback) {
  log.trace('Making request to weather api for current weather')

  // Read envvar
  apiKey = process.env.DARK_SKY_API_TOKEN
  if (apiKey == null || apiKey == "") {
    log.warn('No API key provided for dark sky, cannot get weather information')
    return
  }

  request({
    url: 'https://api.forecast.io/forecast/' + apiKey + '/' + lat + ',' + lng,
    json: true
  },
    function(err, res, body) {
      if (err != null) {
        log.error('Error contacting dark sky api')
      }
      callback(err, body)
    }
  )

}
