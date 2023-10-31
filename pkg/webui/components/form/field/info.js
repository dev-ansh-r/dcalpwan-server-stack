

import React from 'react'
import classnames from 'classnames'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'
import from from '@ttn-lw/lib/from'

import { useFormContext } from '..'

import Tooltip from './tooltip'

import style from './field.styl'

const InfoField = props => {
  const { children, className, title, disabled: fieldDisabled, tooltipId } = props
  const { disabled: formDisabled } = useFormContext()
  const disabled = formDisabled || fieldDisabled
  const cls = classnames(className, style.field, from(style, { disabled }))

  return (
    <div className={cls}>
      {title && (
        <div className={style.label}>
          <Message content={title} className={style.title} />
          {tooltipId && <Tooltip id={tooltipId} glossaryTerm={title} />}
        </div>
      )}
      <div className={classnames(style.componentArea, style.infoArea)}>{children}</div>
    </div>
  )
}

InfoField.propTypes = {
  children: PropTypes.node.isRequired,
  className: PropTypes.string,
  disabled: PropTypes.bool,
  title: PropTypes.message,
  tooltipId: PropTypes.string,
}

InfoField.defaultProps = {
  className: undefined,
  title: undefined,
  disabled: false,
  tooltipId: undefined,
}

export default InfoField
