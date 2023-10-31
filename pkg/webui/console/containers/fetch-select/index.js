

import React, { useCallback, useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'

import Field from '@ttn-lw/components/form/field'
import Select from '@ttn-lw/components/select'

import PropTypes from '@ttn-lw/lib/prop-types'

const formatOptions = options =>
  Object.keys(options).map(key => ({ value: key, label: options[key] }))

const { component, ...fieldPropTypes } = Field.propTypes

export default ({
  optionsSelector,
  errorSelector,
  fetchingSelector,
  fetchOptions,
  defaultWarning,
  defaultTitle,
  optionsFormatter = formatOptions,
  defaultDescription,
  additionalOptions = [],
}) => {
  const FetchSelect = props => {
    const { warning, onChange, ...rest } = props
    const options = [...optionsFormatter(useSelector(optionsSelector)), ...additionalOptions]
    const error = useSelector(errorSelector)
    const fetching = useSelector(fetchingSelector)
    const dispatch = useDispatch()

    useEffect(() => {
      dispatch(fetchOptions())
    }, [dispatch])

    const handleChange = useCallback(
      value => {
        const selectedOption = options.find(option => option.value === value)
        onChange(selectedOption)
      },
      [onChange, options],
    )

    return (
      <Field
        {...rest}
        options={options}
        component={Select}
        isLoading={fetching}
        warning={Boolean(error) ? defaultWarning : warning}
        onChange={handleChange}
      />
    )
  }

  FetchSelect.propTypes = {
    ...fieldPropTypes,
    ...Select.propTypes,
    defaultWarning: PropTypes.message,
    description: PropTypes.message,
    fetchOptions: PropTypes.func.isRequired,
    menuPlacement: PropTypes.oneOf(['top', 'bottom', 'auto']),
    onChange: PropTypes.func,
    options: PropTypes.arrayOf(
      PropTypes.shape({ value: PropTypes.string, label: PropTypes.message }),
    ),
    title: PropTypes.message,
    warning: PropTypes.message,
  }

  FetchSelect.defaultProps = {
    description: defaultDescription,
    menuPlacement: 'auto',
    onChange: () => null,
    options: [],
    title: defaultTitle,
    warning: undefined,
    defaultWarning,
  }

  return FetchSelect
}
