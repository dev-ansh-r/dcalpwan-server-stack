
import React from 'react'

import Select from '@ttn-lw/components/select'

import PropTypes from '@ttn-lw/lib/prop-types'

import {
  LORAWAN_PHY_VERSIONS,
  parseLorawanMacVersion,
  LORAWAN_VERSION_PAIRS,
} from '@console/lib/device-utils'

const PhyVersionInput = props => {
  const { lorawanVersion, onChange, disabled, value, ...rest } = props

  const lorawanVersionRef = React.useRef(lorawanVersion)
  const [options, setOptions] = React.useState(LORAWAN_PHY_VERSIONS)

  React.useEffect(() => {
    const options =
      LORAWAN_VERSION_PAIRS[parseLorawanMacVersion(lorawanVersion)] || LORAWAN_PHY_VERSIONS
    if (options.length === 1) {
      onChange(options[0].value)
    } else if (lorawanVersion !== lorawanVersionRef.current) {
      lorawanVersionRef.current = lorawanVersion
      onChange('')
    }

    setOptions(options)
  }, [lorawanVersion, onChange])

  return (
    <Select
      options={options}
      onChange={onChange}
      disabled={options.length <= 1 || disabled}
      value={value}
      {...rest}
    />
  )
}

PhyVersionInput.propTypes = {
  disabled: PropTypes.bool,
  frequencyPlan: PropTypes.string,
  lorawanVersion: PropTypes.string,
  onChange: PropTypes.func.isRequired,
  value: PropTypes.string,
}

PhyVersionInput.defaultProps = {
  disabled: false,
  value: undefined,
  frequencyPlan: '',
  lorawanVersion: '',
}

export default PhyVersionInput
