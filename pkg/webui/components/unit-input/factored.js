

import React, { useCallback, useState } from 'react'

import PropTypes from '@ttn-lw/lib/prop-types'

import UnitInput from './unit-input'

const FactoredUnitInput = props => {
  const { units, baseUnit, defaultUnit, value, onChange, ...rest } = props

  const [unit, setUnit] = useState(
    defaultUnit
      ? units.find(e => e.value === defaultUnit)?.value || units[0].value
      : units[0].value,
  )
  const unitIndex = units.findIndex(e => e.value === unit)
  const baseUnitIndex = units.findIndex(e => e.value === baseUnit)
  const currentFactor = 1 / (units[baseUnitIndex].factor / units[unitIndex].factor)
  const displayValue = !isNaN(value) && value !== '' ? value / currentFactor : value

  const handleInputChange = useCallback(
    value => {
      onChange(!isNaN(value) && value !== '' ? value * currentFactor : value)
    },
    [currentFactor, onChange],
  )

  const handleUnitChange = useCallback(
    unit => {
      setUnit(unit)
    },
    [setUnit],
  )

  return (
    <UnitInput
      {...rest}
      onInputChange={handleInputChange}
      onUnitChange={handleUnitChange}
      value={{ value: displayValue, unit }}
      units={units}
    />
  )
}

FactoredUnitInput.propTypes = {
  baseUnit: PropTypes.string.isRequired,
  defaultUnit: PropTypes.string,
  onChange: PropTypes.func.isRequired,
  units: PropTypes.arrayOf(
    PropTypes.shape({
      label: PropTypes.message,
      value: PropTypes.string,
      factor: PropTypes.number,
    }),
  ).isRequired,
  value: PropTypes.oneOfType([PropTypes.number, PropTypes.string]),
}

FactoredUnitInput.defaultProps = {
  defaultUnit: undefined,
  value: undefined,
}

export default FactoredUnitInput
