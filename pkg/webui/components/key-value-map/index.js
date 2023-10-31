

import React, { useCallback } from 'react'
import { defineMessages } from 'react-intl'
import classnames from 'classnames'

import Button from '@ttn-lw/components/button'
import Input from '@ttn-lw/components/input'

import PropTypes from '@ttn-lw/lib/prop-types'

import Entry from './entry'

import style from './key-value-map.styl'

const m = defineMessages({
  addEntry: 'Add entry',
})

const KeyValueMap = ({
  addMessage,
  removeMessage,
  additionalInputProps,
  className,
  disabled,
  indexAsKey,
  inputElement,
  isReadOnly,
  keyPlaceholder,
  name,
  onBlur,
  onChange,
  value,
  valuePlaceholder,
  distinctOptions,
  atLeastOneEntry,
}) => {
  const handleEntryChange = useCallback(
    (index, newValues) => {
      onChange(
        value.map((val, idx) => {
          if (index !== idx) {
            return val
          }

          return indexAsKey ? newValues.value : { ...val, ...newValues }
        }),
      )
    },
    [indexAsKey, onChange, value],
  )

  const removeEntry = useCallback(
    index => {
      onChange(value.filter((_, i) => i !== index) || [], true)
    },
    [onChange, value],
  )

  const addEmptyEntry = useCallback(() => {
    const entry = indexAsKey ? '' : { key: '', value: '' }

    onChange([...value, entry])
  }, [indexAsKey, onChange, value])

  return (
    <div data-test-id={'key-value-map'} className={classnames(className, style.container)}>
      <div>
        {value &&
          value.map((individualValue, index) => (
            <Entry
              key={`${name}[${index}]`}
              name={name}
              value={individualValue}
              fieldValue={value}
              keyPlaceholder={keyPlaceholder}
              valuePlaceholder={valuePlaceholder}
              index={index}
              onRemoveButtonClick={removeEntry}
              onChange={handleEntryChange}
              onBlur={onBlur}
              indexAsKey={indexAsKey}
              readOnly={isReadOnly(individualValue)}
              inputElement={inputElement}
              additionalInputProps={additionalInputProps}
              removeMessage={removeMessage}
              distinctOptions={distinctOptions}
              atLeastOneEntry={atLeastOneEntry}
            />
          ))}
      </div>
      <div>
        <Button
          name={`${name}.push`}
          type="button"
          message={addMessage}
          onClick={addEmptyEntry}
          disabled={disabled}
          icon="add"
        />
      </div>
    </div>
  )
}

KeyValueMap.propTypes = {
  addMessage: PropTypes.message,
  additionalInputProps: PropTypes.shape({}),
  atLeastOneEntry: PropTypes.bool,
  className: PropTypes.string,
  disabled: PropTypes.bool,
  distinctOptions: PropTypes.bool,
  indexAsKey: PropTypes.bool,
  inputElement: PropTypes.elementType,
  isReadOnly: PropTypes.func,
  keyPlaceholder: PropTypes.message,
  name: PropTypes.string.isRequired,
  onBlur: PropTypes.func,
  onChange: PropTypes.func,
  removeMessage: PropTypes.message,
  value: PropTypes.arrayOf(
    PropTypes.oneOfType([
      PropTypes.shape({
        key: PropTypes.oneOfType([PropTypes.string, PropTypes.number]).isRequired,
        value: PropTypes.oneOfType([PropTypes.string, PropTypes.number]),
      }),
      PropTypes.string,
    ]),
  ),
  valuePlaceholder: PropTypes.message.isRequired,
}

KeyValueMap.defaultProps = {
  additionalInputProps: {},
  className: undefined,
  onBlur: () => null,
  onChange: () => null,
  value: [],
  addMessage: m.addEntry,
  indexAsKey: false,
  keyPlaceholder: '',
  disabled: false,
  isReadOnly: () => null,
  inputElement: Input,
  removeMessage: undefined,
  distinctOptions: false,
  atLeastOneEntry: false,
}

export default KeyValueMap
