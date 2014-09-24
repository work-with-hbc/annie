root = exports ? window

$ = jquery = require 'components/jquery'

# Simple HTTP client binding for annie.
Annie =
  thing:
    get: (id) ->
      console.log $
    store: (thing) ->
      console.log $


root.Annie = Annie
