

/**
 * Extracts the value of a cookie if it exists.
 *
 * @param {string} key - The key of the entry to extract.
 * @returns {string} The extracted value.
 */
export default key => {
  const matches = document.cookie.match(
    new RegExp(`(?:^|; )${key.replace(/([.$?*|{}()[\]\\/+^])/g, '\\$1')}=([^;]*)`),
  )

  return matches ? decodeURIComponent(matches[1]) : undefined
}
