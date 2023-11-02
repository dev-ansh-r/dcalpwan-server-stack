

import React from 'react'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { unit as unitRegexp } from '@console/lib/regexp'

import EncodedUnitInput from '../encoded'

const units = [
  { label: sharedMessages.milliseconds, value: 'ms' },
  { label: sharedMessages.seconds, value: 's' },
  { label: sharedMessages.minutes, value: 'm' },
  { label: sharedMessages.hours, value: 'h' },
]

const encoder = (value, unit) => {
  if (value === null) {
    return null
  }

  return value ? `${value}${unit}` : unit
}
const decoder = (rawValue = '') => {
  if (rawValue === null) {
    return { value: '', unit: null }
  }
  const value = rawValue.split(unitRegexp)[0]
  const unit = rawValue.split(value)[1] || null
  return {
    value: value ? Number(value) : undefined,
    unit,
  }
}

const DurationInput = props => (
  <EncodedUnitInput
    inputWidth="xxs"
    selectWidth="s"
    units={units}
    encode={encoder}
    decode={decoder}
    {...props}
  />
)

export default DurationInput
