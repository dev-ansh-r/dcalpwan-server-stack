

import React from 'react'
import { defineMessages, useIntl } from 'react-intl'

import Field from '@ttn-lw/components/form/field'
import Select from '@ttn-lw/components/select'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import { SELECT_OTHER_OPTION } from '@console/lib/device-utils'

const m = defineMessages({
  title: 'End device brand',
  noOptionsMessage: 'No matching brand found',
})

const formatOptions = (brands = []) =>
  brands
    .map(brand => ({
      value: brand.brand_id,
      label: brand.name || brand.brand_id,
      profileID: brand.brand_id,
    }))
    .concat([{ value: SELECT_OTHER_OPTION, label: sharedMessages.otherOption }])

const BrandSelect = props => {
  const { appId, name, error, fetching, brands, onChange, ...rest } = props
  const { formatMessage } = useIntl()

  const options = React.useMemo(() => formatOptions(brands), [brands])
  const handleNoOptions = React.useCallback(
    () => formatMessage(m.noOptionsMessage),
    [formatMessage],
  )

  return (
    <Field
      {...rest}
      options={options}
      name={name}
      title={m.title}
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

BrandSelect.propTypes = {
  appId: PropTypes.string.isRequired,
  brands: PropTypes.arrayOf(
    PropTypes.shape({
      brand_id: PropTypes.string.isRequired,
    }),
  ),
  error: PropTypes.error,
  fetching: PropTypes.bool,
  name: PropTypes.string.isRequired,
  onChange: PropTypes.func,
}

BrandSelect.defaultProps = {
  error: undefined,
  fetching: false,
  brands: [],
  onChange: () => null,
}

export default BrandSelect
