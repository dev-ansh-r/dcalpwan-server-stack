

/**
 * @example
 * const object = {
 *  a: 'b',
 *  b: {
 *    a: 'c',
 *    c: 'd',
 *  },
 * }
 *
 * const resultingObject = omitDeep(object, 'a') // result is { b: { c: 'd' } }
 * @param {object|Array} value - Multinested object or array.
 * @param {string[]} key - Array of properties / keys to exclude from `value`.
 * @param {string[]} regexp - Regular expression for excluding properties from `value`.
 * @returns {object|Array} The ininew object/array without specified properties/keys.
 */
const omitDeep = (value, key, regexp = /(?!)/) => {
  if (Array.isArray(value)) {
    return value.map(i => omitDeep(i, key))
  } else if (typeof value === 'object' && value !== null) {
    return Object.keys(value).reduce((newObject, k) => {
      if (key.includes(k) || regexp.test(k)) return newObject
      return Object.assign({ [k]: omitDeep(value[k], key) }, newObject)
    }, {})
  }
  return value
}

export default omitDeep
