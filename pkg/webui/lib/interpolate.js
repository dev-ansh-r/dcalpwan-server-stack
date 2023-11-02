
/**
 * Takes a templated string and interpolates its values using a values object.
 *
 * @param {string} str - The to be interpolated template string.
 * @param {values} values - The values to interpolate the template string with.
 *
 * @returns {string} - The interpolated string.
 */
const interpolate = (str, values = {}) => str.replace(/\{([^}]+)\}/g, (a, b) => values[b])

export default interpolate
