

import unicode from 'unicode-properties'

/**
 * Checks if `password` has enough special characters.
 *
 * @param {string} password - The pasword to validate.
 * @param {number} specialCount - The number of required special characters.
 * @returns {boolean} - `true` if `password` has more or equal number of special characters `specialCount`.
 */
export const hasSpecial = (password = '', specialCount = 0) => {
  let foundSpecial = 0
  for (let i = 0; i < password.length && foundSpecial < specialCount; i++) {
    const codePoint = password.codePointAt(i)
    if (!unicode.isAlphabetic(codePoint) && !unicode.isDigit(codePoint)) {
      foundSpecial++
    }
  }

  return foundSpecial >= specialCount
}

/**
 * Checks if `password` has enough uppercase characters.
 *
 * @param {string} password - The pasword to validate.
 * @param {number} upperCount - The number of required uppercase characters.
 * @returns {boolean} - `true` if `password` has more or equal number of uppercase characters `upperCount`.
 */
export const hasUpper = (password = '', upperCount = 0) => {
  let foundUppercase = 0
  for (let i = 0; i < password.length && foundUppercase < upperCount; i++) {
    if (unicode.isUpperCase(password.codePointAt(i))) {
      foundUppercase++
    }
  }

  return foundUppercase >= upperCount
}

/**
 * Checks if `password` has enough digits.
 *
 * @param {string} password - The pasword to validate.
 * @param {number} digitCount - The number of required digits.
 * @returns {boolean} - `true` if `password` has more or equal number of digits `digitCount`.
 */
export const hasDigit = (password = '', digitCount = 0) => {
  let foundDigits = 0
  for (let i = 0; i < password.length && foundDigits < digitCount; i++) {
    if (unicode.isDigit(password.charCodeAt(i))) {
      foundDigits++
    }
  }

  return foundDigits >= digitCount
}

/**
 * Checks if `password` has minimum allowed length.
 *
 * @param {string} password - The pasword to validate.
 * @param {number} minLength - The minimum allowed length.
 * @returns {boolean} - `true` if `password` has more or equal characters`minLength`.
 */
export const hasMinLength = (password = '', minLength = 0) => password.length >= minLength

/**
 * Checks if `password` has minimum allowed length.
 *
 * @param {string} password - The pasword to validate.
 * @param {number} maxLength - The maximum allowed length.
 * @returns {boolean} - `true` if `password` has less or equal characters `maxLength`.
 */
export const hasMaxLength = (password = '', maxLength = Infinity) => password.length <= maxLength
