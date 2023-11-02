
import React from 'react'
import ReactSwitch from 'react-switch'

import COLORS from '@ttn-lw/constants/colors'

import Icon from '@ttn-lw/components/icon'

const Switch = props => (
  <ReactSwitch
    {...props}
    uncheckedIcon={
      <Icon icon="close" style={{ color: 'white', marginLeft: '3px', fontSize: '1rem' }} />
    }
    checkedIcon={
      <Icon icon="check" style={{ color: 'white', marginLeft: '5px', fontSize: '1rem' }} />
    }
    onColor={COLORS.C_ACTIVE_BLUE}
    activeBoxShadow={`"0 0 3px 5px ${COLORS.C_ACTIVE_BLUE}66, inset 0 0 3px 1px #0002"`}
    height={24}
    width={44}
    data-test-id="switch"
  />
)

Switch.propTypes = ReactSwitch.propTypes
Switch.defaultProps = ReactSwitch.defaultProps

export default Switch
