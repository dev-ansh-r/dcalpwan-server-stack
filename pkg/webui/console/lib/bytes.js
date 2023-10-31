

/**
 * Converts hex encoded string to base64.
 *
 * @param {string} str - Hex encoded string.
 * @returns {string} - `str` base64 encoded.
 */
export const hexToBase64 = str =>
  btoa(
    String.fromCharCode.apply(
      null,
      str
        .replace(/\r|\n/g, '')
        .replace(/([\da-fA-F]{2}) ?/g, '0x$1 ')
        .replace(/ +$/, '')
        .split(' '),
    ),
  )

/**
 * Converts base64 encoded string to hex.
 *
 * @param {string} str - Base64 encoded string.
 * @returns {string} - `str` hex encoded.
 */
export const base64ToHex = str =>
  Array.from(atob(str.replace(/[ \r\n]+$/, '')))
    .map(char => {
      const tmp = char.charCodeAt(0).toString(16)

      return tmp.length > 1 ? tmp : `0${tmp}`
    })
    .join('')
