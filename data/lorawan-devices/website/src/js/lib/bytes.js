
/**
 * Converts base64 encoded string to hex.
 *
 * @param {string} str - Base64 encoded string.
 * @returns {string} - `str` hex encoded.
 */
export const base64ToHex = (str) => {
  return Array.from(atob(str.replace(/[ \r\n]+$/, '')))
    .map((char) => {
      const tmp = char.charCodeAt(0).toString(16)

      return tmp.length > 1 ? tmp : `0${tmp}`
    })
    .join('')
}
