

import React from 'react'
import displayName from 'react-display-name'

import PropTypes from '@ttn-lw/lib/prop-types'
import { warn } from '@ttn-lw/lib/log'

export const withEnv = Component => {
  const Base =
    Component.prototype instanceof React.Component ? React.Component : React.PureComponent

  class WithEnv extends Base {
    static displayName = `WithEnv(${displayName(Component)})`

    static contextTypes = {
      env: PropTypes.env,
    }

    render() {
      const { env } = this.context

      if (!env) {
        warn('No env in context, make sure to use env.Provider')
      }

      return <Component env={env || {}} {...this.props} />
    }
  }

  return WithEnv
}

export class EnvProvider extends React.PureComponent {
  static propTypes = {
    children: PropTypes.node.isRequired,
    env: PropTypes.env.isRequired,
  }

  static childContextTypes = {
    env: PropTypes.env.isRequired,
  }

  getChildContext() {
    return {
      env: this.props.env,
    }
  }

  render() {
    return this.props.children
  }
}

export default withEnv
