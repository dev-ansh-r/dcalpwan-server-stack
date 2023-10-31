

import React from 'react'
import { defineMessages } from 'react-intl'

import Select from '@ttn-lw/components/select'
import Field from '@ttn-lw/components/form/field'

import { getOrganizationId, getUserId } from '@ttn-lw/lib/selectors/id'
import PropTypes from '@ttn-lw/lib/prop-types'

const m = defineMessages({
  title: 'Owner',
  warning: 'There was an error and the list of organizations could not be displayed',
})

const OwnersSelect = props => {
  const {
    autoFocus,
    error,
    fetching,
    menuPlacement,
    name,
    onChange,
    organizations,
    required,
    user,
  } = props

  const options = React.useMemo(() => {
    const usrOption = { label: getUserId(user), value: getUserId(user) }
    const orgsOptions = organizations.map(org => ({
      label: getOrganizationId(org),
      value: getOrganizationId(org),
    }))

    return [usrOption, ...orgsOptions]
  }, [user, organizations])
  const handleChange = React.useCallback(
    value => {
      onChange(options.find(option => option.value === value))
    },
    [onChange, options],
  )

  // Do not show the input when there are no alternative options.
  if (options.length === 1) {
    return null
  }

  return (
    <Field
      component={Select}
      options={options}
      name={name}
      required={required}
      title={m.title}
      autoFocus={autoFocus}
      isLoading={fetching}
      warning={Boolean(error) ? m.warning : undefined}
      menuPlacement={menuPlacement}
      onChange={handleChange}
      defaultValue={options[0]}
    />
  )
}

OwnersSelect.propTypes = {
  autoFocus: PropTypes.bool,
  error: PropTypes.error,
  fetching: PropTypes.bool,
  menuPlacement: PropTypes.oneOf(['top', 'bottom', 'auto']),
  name: PropTypes.string.isRequired,
  onChange: PropTypes.func,
  organizations: PropTypes.arrayOf(PropTypes.organization).isRequired,
  required: PropTypes.bool,
  user: PropTypes.user.isRequired,
}

OwnersSelect.defaultProps = {
  autoFocus: false,
  error: undefined,
  fetching: false,
  onChange: () => null,
  menuPlacement: 'auto',
  required: false,
}

export default OwnersSelect
