

/**
 * Function to return a valid input date string from a Date object.
 *
 * @param {object} d - Date() object.
 * @returns {string} 'yyyy-mm-dd' or undefined.
 */

export default d => {
  if (Object.prototype.toString.call(d) !== '[object Date]' || isNaN(d.getTime())) {
    return undefined
  }
  const mm = d.getMonth() + 1
  const dd = d.getDate()
  const yy = d.getFullYear()
  return `${yy}-${`0${mm}`.slice(-2)}-${`0${dd}`.slice(-2)}`
}
