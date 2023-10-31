

import React from 'react'
import bind from 'autobind-decorator'

import Status from '.'

const containerStyle = {
  width: '120px',
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'space-between',
}

class Toggle extends React.Component {
  state = {
    status: 'unknown',
  }

  @bind
  toggleStatus() {
    const { status } = this.state
    let nextStatus
    switch (status) {
      case 'unknown':
        nextStatus = 'bad'
        break
      case 'bad':
        nextStatus = 'mediocre'
        break
      case 'mediocre':
        nextStatus = 'good'
        break
      case 'good':
        nextStatus = 'unknown'
        break
    }

    this.setState({ status: nextStatus })
  }

  render() {
    const { status } = this.state
    return (
      <div>
        <Status status={status} />
        <br />
        <button onClick={this.toggleStatus}>Toggle</button>
      </div>
    )
  }
}

export default {
  title: 'Status',
}

export const AllTypes = () => (
  <div>
    <div style={containerStyle}>
      <span>Good:</span>
      <Status status="good" />
    </div>
    <div style={containerStyle}>
      <span>Bad:</span>
      <Status status="bad" />
    </div>
    <div style={containerStyle}>
      <span>Mediocre:</span>
      <Status status="mediocre" />
    </div>
    <div style={containerStyle}>
      <span>Unknown:</span>
      <Status status="unknown" />
    </div>
  </div>
)

AllTypes.story = {
  name: 'All types',
}

export const WithLabel = () => (
  <div>
    <div style={containerStyle}>
      <Status label="Network Status" status="good" />
    </div>
    <div style={containerStyle}>
      <Status label="Network Status" status="bad" />
    </div>
    <div style={containerStyle}>
      <Status label="Network Status" status="mediocre" />
    </div>
    <div style={containerStyle}>
      <Status label="Network Status" status="unknown" />
    </div>
  </div>
)

WithLabel.story = {
  name: 'With label',
}

export const WithoutPulse = () => (
  <div>
    <div style={containerStyle}>
      <Status label="No Pulse" status="good" pulse={false} />
    </div>
    <div style={containerStyle}>
      <Status label="No Pulse" status="bad" pulse={false} />
    </div>
    <div style={containerStyle}>
      <Status label="No Pulse" status="mediocre" pulse={false} />
    </div>
    <div style={containerStyle}>
      <Status label="No Pulse" status="unknown" pulse={false} />
    </div>
  </div>
)

export const _Toggle = () => <Toggle />
