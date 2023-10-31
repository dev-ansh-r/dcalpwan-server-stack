

import React from 'react'
import { ToastContainer as Container, cssTransition } from 'react-toastify'

import PropTypes from '@ttn-lw/lib/prop-types'

import createToast from './toast'

import './react-toastify.styl'
import style from './toast.styl'

const ToastContainer = props => (
  <Container toastClassName={style.toast} bodyClassName={style.body} {...props} />
)

ToastContainer.propTypes = {
  autoClose: PropTypes.oneOfType([PropTypes.bool, PropTypes.number]),
  closeButton: PropTypes.oneOfType([PropTypes.bool, PropTypes.element]),
  closeOnClick: PropTypes.bool,
  hideProgressBar: PropTypes.bool,
  limit: PropTypes.number,
  pauseOnFocusLoss: PropTypes.bool,
  pauseOnHover: PropTypes.bool,
  position: PropTypes.oneOf([
    'bottom-right',
    'bottom-left',
    'top-right',
    'top-left',
    'top-center',
    'bottom-center',
  ]),
  transition: PropTypes.func,
}

ToastContainer.defaultProps = {
  autoClose: undefined,
  position: 'bottom-right',
  closeButton: false,
  hideProgressBar: true,
  pauseOnHover: true,
  closeOnClick: true,
  pauseOnFocusLoss: true,
  limit: 2,
  transition: cssTransition({ enter: style.slideInRight, exit: style.slideOutRight }),
}

const toast = createToast()

export { toast as default, ToastContainer }
