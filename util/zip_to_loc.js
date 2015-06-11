
var log = require('tablog')
var request = require('request')

// Converts a zip code into a more specific location object using the zipcodeapi
// @arg zipcode the zip code you want to covert (string)
// @arg callback function(err, obj)
//
// obj: contains fields like "lat", "lng", "city", "state", "timezone.timezone_abbr"
// See full documentation on http://zipcodeapi.com
module.exports = function(zipcode, callback) {
  log.trace('Making request to zip code api for location object')

  // Read the environ
  var apiKey = process.env.ZIP_CODE_API_TOKEN
  if (apiKey == null || apiKey === "") {
    log.warn('Attempted to access zip code api but no api key set')
    return
  }

  // Make the request to the zip code api
  request({
    url: 'https://www.zipcodeapi.com/rest/' + apiKey + '/info.json/' + zipcode + '/degrees',
    json: true
  },
    function(err, res, body) {
      if (err != null) {
        log.error('Error contacting zipcode api')
      }
      callback(err, body)
  })

}
