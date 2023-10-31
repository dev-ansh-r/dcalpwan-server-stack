

import React from 'react'

import Input from '@ttn-lw/components/input'

import PropTypes from '@ttn-lw/lib/prop-types'

const DevAddrInput = props => {
  const {
    className,
    id,
    name,
    onFocus,
    onChange,
    onBlur,
    value,
    disabled,
    autoFocus,
    warning,
    error: fieldError,
    loading: fieldLoading,
    onGenerate,
    generatedError,
    generatedLoading,
    ...rest
  } = props

  const showLoading = fieldLoading || generatedLoading
  const showError = fieldError
  // Always show field validation error first.
  const showWarning = fieldError ? false : Boolean(warning) && generatedError

  return (
    <Input.Generate
      type="byte"
      id={id}
      min={4}
      max={4}
      className={className}
      name={name}
      onChange={onChange}
      onBlur={onBlur}
      onFocus={onFocus}
      value={value}
      error={showError}
      warning={showWarning}
      loading={showLoading}
      disabled={disabled}
      autoFocus={autoFocus}
      onGenerateValue={onGenerate}
      mayGenerateValue={!fieldLoading && !disabled && !generatedLoading}
      {...rest}
    />
  )
}

DevAddrInput.propTypes = {
  autoFocus: PropTypes.bool,
  className: PropTypes.string,
  disabled: PropTypes.bool,
  error: PropTypes.bool,
  generatedError: PropTypes.bool.isRequired,
  generatedLoading: PropTypes.bool.isRequired,
  id: PropTypes.string.isRequired,
  loading: PropTypes.bool,
  name: PropTypes.string.isRequired,
  onBlur: PropTypes.func.isRequired,
  onChange: PropTypes.func.isRequired,
  onFocus: PropTypes.func,
  onGenerate: PropTypes.func.isRequired,
  value: PropTypes.string,
  warning: PropTypes.bool,
}

DevAddrInput.defaultProps = {
  className: undefined,
  onFocus: () => null,
  disabled: false,
  error: false,
  warning: false,
  autoFocus: false,
  value: undefined,
  loading: false,
}

export default DevAddrInput
