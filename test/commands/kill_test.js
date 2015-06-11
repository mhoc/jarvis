
var assert = require('assert')
var cmd = require('../../commands/kill')
require('tablog').setLevel(5)

describe('Kill Command', function() {
  describe('Setup', function() {
    it('Command should have description', function() {
      assert(cmd.description)
    })
    it('Command should define at least one match', function() {
      assert(cmd.match.length >= 1)
    })
    it('Command should define a run function', function() {
      assert(cmd.run)
    })
  })
  describe('Run', function() {
    it('Run should fail if user is not authed', function() {
      msg = {'user': '12345'}
      cmd.run(msg, function(){})
    })
  })
})
