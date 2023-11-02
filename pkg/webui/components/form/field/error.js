

import classNames from 'classnames'
import React from 'react'

import Icon from '@ttn-lw/components/icon'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './field.styl'

const FieldError = ({ content, error, warning, title, className, id }) => {
  const icon = error ? 'error' : 'warning'
  const contentValues = content.values || {}
  const classname = classNames(style.message, className, {
    [style.show]: content && content !== '',
    [style.hide]: !content || content === '',
    [style.err]: error,
    [style.warn]: warning,
  })

  if (title) {
    contentValues.field = <Message content={title} className={style.name} key={title.id || title} />
  }

  return (
    <div className={classname} id={id}>
      <Icon icon={icon} className={style.icon} />
      <Message content={content.message || content} values={contentValues} />
    </div>
  )
}

FieldError.propTypes = {
  className: PropTypes.string,
  content: PropTypes.oneOfType([
    PropTypes.error,
    PropTypes.shape({
      message: PropTypes.error.isRequired,
      values: PropTypes.shape({}).isRequired,
    }),
  ]).isRequired,
  error: PropTypes.bool,
  id: PropTypes.string.isRequired,
  title: PropTypes.message,
  warning: PropTypes.bool,
}

FieldError.defaultProps = {
  className: undefined,
  title: undefined,
  warning: false,
  error: false,
}

export default FieldError
