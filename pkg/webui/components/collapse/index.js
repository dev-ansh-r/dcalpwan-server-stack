

import React from 'react'
import classnames from 'classnames'
import { defineMessages } from 'react-intl'

import Button from '@ttn-lw/components/button'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './collapse.styl'

const m = defineMessages({
  collapse: 'Collapse',
  expand: 'Expand',
})

const Collapse = props => {
  const { className, id, title, description, initialCollapsed, children, disabled } = props

  const [collapsed, setCollapsed] = React.useState(initialCollapsed)
  const onCollapsedChange = React.useCallback(() => {
    setCollapsed(collapsed => !collapsed)
  }, [])

  const isOpen = !collapsed && !disabled

  const cls = classnames(className, style.section)
  return (
    <section className={cls} data-test-id={id}>
      <div className={style.header}>
        <div className={style.headerContent}>
          <Message className={style.title} component="h3" content={title} />
          <Message className={style.description} component="p" content={description} />
        </div>
        <div className={style.button}>
          <Button
            type="button"
            className={style.expandButton}
            disabled={disabled}
            message={collapsed ? m.expand : m.collapse}
            onClick={onCollapsedChange}
          />
        </div>
      </div>
      {isOpen && <div className={style.content}>{children}</div>}
    </section>
  )
}

Collapse.propTypes = {
  children: PropTypes.node.isRequired,
  className: PropTypes.string,
  description: PropTypes.message.isRequired,
  disabled: PropTypes.bool,
  id: PropTypes.string,
  initialCollapsed: PropTypes.bool,
  title: PropTypes.message.isRequired,
}

Collapse.defaultProps = {
  className: undefined,
  disabled: false,
  id: 'collapsible-section',
  initialCollapsed: true,
}

export default Collapse
