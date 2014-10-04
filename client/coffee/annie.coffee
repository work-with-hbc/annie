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
    getListOfThingUrl: (part='') -> Annie.getBaseUrl("list/thing/#{part}")

    # Retrieve thing from annie with id.
    get: (id, onSuccess, onError) ->
      xhr = $.get @getThingUrl id

      xhr.success (data) => onSuccess @parseGetThingResp data
      xhr.fail onError if onError?

    # Retrieve list from annie with id.
    getList: (id, onSuccess, onError) ->
      xhr = $.get @getListOfThingUrl "#{id}"

      xhr.success (data) => onSuccess @parseGetThingsResp data
      xhr.fail onError if onError?

    # Store thing to annie.
    store: (thing, onSuccess, onError) ->
      xhr = $.ajax
        url: @getThingUrl()
        method: 'POST'
        contentType: 'application/json'
        data: JSON.stringify thing: @prepareThing thing

      xhr.success (data) => onSuccess @parseStoreThingResp data
      xhr.fail onError if onError?

    # Store something to annie with id
    storeWithId: (id, thing, onSuccess, onError) ->
      xhr = $.ajax
        url: @getThingUrl id
        method: 'PUT'
        contentType: 'application/json'
        data: JSON.stringify thing: @prepareThing thing

      xhr.success (data) => onSuccess @parseStoreThingResp data
      xhr.fail onError if onError?

    # Store a list to annie.
    storeList: (things, onSuccess, onError) ->
      xhr = $.ajax
        url: @getListOfThingUrl()
        method: 'POST'
        contentType: 'application/json'
        data: JSON.stringify things: (@prepareThing thing for thing in things)

      xhr.success (data) => onSuccess @parseStoreThingResp data
      xhr.fail onError if onError?

    # Store a list to annie with id.
    storeListWithId: (id, things, onSuccess, onError) ->
      xhr = $.ajax
        url: @getListOfThingUrl id
        method: 'PUT'
        contentType: 'application/json'
        data: JSON.stringify things: (@prepareThing thing for thing in things)

      xhr.success (data) => onSuccess @parseStoreThingResp data
      xhr.fail onError if onError?

    # Pust thing to list to annie.
    pushThingToList: (id, thing, onSuccess, onError) ->
      xhr = $.ajax
        url: @getListOfThingUrl "#{id}/item"
        method: 'PUT'
        contentType: 'application/json'
        data: JSON.stringify thing: @prepareThing thing

      xhr.success (data) => onSuccess @parseStoreThingResp data
      xhr.fail onError if onError?

    prepareThing: (thing) -> JSON.stringify thing
    parseStoreThingResp: (data) -> data.id
    parseGetThingResp: (data) -> JSON.parse data.value
    parseGetThingsResp: (data) ->
      return unless data.value?
      JSON.parse thing for thing in data.value


root.Annie = Annie
