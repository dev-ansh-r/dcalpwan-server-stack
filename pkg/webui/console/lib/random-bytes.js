

/**
 * Generates a random string of the given length.
 * @param {number} len - The length of the string to generate.
 * @param {string} type - The type of the string to generate.
 * @returns {string} The generated string.
 * @example
 * const randomString = randomBytes(16)
 * @example
 * const randomString = randomBytes(16, 'base64')
 */
export default (len, type = 'hex') => {
  let bytes
  if (window.crypto) {
    bytes = crypto.getRandomValues(new Uint8Array(Math.floor(len / 2)))
  } else {
    const byteLength = Math.floor(len / 2)
    bytes = new Uint8Array(byteLength)
    for (let i = 0; i < byteLength; i++) {
      bytes[i] = Math.floor(Math.random() * 256)
    }
  }
  switch (type) {
    case 'base64':
      return btoa(String.fromCharCode(...bytes))
    case 'hex':
    default:
      return Array.from(bytes, byte => byte.toString(16).padStart(2, '0'))
        .join('')
        .toUpperCase()
  }
}
