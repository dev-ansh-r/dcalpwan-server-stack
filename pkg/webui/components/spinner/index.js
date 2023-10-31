

import React, { useEffect, useState } from 'react'
import classnames from 'classnames'
import PropTypes from 'prop-types'

import from from '@ttn-lw/lib/from'

import style from './spinner.styl'

const id = `grad-${Math.round(Math.random() * 10000)}`

const Spinner = ({
  after,
  center,
  children,
  className,
  faded = false,
  micro = false,
  small,
  inline = false,
}) => {
  const [visible, setVisible] = useState(false)
  const visibilityTimeout = setTimeout(() => setVisible(true), after)

  useEffect(
    () => () => {
      clearTimeout(visibilityTimeout)
    },
    [visibilityTimeout],
  )

  const classname = classnames(
    style.box,
    className,
    ...from(style, {
      center,
      small,
      micro,
      faded,
      visible,
      inline,
    }),
  )

  return (
    <div className={classname}>
      <svg className={style.spinner} viewBox="0 0 100 100">
        <defs>
          <linearGradient id={id}>
            <stop offset="0%" className={style.stop} />
            <stop offset="100%" className={style.stop} stopColor="white" stopOpacity="0" />
          </linearGradient>
        </defs>
        <g transform="translate(50, 50)">
          <circle cx="0" cy="0" r={micro ? 35 : 40} className={style.bar} stroke={`url(#${id})`} />
        </g>
        <circle cx="50" cy="50" r={micro ? 35 : 40} className={style.circle} />
      </svg>
      <div className={style.message}>{children}</div>
    </div>
  )
}

Spinner.propTypes = {
  after: PropTypes.number,
  center: PropTypes.bool,
  children: PropTypes.node,
  className: PropTypes.string,
  faded: PropTypes.bool,
  inline: PropTypes.bool,
  micro: PropTypes.bool,
  small: PropTypes.bool,
}

Spinner.defaultProps = {
  after: 350,
  center: false,
  children: undefined,
  className: undefined,
  faded: false,
  inline: false,
  micro: false,
  small: false,
}

export default Spinner
