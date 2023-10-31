

class Bytes {
  /**
   * Converts hex encoded string to base64.
   *
   * @param {string} str - Hex encoded string.
   * @returns {string} - `str` base64 encoded.
   */
  static hexToBase64(str) {
    return btoa(
      String.fromCharCode.apply(
        null,
        str
          .replace(/\r|\n/g, '')
          .replace(/([\da-fA-F]{2}) ?/g, '0x$1 ')
          .replace(/ +$/, '')
          .split(' '),
      ),
    )
  }

  static base64ToHex(str) {
    /**
     * Converts base64 encoded string to hex.
     *
     * @param {string} str - Base64 encoded string.
     * @returns {string} - `str` hex encoded.
     */
    return Array.from(atob(str.replace(/[ \r\n]+$/, '')))
      .map(char => {
        const tmp = char.charCodeAt(0).toString(16)
        return tmp.length > 1 ? tmp : `0${tmp}`
      })
      .join('')
  }
}

export default Bytes
