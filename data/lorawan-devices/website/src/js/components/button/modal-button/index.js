
import React from 'react'
import PropTypes from 'prop-types'

import PortalledModal from '../../modal/portalled'
import Button from '..'

const ModalButton = function (props) {
  const [modalVisible, setModalVisible] = React.useState(false)
  const {
    modalData,
    onApprove,
    text,
    cancelButtonMessage,
    approveButtonMessage,
    helpMessage,
    ...rest
  } = props

  function handleClick() {
    setModalVisible(true)
  }

  function handleComplete(confirmed) {
    if (confirmed) {
      onApprove()
    }
    setModalVisible(false)
  }

  const modalComposedData = {
    approval: true,
    danger: true,
    buttonMessage: text,
    title: text,
    onComplete: handleComplete,
    cancelButtonMessage,
    approveButtonMessage,
    helpMessage,
    ...modalData,
  }

  return (
    <React.Fragment>
      <PortalledModal visible={modalVisible} modal={modalComposedData} />
      <Button onClick={handleClick} text={text} {...rest} />
    </React.Fragment>
  )
}

ModalButton.propTypes = {
  approveButtonMessage: PropTypes.string,
  cancelButtonMessage: PropTypes.string,
  helpMessage: PropTypes.string,
  modalData: PropTypes.shape(),
  onApprove: PropTypes.func,
  text: PropTypes.string,
}

ModalButton.defaultProps = {
  approveButtonMessage: 'Approve',
  cancelButtonMessage: 'Cancel',
  helpMessage: '',
  modalData: {},
  onApprove: () => {},
  text: '',
}

export default ModalButton
