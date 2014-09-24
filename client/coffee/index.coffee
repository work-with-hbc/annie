fs = require 'fs'
Duo = require 'duo'
{join: join} = require('path')
coffeescript = require('coffee-script')


# Compile coffee to js.
coffee = (file, entry) ->
  return unless file.type == 'coffee'

  file.src = coffeescript.compile file.src
  file.type = 'js'

  return file


# Output file.
out = join __dirname, 'annie.js'

duo = (Duo __dirname)
  .development true
  .use coffee
  .entry 'annie.coffee'


# Build script.
duo.run (err, src) ->
  throw err if err

  fs.writeFileSync out, src

  len = Buffer.byteLength src
  console.log "All done! Wrote #{len / 1024 | 0}kb"
