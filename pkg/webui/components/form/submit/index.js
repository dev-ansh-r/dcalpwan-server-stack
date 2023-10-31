

import React from 'react'

import PropTypes from '@ttn-lw/lib/prop-types'

import { useFormContext } from '..'

const FormSubmit = props => {
  const { component: Component, disabled, ...rest } = props
  const formContext = useFormContext()

  const submitProps = {
    isValid: formContext.isValid,
    isSubmitting: formContext.isSubmitting,
    isValidating: formContext.isValidating,
    submitCount: formContext.submitCount,
    dirty: formContext.dirty,
    validateForm: formContext.validateForm,
    validateField: formContext.validateField,
    disabled: formContext.disabled || disabled,
  }

  return <Component {...rest} {...submitProps} />
}

FormSubmit.propTypes = {
  component: PropTypes.oneOfType([PropTypes.func, PropTypes.string]),
  disabled: PropTypes.bool,
}

FormSubmit.defaultProps = {
  component: 'button',
  disabled: false,
}

export default FormSubmit
