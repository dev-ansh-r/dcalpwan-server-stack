

/* eslint-disable no-console */
/* eslint-disable import/no-commonjs */

const fs = require('fs')

const api = require('../generated/api.json')
const allowedFieldMaskPaths = require('../generated/allowed-field-mask-paths.json')

const map = files => {
  const result = {}
  const paramRegex = /{([a-z._]+)}/gm

  for (const file of files) {
    for (const service of file.services) {
      result[service.name] = {}

      for (const method of service.methods) {
        result[service.name][method.name] = {
          file: file.name,
          http: [],
          allowedFieldMaskPaths: allowedFieldMaskPaths[`/${service.fullName}/${method.name}`],
        }

        if (method.options && method.options['google.api.http']) {
          for (const rule of method.options['google.api.http'].rules) {
            rule.parameters = []
            let match
            while ((match = paramRegex.exec(rule.pattern)) !== null) {
              rule.parameters.push(match[1])
            }

            rule.method = rule.method.toLowerCase()
            if (method.responseStreaming) {
              rule.stream = method.responseStreaming
            }
            result[service.name][method.name].http.push(rule)
          }
        }
      }
    }
  }

  return result
}

fs.writeFile(
  `${__dirname}/../generated/api-definition.json`,
  JSON.stringify(map(api.files), null, 2),
  err => {
    if (err) {
      return console.error(err)
    }
    console.log('File saved.')
  },
)
