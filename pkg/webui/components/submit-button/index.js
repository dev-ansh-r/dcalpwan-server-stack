

import React from 'react'

import Button from '@ttn-lw/components/button'

import PropTypes from '@ttn-lw/lib/prop-types'

const SubmitButton = ({ disabled, icon, isSubmitting, isValidating, message, ...rest }) => {
  const buttonLoading = isSubmitting || isValidating

  return (
    <Button
      primary
      {...rest}
      type="submit"
      icon={icon}
      message={message}
      disabled={disabled}
      busy={buttonLoading}
    />
  )
}

SubmitButton.propTypes = {
  disabled: PropTypes.bool,
  icon: PropTypes.string,
  isSubmitting: PropTypes.bool.isRequired,
  isValidating: PropTypes.bool.isRequired,
  message: PropTypes.message,
}

SubmitButton.defaultProps = {
  disabled: false,
  icon: undefined,
  message: undefined,
}

export default SubmitButton
