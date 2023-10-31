

import React from 'react'
import { defineMessages } from 'react-intl'

import Button from '@ttn-lw/components/button'

import PropTypes from '@ttn-lw/lib/prop-types'

import Input from '.'

const m = defineMessages({
  generate: 'Generate',
})

const GenerateInput = props => {
  const { onChange } = props
  const { mayGenerateValue, onGenerateValue, action, ...rest } = props

  const handleGenerateValue = React.useCallback(async () => {
    if (mayGenerateValue) {
      const generatedValue = await onGenerateValue()

      onChange(generatedValue, true)
    }
  }, [mayGenerateValue, onChange, onGenerateValue])

  const generateAction = React.useMemo(
    () => ({
      icon: 'autorenew',
      type: 'button',
      disabled: !mayGenerateValue,
      onClick: handleGenerateValue,
      message: m.generate,
      ...action,
    }),
    [action, handleGenerateValue, mayGenerateValue],
  )

  return <Input {...rest} action={generateAction} />
}

GenerateInput.propTypes = {
  action: PropTypes.shape({
    ...Button.propTypes,
  }),
  mayGenerateValue: PropTypes.bool,
  onChange: PropTypes.func.isRequired,
  onGenerateValue: PropTypes.func.isRequired,
}

GenerateInput.defaultProps = {
  mayGenerateValue: true,
  action: {},
}

export default GenerateInput
