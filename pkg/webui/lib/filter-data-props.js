
/**
 * @example
 * const props = {
 *  className: 'objClass',
 *  counter: 3,
 *  'data-test-id': 'object-test',
 * }
 *
 * const testOnly = filterDataProps(props) // result is {'data-test-id': 'object-test'}
 *
 * @param {object} props - Multinested object.
 * @returns {object} New object with input object properties that start with `data`.
 */
export default props =>
  Object.keys(props)
    .filter(key => key.startsWith('data-'))
    .reduce((acc, key) => {
      acc[key] = props[key]
      return acc
    }, {})
