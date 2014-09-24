root = exports ? window.Annie

$ = jquery = require 'jquery'

# Simple HTTP client binding for annie.
Annie =

  # API main endpoint.
  baseUrl: null

  setBaseUrl: (@baseUrl) ->

  getBaseUrl: -> @baseUrl

  thing:

    get: (id) ->
      console.log $

    store: (thing) ->
      console.log $


root.Annie = Annie
