

import React from 'react'
import classnames from 'classnames'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './stepper.styl'

const Stepper = props => {
  const { className, children, stepCountStart, currentStep, status, vertical } = props

  const cls = classnames(className, style.stepper, {
    [style.vertical]: vertical,
  })
  const steps = React.Children.map(children, (child, index) => {
    if (!Boolean(child)) {
      return null
    }

    if (React.isValidElement(child) && child.type.displayName === 'Stepper.Step') {
      const stepNumber = index + stepCountStart
      let stepStatus = status
      if (stepNumber < currentStep) {
        stepStatus = 'success'
      } else if (stepNumber > currentStep) {
        stepStatus = 'wait'
      }

      const props = {
        ...child.props,
        stepNumber,
        vertical,
        transitionFailed: status === 'failure' && stepNumber === currentStep - 1,
        active: stepNumber === currentStep,
        status: stepStatus,
      }

      return React.cloneElement(child, props)
    }

    return child
  })

  return <ol className={cls}>{steps}</ol>
}

Stepper.propTypes = {
  children: PropTypes.oneOfType([PropTypes.node, PropTypes.arrayOf(PropTypes.node)]),
  className: PropTypes.string,
  currentStep: PropTypes.number.isRequired,
  status: PropTypes.oneOf(['success', 'failure', 'current', 'wait']),
  stepCountStart: PropTypes.number,
  vertical: PropTypes.bool,
}

Stepper.defaultProps = {
  className: undefined,
  stepCountStart: 1,
  status: 'current',
  children: [],
  vertical: false,
}

export default Stepper
