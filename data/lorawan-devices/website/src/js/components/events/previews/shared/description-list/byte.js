
import React from 'react'

import SafeInspector from '../../../../safe-inspector'
import PropTypes from 'prop-types'

import { base64ToHex } from '../../../../../lib/bytes'

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
      {isEmpty ? <div className={style.emptyByte}>Empty</div> : <SafeInspector data={data} small />}
    </DescriptionListItem>
  )
}

DescriptionListByteItem.propTypes = {
  convertToHex: PropTypes.bool,
  data: PropTypes.string,
  title: PropTypes.string,
}

DescriptionListByteItem.defaultProps = {
  convertToHex: false,
  data: undefined,
  title: '',
}

export default DescriptionListByteItem
