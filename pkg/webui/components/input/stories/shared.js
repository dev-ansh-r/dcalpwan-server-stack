
/* eslint-disable react/prop-types, import/prefer-default-export */

import React from 'react'
import bind from 'autobind-decorator'

import Input from '..'

class Example extends React.Component {
  constructor(props) {
    super(props)

    this.state = {
      value: props.value || '',
    }
  }

  @bind
  onChange(value) {
    this.setState({ value })
  }

  render() {
    const { type, component: Component, ...props } = this.props
    const { value } = this.state

    const InputComponent = Component ? Component : Input

    return <InputComponent {...props} type={type} onChange={this.onChange} value={value} />
  }
}

export { Example }
