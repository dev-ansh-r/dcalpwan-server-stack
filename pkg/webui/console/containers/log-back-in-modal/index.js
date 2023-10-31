

import React from 'react'
import { defineMessages } from 'react-intl'

import Modal from '@ttn-lw/components/modal'

const m = defineMessages({
  modalTitle: 'Please sign in again',
  modalMessage:
    "You were signed out of the Console. You can press 'Reload' to log back into the Console again.",
  buttonMessage: 'Reload',
})

const reload = () => {
  window.location.reload()
}

const LogBackInModal = () => (
  <Modal
    approval={false}
    buttonMessage={m.buttonMessage}
    message={m.modalMessage}
    title={m.modalTitle}
    onComplete={reload}
    approveButtonProps={{ icon: 'refresh' }}
  />
)

export default LogBackInModal
