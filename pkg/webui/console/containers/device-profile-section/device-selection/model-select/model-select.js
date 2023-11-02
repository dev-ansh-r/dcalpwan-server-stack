

import React from 'react'
import { defineMessages, useIntl } from 'react-intl'

import Field from '@ttn-lw/components/form/field'
import Select from '@ttn-lw/components/select'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import { SELECT_OTHER_OPTION } from '@console/lib/device-utils'

const m = defineMessages({
  noOptionsMessage: 'No matching model found',
})

const formatOptions = (models = []) =>
  models
    .map(model => ({
      value: model.model_id,
      label: model.name,
    }))
    .concat([{ value: SELECT_OTHER_OPTION, label: sharedMessages.otherOption }])

const ModelSelect = props => {
  const { appId, brandId, name, error, fetching, models, listModels, onChange, ...rest } = props
  const { formatMessage } = useIntl()

  React.useEffect(() => {
    listModels(appId, brandId, {}, [
      'name',
      'description',
      'firmware_versions',
      'hardware_versions',
      'key_provisioning',
      'photos',
      'product_url',
      'datasheet_url',
    ])
  }, [appId, brandId, listModels])

  const options = React.useMemo(() => formatOptions(models), [models])
  const handleNoOptions = React.useCallback(
    () => formatMessage(m.noOptionsMessage),
    [formatMessage],
  )

  return (
    <Field
      {...rest}
      options={options}
      name={name}
      title={sharedMessages.model}
      component={Select}
      isLoading={fetching}
      warning={Boolean(error) ? sharedMessages.endDeviceModelsUnavailable : undefined}
      onChange={onChange}
      noOptionsMessage={handleNoOptions}
      placeholder={sharedMessages.typeToSearch}
      autoFocus
    />
  )
}

ModelSelect.propTypes = {
  appId: PropTypes.string.isRequired,
  brandId: PropTypes.string.isRequired,
  error: PropTypes.error,
  fetching: PropTypes.bool,
  listModels: PropTypes.func.isRequired,
  models: PropTypes.arrayOf(
    PropTypes.shape({
      model_id: PropTypes.string.isRequired,
      name: PropTypes.string.isRequired,
    }),
  ),
  name: PropTypes.string.isRequired,
  onChange: PropTypes.func,
}

ModelSelect.defaultProps = {
  error: undefined,
  fetching: false,
  models: [],
  onChange: () => null,
}

export default ModelSelect
