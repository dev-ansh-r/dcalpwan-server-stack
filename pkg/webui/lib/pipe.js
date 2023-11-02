

/**
 * Performs left-to-right function composition.
 *
 * @param {Function[]}funcs - A list of functions.
 * @returns {Function} A single function composed from `funcs`.
 */
const pipe = (...funcs) =>
  funcs.reduce(
    (a, b) =>
      (...args) =>
        a(b(...args)),
    arg => arg,
  )

export default pipe
