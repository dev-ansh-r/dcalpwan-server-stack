

import React from 'react'
import { unstable_usePrompt } from 'react-router-dom'

import PortalledModal from '@ttn-lw/components/modal/portalled'

import PropTypes from '@ttn-lw/lib/prop-types'

const Prompt = props => {
  const { modal, children, message, when } = props
  const [showModal, setShowModal] = React.useState(false)

  // The usage of `unstable_usePrompt` might change as the library updates.
  const continueNavigation = unstable_usePrompt(when, message)

  const handleModalComplete = React.useCallback(
    approved => {
      setShowModal(false)
      if (approved) {
        continueNavigation()
      }
    },
    [continueNavigation],
  )

  React.useEffect(() => {
    if (when) {
      setShowModal(true)
    }
  }, [when])

  return (
    <PortalledModal visible={showModal} {...modal} approval onComplete={handleModalComplete}>
      {children}
    </PortalledModal>
  )
}

Prompt.propTypes = {
  children: PropTypes.oneOfType([PropTypes.arrayOf(PropTypes.node), PropTypes.node]),
  message: PropTypes.string.isRequired,
  modal: PropTypes.shape({ ...PortalledModal.Modal.propTypes }).isRequired,
  when: PropTypes.bool.isRequired,
}

Prompt.defaultProps = {
  children: undefined,
}

export default Prompt
