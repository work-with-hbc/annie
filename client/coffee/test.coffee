should = require 'should'

{Annie} = require './annie'

describe 'Annie basic usage', ->

  describe 'base url', ->
    it 'should return given url after set base url', ->
      baseUrl = 'http://annie.test/api/v1'
      
      Annie.setBaseUrl baseUrl
      Annie.getBaseUrl().should.be.exactly "#{baseUrl}/"

describe 'Annie thing API', ->

  describe 'get a thing', ->

  describe 'remember a thing', ->
