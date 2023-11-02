
import React from 'react'

import Icon from '@ttn-lw/components/icon'

import Tooltip from '.'

const containerStyles = {
  width: '300px',
  margin: '0 auto',
}

const splitterStyles = {
  marginTop: '40px',
}

export default {
  title: 'Tooltip',
}

export const TextContent = () => (
  <div style={containerStyles}>
    <div style={splitterStyles} />
    <Tooltip content="Tooltip with top placement" placement="top">
      <Icon icon="info" />
    </Tooltip>
    <div style={splitterStyles} />
    <Tooltip content="Tooltip with top-start placement" placement="top-start">
      <Icon icon="info" />
    </Tooltip>
    <div style={splitterStyles} />
    <Tooltip content="Tooltip with top-end placement" placement="top-end">
      <Icon icon="info" />
    </Tooltip>
    <div style={splitterStyles} />
    <Tooltip content="Tooltip with bottom placement" placement="bottom">
      <Icon icon="info" />
    </Tooltip>
    <div style={splitterStyles} />
    <Tooltip content="Tooltip with bottom-start placement" placement="bottom-start">
      <Icon icon="info" />
    </Tooltip>
    <div style={splitterStyles} />
    <Tooltip content="Tooltip with bottom-end placement" placement="bottom-end">
      <Icon icon="info" />
    </Tooltip>
    <div style={splitterStyles} />
    <Tooltip content="Tooltip with right placement" placement="right">
      <Icon icon="info" />
    </Tooltip>
    <div style={splitterStyles} />
    <Tooltip content="Tooltip with right-start placement" placement="right-start">
      <Icon icon="info" />
    </Tooltip>
    <div style={splitterStyles} />
    <Tooltip content="Tooltip with right-end placement" placement="right-end">
      <Icon icon="info" />
    </Tooltip>
    <div style={splitterStyles} />
    <Tooltip content="Tooltip with left placement" placement="left">
      <Icon icon="info" />
    </Tooltip>
    <div style={splitterStyles} />
    <Tooltip content="Tooltip with left-start placement" placement="left-start">
      <Icon icon="info" />
    </Tooltip>
    <div style={splitterStyles} />
    <Tooltip content="Tooltip with left-end placement" placement="left-end">
      <Icon icon="info" />
    </Tooltip>
  </div>
)

TextContent.story = {
  name: 'Text content',
}
