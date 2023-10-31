

import React from 'react'

import ProgressBar from '.'

class Helper extends React.Component {
  state = {
    percentage: 0,
  }

  increment() {
    const { percentage } = this.state
    const newVal = percentage > 100 ? 0 : percentage + Math.random() * 15

    this.setState({ percentage: newVal })
  }

  componentDidMount() {
    setInterval(this.increment.bind(this), 1000)
  }

  render() {
    const { percentage } = this.state

    return <ProgressBar percentage={percentage} showStatus />
  }
}

export default {
  title: 'ProgressBar',
}

export const Default = () => <ProgressBar percentage={50} />
export const WithStatus = () => <ProgressBar current={24} target={190} showStatus />
export const WithEtaEstimation = () => <Helper />

WithEtaEstimation.story = {
  name: 'With ETA Estimation',
}
