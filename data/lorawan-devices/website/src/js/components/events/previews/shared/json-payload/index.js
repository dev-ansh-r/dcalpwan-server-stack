
import React from 'react'
import classnames from 'classnames'

import PropTypes from 'prop-types'

import style from './json-payload.styl'

const OBJECT_PLACEHOLDER = '{…}'
const ARRAY_PLACEHOLDER = '[…]'

const KeyValue = ({ entryKey, entryValue, showKey }) => {
  let renderValue
  if (Array.isArray(entryValue)) {
    renderValue = ARRAY_PLACEHOLDER
  } else if (typeof entryValue === 'object') {
    renderValue = OBJECT_PLACEHOLDER
  } else {
    renderValue = JSON.stringify(entryValue)
  }

  const valueCls = classnames(style.value, {
    [style.string]: typeof entryValue === 'string',
    [style.number]: typeof entryValue === 'number',
    [style.boolean]: typeof entryValue === 'boolean',
    [style.value]: typeof entryValue === 'object',
  })

  const entries = []

  if (showKey) {
    entries.push(
      <span key={entryKey} className={style.key}>
        {entryKey}
      </span>,
    )
  }

  entries.push(
    <span key={`${entryKey}-value`} className={valueCls}>
      {renderValue}
    </span>,
  )

  return entries
}

KeyValue.propTypes = {
  entryKey: PropTypes.string.isRequired,
  entryValue: PropTypes.oneOfType([
    PropTypes.string,
    PropTypes.bool,
    PropTypes.number,
    PropTypes.shape({}),
    PropTypes.array,
  ]),
  showKey: PropTypes.bool,
}

KeyValue.defaultProps = {
  showKey: false,
}

const JSONPayload = React.memo((props) => {
  const { data } = props

  try {
    JSON.stringify(data)
  } catch (e) {
    return <div className={style.invalid}>Invalid JSON</div>
  }

  const dataArray = Array.isArray(data)
  const dataObject = !dataArray && data !== null && typeof data === 'object'

  const cls = classnames(style.jsonPayload, {
    [style.jsonPayloadArray]: dataArray,
    [style.jsonPayloadObject]: dataObject,
  })

  const content = Object.keys(data).map((key) => (
    <KeyValue key={key} entryKey={key} entryValue={data[key]} showKey={dataObject} />
  ))

  return (
    <div title={JSON.stringify(data)} className={cls}>
      {content}
    </div>
  )
})

JSONPayload.propTypes = {
  data: PropTypes.oneOfType([
    PropTypes.string,
    PropTypes.bool,
    PropTypes.number,
    PropTypes.shape({}),
    PropTypes.array,
  ]).isRequired,
}

JSONPayload.defaultProps = {}

export default JSONPayload
