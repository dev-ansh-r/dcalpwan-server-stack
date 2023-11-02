

import React, { useCallback } from 'react'
import PropTypes from 'prop-types'
import { defineMessages } from 'react-intl'

import FileInput from '../../../file-input'
import style from '../../qr.styl'

const m = defineMessages({
  uploadImage: 'Upload a Photo',
  qrCodeNotFound: 'QR code not found',
})

const Capture = props => {
  const { onRead } = props

  const handleChange = useCallback(
    data => {
      const image = new Image()
      image.src = data
      image.onload = () => {
        onRead(image, image.width, image.height)
      }
    },
    [onRead],
  )

  const handleDataTransform = useCallback(content => content, [])

  return (
    <div className={style.captureWrapper}>
      <FileInput
        name="captureFileInput"
        id="captureFileInput"
        onChange={handleChange}
        message={m.uploadImage}
        dataTransform={handleDataTransform}
        providedMessage={m.uploadImage}
        image
        warningSize={0}
        largeFileWarningMessage={m.qrCodeNotFound}
        center
      />
    </div>
  )
}

Capture.propTypes = {
  onRead: PropTypes.func.isRequired,
}

export default Capture
