

import React from 'react'
import { defineMessages } from 'react-intl'
import { useSelector } from 'react-redux'

import Field from '@ttn-lw/components/form/field'
import Select from '@ttn-lw/components/select'
import { useFormContext } from '@ttn-lw/components/form'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import { SELECT_OTHER_OPTION } from '@console/lib/device-utils'

import { selectDeviceModelFirmwareVersions } from '@console/store/selectors/device-repository'

const m = defineMessages({
  title: 'Profile (Region)',
})

const formatOptions = (profiles = []) =>
  profiles
    .map(profile => ({
      value: profile,
      label: profile,
    }))
    .concat([{ value: SELECT_OTHER_OPTION, label: sharedMessages.otherOption }])

const BandSelect = props => {
  const { name, onChange, brandId, modelId, fwVersion, ...rest } = props
  const { setFieldValue } = useFormContext()
  const versions = useSelector(state => selectDeviceModelFirmwareVersions(state, brandId, modelId))
  const version = versions.find(v => v.version === fwVersion) || { profiles: [] }
  const profiles = Object.keys(version.profiles)

  const options = React.useMemo(() => formatOptions(profiles), [profiles])
  const onlyOption = options.length > 0 && options.length <= 2 ? options[0].value : undefined

  React.useEffect(() => {
    if (onlyOption) {
      setFieldValue('version_ids.band_id', onlyOption)
    }
  }, [onlyOption, setFieldValue])

  return (
    <Field {...rest} options={options} name={name} title={m.title} component={Select} autoFocus />
  )
}

BandSelect.propTypes = {
  brandId: PropTypes.string.isRequired,
  fwVersion: PropTypes.string.isRequired,
  modelId: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  onChange: PropTypes.func,
}

BandSelect.defaultProps = {
  onChange: () => null,
}

export default BandSelect
