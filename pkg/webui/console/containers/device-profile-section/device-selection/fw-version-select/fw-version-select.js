

import React from 'react'
import { defineMessages } from 'react-intl'

import Field from '@ttn-lw/components/form/field'
import Select from '@ttn-lw/components/select'
import { useFormContext } from '@ttn-lw/components/form'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import { SELECT_OTHER_OPTION } from '@console/lib/device-utils'

const m = defineMessages({
  title: 'Firmware Ver.',
})

const formatOptions = (versions = []) =>
  versions
    .map(version => ({
      value: version.version,
      label: version.version,
    }))
    .concat([{ value: SELECT_OTHER_OPTION, label: sharedMessages.otherOption }])

const FirmwareVersionSelect = props => {
  const { name, versions, onChange, ...rest } = props
  const { setFieldValue } = useFormContext()

  const options = React.useMemo(() => formatOptions(versions), [versions])

  React.useEffect(() => {
    if (options.length > 0 && options.length <= 2) {
      setFieldValue('version_ids.firmware_version', options[0].value)
    }
  }, [setFieldValue, options])

  return (
    <Field {...rest} options={options} name={name} title={m.title} component={Select} autoFocus />
  )
}

FirmwareVersionSelect.propTypes = {
  name: PropTypes.string.isRequired,
  onChange: PropTypes.func,
  versions: PropTypes.arrayOf(
    PropTypes.shape({
      version: PropTypes.string.isRequired,
    }),
  ),
}

FirmwareVersionSelect.defaultProps = {
  versions: [],
  onChange: () => null,
}

export default FirmwareVersionSelect
