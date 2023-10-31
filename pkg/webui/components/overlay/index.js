

import React from 'react'
import classnames from 'classnames'

import Spinner from '@ttn-lw/components/spinner'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import style from './overlay.styl'

const Overlay = ({
  className,
  overlayClassName,
  spinnerClassName,
  spinnerMessage,
  visible,
  loading,
  children,
}) => (
  <div className={classnames(className, style.overlayWrapper)}>
    <div
      className={classnames(overlayClassName, style.overlay, {
        [style.overlayVisible]: visible,
      })}
    />
    {visible && loading && (
      <Spinner center className={classnames(spinnerClassName, style.overlaySpinner)}>
        <Message content={spinnerMessage} />
      </Spinner>
    )}
    {children}
  </div>
)

Overlay.propTypes = {
  children: PropTypes.node.isRequired,
  className: PropTypes.string,
  /**
   * A flag specifying whether the overlay should display the loading spinner.
   */
  loading: PropTypes.bool,
  overlayClassName: PropTypes.string,
  spinnerClassName: PropTypes.string,
  spinnerMessage: PropTypes.message,
  /** A flag specifying whether the overlay is visible or not. */
  visible: PropTypes.bool.isRequired,
}

Overlay.defaultProps = {
  className: undefined,
  overlayClassName: undefined,
  spinnerClassName: undefined,
  spinnerMessage: sharedMessages.fetching,
  loading: false,
}

export default Overlay
