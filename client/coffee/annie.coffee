if window?
  root = window
else
  root = exports

$ = jquery = require 'jquery'

# Simple HTTP client binding for annie.
Annie =

  # API main endpoint.
  baseUrl: null

  setBaseUrl: (@baseUrl) ->

  getBaseUrl: (part='') -> "#{@baseUrl}/#{part}"


  # Thing API
  thing:

    getThingUrl: (part='') -> Annie.getBaseUrl("thing/#{part}")

    # Try to retrieve thing from annie.
    get: (id, onSuccess, onError) ->
      xhr = $.get @getThingUrl "#{id}"

      xhr.success (data) => onSuccess @parseGetThingResp data
      xhr.fail onError if onError?

    # Store thing to annie.
    store: (thing, onSuccess, onError) ->
      xhr = $.ajax
        url: @getThingUrl()
        method: 'POST'
        contentType: 'application/json'
        data: JSON.stringify thing: thing

      xhr.success (data) => onSuccess @parseSetThingResp data
      xhr.fail onError if onError?

    parseGetThingResp: (data) -> data.value
    parseSetThingResp: (data) -> data.id


root.Annie = Annie
