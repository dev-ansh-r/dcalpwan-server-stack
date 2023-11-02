

import React, { useCallback, useState } from 'react'

import PropTypes from '@ttn-lw/lib/prop-types'

import UnitInput from './unit-input'

const EncodedUnitInput = props => {
  const { onChange, encode, decode, value, units, ...rest } = props

  const decodedValue = decode(value)
  const { unit } = decodedValue
  const [storedUnit, setStoredUnit] = useState(unit || units[0].value)
  decodedValue.unit = decodedValue.unit || storedUnit

  const handleInputChange = useCallback(
    inputValue => {
      onChange(encode(inputValue, unit || storedUnit))
    },
    [onChange, encode, unit, storedUnit],
  )

  const handleUnitChange = useCallback(
    unit => {
      setStoredUnit(unit)
      onChange(encode(decodedValue.value, unit), true)
    },
    [onChange, encode, decodedValue.value],
  )

  return (
    <UnitInput
      {...rest}
      onInputChange={handleInputChange}
      onUnitChange={handleUnitChange}
      value={decodedValue}
      storedUnit={storedUnit}
      units={units}
    />
  )
}

EncodedUnitInput.propTypes = {
  decode: PropTypes.func.isRequired,
  encode: PropTypes.func.isRequired,
  onChange: PropTypes.func.isRequired,
  units: PropTypes.arrayOf(
    PropTypes.shape({
      label: PropTypes.message,
      value: PropTypes.string,
      factor: PropTypes.number,
    }),
  ).isRequired,
  value: PropTypes.string,
}

EncodedUnitInput.defaultProps = {
  value: undefined,
}

export default EncodedUnitInput
