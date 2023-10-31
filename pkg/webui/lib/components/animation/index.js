

import React, { Component } from 'react'
import lottie from 'lottie-web'

import PropTypes from '@ttn-lw/lib/prop-types'

export default class Animation extends Component {
  static propTypes = {
    animationData: PropTypes.shape({}).isRequired,
    lottieConfig: PropTypes.shape({}),
  }

  static defaultProps = {
    lottieConfig: {},
  }

  constructor(props) {
    super(props)

    this.containerRef = React.createRef()
    this.instance = null
  }

  componentDidMount() {
    const { lottieConfig } = this.props

    this.instance = lottie.loadAnimation({
      container: this.containerRef.current,
      renderer: 'svg',
      loop: false,
      autoplay: false,
      animationData: this.props.animationData,
      ...lottieConfig,
    })
  }

  render() {
    const { lottieConfig, animationData, ...rest } = this.props

    return (
      <div
        ref={this.containerRef}
        onMouseEnter={this.handleOnMouseEnter}
        onMouseLeave={this.handleOnMouseLeave}
        {...rest}
      />
    )
  }
}
