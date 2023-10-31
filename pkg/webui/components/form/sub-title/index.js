

import React from 'react'
import classnames from 'classnames'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './sub-title.styl'

const FormSubTitle = React.memo(props => {
  const { className, title } = props

  return (
    <Message
      className={classnames(style.title, className)}
      content={title}
      component="h3"
      firstToUpper
    />
  )
})

FormSubTitle.propTypes = {
  className: PropTypes.string,
  title: PropTypes.message.isRequired,
}

FormSubTitle.defaultProps = {
  className: undefined,
}

export default FormSubTitle
