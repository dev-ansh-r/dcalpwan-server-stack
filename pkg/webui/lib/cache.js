

import { selectApplicationRootPath } from '@ttn-lw/lib/selectors/env'
import stringToHash from '@ttn-lw/lib/string-to-hash'

const hash = stringToHash(selectApplicationRootPath())
const hashKey = key => `${key}-${hash}`

export const get = key => {
  const hashedKey = hashKey(key)
  const value = localStorage.getItem(hashedKey)
  try {
    return JSON.parse(value)
  } catch (e) {
    return value
  }
}

export const set = (key, val) => {
  const hashedKey = hashKey(key)
  const value = JSON.stringify(val)
  localStorage.setItem(hashedKey, value)
}

export const remove = key => {
  const hashedKey = hashKey(key)
  return localStorage.removeItem(hashedKey)
}

export const clearAll = () => localStorage.clear()
