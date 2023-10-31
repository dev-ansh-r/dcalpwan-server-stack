

import React, { useCallback } from 'react'
import { useSelector } from 'react-redux'
import { Navigate } from 'react-router-dom'

import toast from '@ttn-lw/components/toast'

import PropTypes from '../../../lib/prop-types'

const Require = ({ children, featureCheck, condition, otherwise }) => {
  const rights = useSelector(state => featureCheck?.rightsSelector(state))
  const combinedCondition = condition || (Boolean(featureCheck) && featureCheck.check(rights))

  const alternativeRender = useCallback(() => {
    if (typeof otherwise === 'object') {
      const { render, redirect, message } = otherwise

      if (message) {
        toast({
          type: toast.types.WARNING,
          message,
        })
      }

      if (typeof redirect === 'string') {
        return <Navigate to={redirect} />
      } else if (typeof render === 'function') {
        return render()
      }
    }

    return null
  }, [otherwise])

  if (!combinedCondition) {
    return alternativeRender()
  }

  return children
}

Require.propTypes = {
  children: PropTypes.node.isRequired,
  otherwise: PropTypes.shape({
    redirect: PropTypes.oneOfType([PropTypes.string, PropTypes.func]),
    render: PropTypes.func,
    message: PropTypes.message,
  }),
}
Require.defaultProps = {
  otherwise: undefined,
  condition: false,
}

export default Require
