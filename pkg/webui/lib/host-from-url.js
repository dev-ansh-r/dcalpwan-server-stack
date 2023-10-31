

/**
 * Extracts hostname from the `url`.
 *
 * @param {string} url - The URL string.
 * @returns {string?} - The hostname of the `url` or undefined.
 */
export default url => {
  try {
    if (url.match(/^[a-zA-Z0-9]+:\/\/.*/)) {
      return new URL(url).hostname
    }

    return new URL(`http://${url}`).hostname
  } catch (error) {
    return url
  }
}
