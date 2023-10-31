

import React, { useEffect } from 'react'
import { useNavigate, useParams } from 'react-router-dom'

import PropTypes from '../prop-types'

import GenericNotFound from './full-view-error/not-found'

const render = ({ Component, element }) => {
  if (Component) {
    return <Component />
  } else if (element) {
    return element
  }

  return null
}

const checkParams = (check, params) =>
  Object.keys(check).every(paramKey => {
    const singleCheck = check[paramKey]
    if (singleCheck instanceof RegExp) {
      return singleCheck.test(params[paramKey])
    } else if (typeof singleCheck === 'function') {
      return singleCheck(params)
    }
    return false
  })

const ValidateRouteParam = ({ check, Component, element, otherwise }) => {
  const params = useParams()
  const navigate = useNavigate()

  const shouldRender = checkParams(check, params)

  useEffect(() => {
    if (!shouldRender && otherwise.redirect) {
      navigate(otherwise.redirect)
    }
  }, [shouldRender, navigate, otherwise.redirect])

  if (shouldRender) {
    return render({ Component, element })
  }

  if (otherwise.render) {
    return otherwise.render()
  }
}

ValidateRouteParam.propTypes = {
  Component: PropTypes.func,
  check: PropTypes.objectOf(PropTypes.oneOfType([PropTypes.func, PropTypes.instanceOf(RegExp)]))
    .isRequired,
  element: PropTypes.node,
  otherwise: PropTypes.shape({
    render: PropTypes.func,
    redirect: PropTypes.string,
  }),
}

ValidateRouteParam.defaultProps = {
  Component: undefined,
  element: undefined,
  otherwise: {
    render: () => <GenericNotFound />,
  },
}

export default ValidateRouteParam
