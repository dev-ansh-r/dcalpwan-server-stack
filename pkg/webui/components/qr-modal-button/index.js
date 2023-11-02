

import React, { useCallback } from 'react'
import { defineMessages } from 'react-intl'

import Link from '@ttn-lw/components/link'
import ModalButton from '@ttn-lw/components/button/modal-button'

import Message from '@ttn-lw/lib/components/message'
import ErrorMessage from '@ttn-lw/lib/components/error-message'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import DataSheet from '../data-sheet'
import QR from '../qr'

const QrScanDoc = (
  <Link.Anchor external secondary href="https://www.thethingsindustries.com/docs/">
    Having trouble?
  </Link.Anchor>
)

const m = defineMessages({
  scanEndDeviceContinue: 'Please scan the QR code to continue. {qrScanDoc}',
  invalidData:
    'Invalid QR code data. Please note that only TR005 LoRaWAN® Device Identification QR Code can be scanned. Some devices have unrelated QR codes printed on them that cannot be used.',
  apply: 'Apply',
  scanAgain: 'Scan again',
})

const QRModalButton = props => {
  const { message, onApprove, onCancel, onRead, qrData } = props

  const handleRead = useCallback(
    val => {
      onRead(val)
    },
    [onRead],
  )

  const modalData = (
    <div style={{ width: '100%' }}>
      {qrData.data ? (
        qrData.valid ? (
          <DataSheet data={qrData.data} />
        ) : (
          <ErrorMessage content={m.invalidData} />
        )
      ) : (
        <>
          <QR onChange={handleRead} />
          <Message
            content={m.scanEndDeviceContinue}
            values={{ qrScanDoc: QrScanDoc }}
            component="span"
          />
        </>
      )}
    </div>
  )

  return (
    <ModalButton
      type="button"
      icon="camera_alt"
      onCancel={onCancel}
      onApprove={onApprove}
      message={message}
      modalData={{
        title: sharedMessages.scanEndDevice,
        children: modalData,
        buttonMessage: m.apply,
        approveButtonProps: {
          disabled: !qrData.valid,
        },
        cancelButtonMessage: qrData.data ? m.scanAgain : sharedMessages.cancel,
        cancelButtonProps: qrData.data ? { onClick: onCancel } : {},
        danger: false,
      }}
    />
  )
}

QRModalButton.propTypes = {
  message: PropTypes.message.isRequired,
  onApprove: PropTypes.func.isRequired,
  onCancel: PropTypes.func.isRequired,
  onRead: PropTypes.func.isRequired,
  qrData: PropTypes.shape({
    valid: PropTypes.bool,
    data: PropTypes.arrayOf(PropTypes.shape()),
  }),
}

QRModalButton.defaultProps = {
  qrData: undefined,
}

export default QRModalButton
