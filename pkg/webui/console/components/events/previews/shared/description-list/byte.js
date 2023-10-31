

import React from 'react'

import SafeInspector from '@ttn-lw/components/safe-inspector'

import Message from '@ttn-lw/lib/components/message'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'

import { base64ToHex } from '@console/lib/bytes'

import DescriptionListItem from './item'

import style from './description-list.styl'

const DescriptionListByteItem = ({ title, data: rawData, convertToHex }) => {
  if (rawData === undefined) {
    return null
  }

  const isEmpty = rawData.length === 0
  const data = convertToHex ? base64ToHex(rawData) : rawData

  return (
    <DescriptionListItem title={title}>
      {isEmpty ? (
        <Message className={style.emptyByte} firstToLower content={sharedMessages.empty} />
      ) : (
        <SafeInspector
          data={data}
          hideable={false}
          truncateAfter={8}
          initiallyVisible
          small
          disableResize
        />
      )}
    </DescriptionListItem>
  )
}

DescriptionListByteItem.propTypes = {
  convertToHex: PropTypes.bool,
  data: PropTypes.string,
  title: PropTypes.message,
}

DescriptionListByteItem.defaultProps = {
  convertToHex: false,
  data: undefined,
  title: undefined,
}

export default DescriptionListByteItem
