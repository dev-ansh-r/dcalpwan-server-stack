

import React from 'react'

import Icon from '@ttn-lw/components/icon'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'

import m from './messages'

import style from './events.styl'

class EventErrorBoundary extends React.Component {
  state = { hasErrored: false, error: undefined, expanded: false }

  static getDerivedStateFromError() {
    return { hasErrored: true }
  }

  render() {
    const { hasErrored } = this.state
    const { children } = this.props

    if (hasErrored) {
      return (
        <div className={style.cellError}>
          <Icon icon="error" className={style.eventIcon} />
          <Message content={m.errorOverviewEntry} />
        </div>
      )
    }

    return children
  }
}

EventErrorBoundary.propTypes = {
  children: PropTypes.node.isRequired,
}

export default EventErrorBoundary
