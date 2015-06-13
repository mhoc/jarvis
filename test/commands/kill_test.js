
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
})
