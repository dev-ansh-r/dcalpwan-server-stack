

import React, { useCallback } from 'react'
import { components } from 'react-select'
import AsyncSelect from 'react-select/async'
import { useIntl } from 'react-intl'
import classnames from 'classnames'
import { debounce } from 'lodash'

import PropTypes from '@ttn-lw/lib/prop-types'

import Icon from '../icon'
import { useFormContext } from '../form'

import style from './select.styl'

const customOption = props => {
  const { showOptionIcon } = props.selectProps

  return (
    <components.Option {...props}>
      {showOptionIcon && <Icon icon={props.data.icon} className="mr-cs-xs" />}
      <b>{props.label}</b>
    </components.Option>
  )
}

customOption.propTypes = {
  data: PropTypes.shape({
    icon: PropTypes.string,
  }).isRequired,
  label: PropTypes.string.isRequired,
}

const Input = props => {
  const { selectProps } = props

  return <components.Input {...props} aria-describedby={selectProps['aria-describedby']} />
}

Input.propTypes = {
  selectProps: PropTypes.shape({
    'aria-describedby': PropTypes.string,
  }).isRequired,
}

const SuggestedSelect = props => {
  const {
    value,
    name,
    onBlur,
    onChange,
    loadOptions,
    className,
    options,
    inputWidth,
    onFocus,
    disabled,
    error,
    warning,
    id,
    placeholder,
    showOptionIcon,
    customComponents,
    ...rest
  } = props

  const { formatMessage } = useIntl()
  const { values } = useFormContext()

  const handleChange = useCallback(
    selectedValue => {
      onChange(selectedValue)
    },
    [onChange],
  )

  const handleBlur = useCallback(
    event => {
      // https://github.com/JedWatson/react-select/issues/3523
      // Make sure the input name is always present in the event object.
      event.target.name = name

      if (typeof values[name] !== 'undefined') {
        // https://github.com/JedWatson/react-select/issues/3175
        event.target.value = values[name]
      }

      onBlur(event)
    },
    [onBlur, name, values],
  )

  const debouncedFetch = debounce((query, callback) => {
    loadOptions(query).then(result => callback(result))
  }, 500)

  const cls = classnames(className, style.container, style[`input-width-${inputWidth}`], {
    [style.error]: error,
    [style.warning]: warning,
  })

  return (
    <AsyncSelect
      {...rest}
      loadOptions={debouncedFetch}
      className={cls}
      inputId={id}
      classNamePrefix="select"
      onChange={handleChange}
      onFocus={onFocus}
      onBlur={handleBlur}
      isDisabled={disabled}
      isClearable
      value={value}
      name={name}
      showOptionIcon={showOptionIcon}
      components={{ Input, Option: customOption, ...customComponents }}
      aria-describedby={rest['aria-describedby']}
      placeholder={Boolean(placeholder) ? formatMessage(placeholder) : undefined}
    />
  )
}

SuggestedSelect.propTypes = {
  className: PropTypes.string,
  customComponents: PropTypes.shape({
    Option: PropTypes.func,
    SingleValue: PropTypes.func,
  }),
  disabled: PropTypes.bool,
  error: PropTypes.bool,
  id: PropTypes.string,
  inputWidth: PropTypes.inputWidth,
  loadOptions: PropTypes.func,
  menuPlacement: PropTypes.string,
  name: PropTypes.string.isRequired,
  onBlur: PropTypes.func,
  onChange: PropTypes.func,
  onFocus: PropTypes.func,
  options: PropTypes.arrayOf(
    PropTypes.shape({
      value: PropTypes.oneOfType([PropTypes.string, PropTypes.number]),
      label: PropTypes.message,
    }),
  ),
  placeholder: PropTypes.message,
  showOptionIcon: PropTypes.bool,
  value: PropTypes.oneOfType([PropTypes.string, PropTypes.shape({})]),
  warning: PropTypes.bool,
}

SuggestedSelect.defaultProps = {
  className: undefined,
  onChange: () => null,
  onBlur: () => null,
  onFocus: () => null,
  options: [],
  disabled: false,
  error: false,
  warning: false,
  value: undefined,
  id: undefined,
  inputWidth: 'm',
  placeholder: undefined,
  menuPlacement: 'auto',
  loadOptions: () => null,
  showOptionIcon: false,
  customComponents: {},
}

export default SuggestedSelect
