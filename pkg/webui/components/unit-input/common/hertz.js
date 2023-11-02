

import React from 'react'

import FactoredUnitInput from '../factored'

const units = [
  { label: 'Hz', value: 'hz', factor: 1 },
  { label: 'kHz', value: 'khz', factor: 1000 },
  { label: 'MHz', value: 'mhz', factor: 1000000 },
]

const HertzInput = props => (
  <FactoredUnitInput
    inputWidth="xs"
    units={units}
    baseUnit={units[0].value}
    defaultUnit="mhz"
    {...props}
  />
)

export default HertzInput
