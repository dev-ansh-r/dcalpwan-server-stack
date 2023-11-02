

import toInputDate from '@ttn-lw/lib/to-input-date'

export const encodeExpiryDate = value => {
  if (value) {
    return new Date(value).toISOString()
  }

  return null
}

export const decodeExpiryDate = value => {
  if (value) {
    return toInputDate(new Date(value))
  }

  return ''
}
