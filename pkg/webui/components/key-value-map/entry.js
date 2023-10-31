

import React, { useCallback, useMemo, useState } from 'react'
import { defineMessages } from 'react-intl'
import classnames from 'classnames'

import Button from '@ttn-lw/components/button'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'

import style from './key-value-map.styl'

const m = defineMessages({
  deleteEntry: 'Delete entry',
})

const Entry = ({
  readOnly,
  name,
  value,
  fieldValue,
  index,
  onRemoveButtonClick,
  onChange,
  onBlur,
  inputElement: InputElement,
  indexAsKey,
  valuePlaceholder,
  keyPlaceholder,
  additionalInputProps,
  removeMessage,
  distinctOptions,
  atLeastOneEntry,
}) => {
  const [currentValue, setCurrentValue] = useState(value)
  const [newOptions, setNewOptions] = useState(undefined)
  const { options, ...additionalInputPropsRest } = additionalInputProps
  const _getKeyInputName = useMemo(() => `${name}[${index}].key`, [index, name])

  const _getValueInputName = useMemo(() => `${name}[${index}].value`, [index, name])

  const handleRemoveButtonClick = useCallback(
    event => {
      onRemoveButtonClick(index, event)
    },
    [index, onRemoveButtonClick],
  )

  const handleKeyChanged = useCallback(
    newKey => {
      onChange(index, { key: newKey })
    },
    [index, onChange],
  )

  const handleValueChanged = useCallback(
    newValue => {
      setCurrentValue(newValue)
      onChange(index, { value: newValue })
    },
    [index, onChange],
  )

  const handleBlur = useCallback(
    event => {
      const { relatedTarget } = event
      const nextTarget = relatedTarget || {}

      if (nextTarget.name !== _getKeyInputName && nextTarget.name !== _getValueInputName) {
        onBlur({
          target: {
            name,
            value,
          },
        })
      }
    },
    [onBlur, name, value, _getKeyInputName, _getValueInputName],
  )

  const handleOptionComposition = useCallback(() => {
    let newOptions
    if (currentValue) {
      newOptions = options.filter(v => !fieldValue.includes(v.value) || v.value === currentValue)
    } else {
      newOptions = options.filter(v => !fieldValue.includes(v.value))
    }
    setNewOptions(newOptions)
  }, [currentValue, options, fieldValue])

  const showRemoveButton = atLeastOneEntry ? index !== 0 : true

  return (
    <div className={style.entriesRow}>
      {!indexAsKey && (
        <InputElement
          data-test-id={_getKeyInputName}
          className={style.input}
          name={_getKeyInputName}
          placeholder={keyPlaceholder}
          type="text"
          onChange={handleKeyChanged}
          onBlur={handleBlur}
          value={value.key}
          readOnly={readOnly}
          code
          {...additionalInputProps}
        />
      )}
      <InputElement
        data-test-id={_getValueInputName}
        className={classnames(style.input, { [style.inputIndexAsKey]: indexAsKey })}
        name={_getValueInputName}
        placeholder={valuePlaceholder}
        type="text"
        onChange={handleValueChanged}
        onBlur={handleBlur}
        onFocus={distinctOptions && options ? handleOptionComposition : undefined}
        value={indexAsKey ? value : value.value}
        readOnly={readOnly}
        code
        options={options ? newOptions ?? options : undefined}
        {...additionalInputPropsRest}
      />
      {showRemoveButton && (
        <Button
          type="button"
          onClick={handleRemoveButtonClick}
          icon="delete"
          title={m.deleteEntry}
          message={removeMessage}
          disabled={readOnly}
          danger={!Boolean(removeMessage)}
        />
      )}
    </div>
  )
}

Entry.propTypes = {
  additionalInputProps: PropTypes.shape({
    options: PropTypes.array,
  }).isRequired,
  atLeastOneEntry: PropTypes.bool,
  distinctOptions: PropTypes.bool,
  fieldValue: PropTypes.any,
  index: PropTypes.number.isRequired,
  indexAsKey: PropTypes.bool.isRequired,
  inputElement: PropTypes.elementType.isRequired,
  keyPlaceholder: PropTypes.message.isRequired,
  name: PropTypes.string.isRequired,
  onBlur: PropTypes.func.isRequired,
  onChange: PropTypes.func.isRequired,
  onRemoveButtonClick: PropTypes.func.isRequired,
  readOnly: PropTypes.bool,
  removeMessage: PropTypes.message,
  value: PropTypes.oneOfType([
    PropTypes.shape({
      key: PropTypes.oneOfType([PropTypes.string, PropTypes.number]),
      value: PropTypes.oneOfType([PropTypes.string, PropTypes.number]),
    }),
    PropTypes.string,
  ]),
  valuePlaceholder: PropTypes.message.isRequired,
}

Entry.defaultProps = {
  value: undefined,
  readOnly: false,
  removeMessage: sharedMessages.remove,
  distinctOptions: false,
  fieldValue: undefined,
  atLeastOneEntry: false,
}

export default Entry
