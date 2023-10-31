

/* eslint-disable no-console */

import dev from './dev'

export const warn = (...args) => {
  if (dev) {
    console.warn(...args)
  }
}

export const error = (...args) => {
  if (dev) {
    console.warn(...args)
  }
}

const log = (...args) => {
  if (dev) {
    console.log(...args)
  }
}

export default log
