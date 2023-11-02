

import React from 'react'
import { defineMessages } from 'react-intl'

import Field from '@ttn-lw/components/form/field'
import Select from '@ttn-lw/components/select'
import { useFormContext } from '@ttn-lw/components/form'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import { SELECT_OTHER_OPTION, SELECT_UNKNOWN_HW_OPTION } from '@console/lib/device-utils'

const m = defineMessages({
  title: 'Hardware Ver.',
})

const formatOptions = (versions = []) =>
  versions
    .map(version => ({
      value: version.version,
      label: version.version,
    }))
    .concat([{ value: SELECT_OTHER_OPTION, label: sharedMessages.otherOption }])

const HardwareVersionSelect = props => {
  const { name, versions, onChange, ...rest } = props
  const { setFieldValue } = useFormContext()

  const options = React.useMemo(() => {
    const opts = formatOptions(versions)
    // When only the `Other...` option is available (so end device model has no hw versions defined
    // in the device repository) add another pseudo option that represents absence of hw versions.
    if (opts.length === 1) {
      opts.unshift({ value: SELECT_UNKNOWN_HW_OPTION, label: sharedMessages.unknownHwOption })
    }

    return opts
  }, [versions])

  React.useEffect(() => {
    if (options.length > 0 && options.length <= 2) {
      setFieldValue('version_ids.hardware_version', options[0].value)
    }
  }, [options, setFieldValue])

  return (
    <Field {...rest} options={options} name={name} title={m.title} component={Select} autoFocus />
  )
}

HardwareVersionSelect.propTypes = {
  name: PropTypes.string.isRequired,
  onChange: PropTypes.func,
  versions: PropTypes.arrayOf(
    PropTypes.shape({
      version: PropTypes.string.isRequired,
    }),
  ),
}

HardwareVersionSelect.defaultProps = {
  versions: [],
  onChange: () => null,
}

export default HardwareVersionSelect
