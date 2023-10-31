

import React, { useCallback } from 'react'
import { defineMessages } from 'react-intl'

import Modal from '@ttn-lw/components/modal'

import PropTypes from '@ttn-lw/lib/prop-types'

const m = defineMessages({
  modalTitle: 'Confirm navigation',
  modalMessage:
    'Are you sure you want to leave this page? Your current changes have not been saved yet.',
})

const MoveAwayModal = ({ blocker }) => {
  const handleComplete = useCallback(
    result => {
      if (result) {
        return blocker.proceed?.()
      }
      return blocker.reset?.()
    },
    [blocker],
  )

  return (
    blocker.state === 'blocked' && (
      <Modal
        buttonMessage={m.modalTitle}
        message={m.modalMessage}
        title={m.modalTitle}
        onComplete={handleComplete}
      />
    )
  )
}

MoveAwayModal.propTypes = {
  blocker: PropTypes.shape({
    proceed: PropTypes.func,
    reset: PropTypes.func,
    state: PropTypes.string,
  }).isRequired,
}

export default MoveAwayModal
