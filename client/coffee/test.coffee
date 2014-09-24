should = require 'should'

{Annie} = require './annie'

describe 'Annie basic usage', ->

  describe 'base url', ->
    it 'should return given url after set base url', ->
      baseUrl = 'http://annie.test'
      
      Annie.setBaseUrl baseUrl
      baseUrl.should.be.exactly Annie.getBaseUrl()

describe 'Annie thing API', ->

  describe 'get a thing', ->

  describe 'remember a thing', ->
