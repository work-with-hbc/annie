Annie.setBaseUrl 'http://127.0.0.1:1234/api/v1'

testThing = ->
  console.log 'test thing api'

  testValue = 'test value'
  Annie.thing.store testValue, (id) ->
    Annie.thing.get id, (value) ->
      console.log "expected #{testValue}, got #{value}"

  testKey = "test-key#{Math.random()}"
  Annie.thing.storeWithId testKey, testValue, (id) ->
    console.log "expected #{testKey}, got #{id}"
    
    Annie.thing.get id, (value) ->
      console.log "expected #{testValue}, got #{value}"


testList = ->
  console.log 'test list of thing api'

  testValues = ['test', 'value']
  Annie.thing.storeList testValues, (id) ->
    Annie.thing.getList id, (values) ->
      console.log "expected #{testValues}, got #{values}"

  testKey = "test-key#{Math.random()}"
  Annie.thing.storeListWithId testKey, testValues, (id) ->
    console.log "expected #{testKey}, got #{id}"
    
    Annie.thing.getList id, (values) ->
      console.log "expected #{testValues}, got #{values}"

testPushList = ->
  testValues = ['test', 'value']
  testKey = "test-key#{Math.random()}"
  pushed = 'foobar'

  Annie.thing.storeListWithId testKey, testValues, (id) ->

  testValues.push pushed
  Annie.thing.pushThingToList testKey, pushed, (id) ->
    console.log "expected #{testKey}, got #{id}"
    
    Annie.thing.getList id, (values) ->
      console.log "expected #{testValues}, got #{values}"


testThing()
testList()
testPushList()
