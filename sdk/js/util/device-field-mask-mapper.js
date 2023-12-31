

/* eslint-disable no-console */
/* eslint-disable import/no-commonjs */
/* eslint-disable no-invalid-this */

const fs = require('fs')

const traverse = require('traverse')

const allowedPaths = require('../generated/allowed-field-mask-paths.json')

const fieldMasks = {
  is: {
    get: allowedPaths['/ttn.lorawan.v3.EndDeviceRegistry/Get'],
    set: allowedPaths['/ttn.lorawan.v3.EndDeviceRegistry/Update'],
  },
  as: {
    get: allowedPaths['/ttn.lorawan.v3.AsEndDeviceRegistry/Get'],
    set: allowedPaths['/ttn.lorawan.v3.AsEndDeviceRegistry/Set'],
  },
  ns: {
    get: allowedPaths['/ttn.lorawan.v3.NsEndDeviceRegistry/Get'],
    set: allowedPaths['/ttn.lorawan.v3.NsEndDeviceRegistry/Set'],
  },
  js: {
    get: allowedPaths['/ttn.lorawan.v3.JsEndDeviceRegistry/Get'],
    set: allowedPaths['/ttn.lorawan.v3.JsEndDeviceRegistry/Set'],
  },
}

const result = {}

for (const component in fieldMasks) {
  // Write get components.
  for (const fieldMask of fieldMasks[component].get) {
    const path = [...fieldMask.split('.'), '_root']
    const val = traverse(result).get(path)
    if (val) {
      if (typeof val[0] === 'string') {
        traverse(result).set(path, [[val[0], component], val[1]])
      } else {
        traverse(result).set(path, [[...val[0], component], val[1]])
      }
    } else {
      traverse(result).set(path, [component, 'read_only'])
    }
  }

  // Write set components.
  for (const fieldMask of fieldMasks[component].set) {
    const path = [...fieldMask.split('.'), '_root']
    const val = traverse(result).get(path)
    if (val) {
      if (typeof val[1] === 'string') {
        traverse(result).set(path, [
          val[0],
          val[1] !== 'read_only' ? [val[1], component] : component,
        ])
      } else {
        traverse(result).set(path, [val[0], [...val[1], component]])
      }
    } else {
      traverse(result).set(path, [component, component])
    }
  }
}

// Rewrite single `_root` entries as plain array leaf.
traverse(result).forEach(function () {
  if (Object.keys(this.node).length === 1 && this.node._root) {
    this.update(this.node._root)
  }
})

fs.writeFile(
  `${__dirname}/../generated/device-entity-map.json`,
  JSON.stringify(result, null, 2),
  err => {
    if (err) {
      return console.error(err)
    }
    console.log('File saved.')
  },
)
