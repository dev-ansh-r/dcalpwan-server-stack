

import React from 'react'

import Spinner from '@ttn-lw/components/spinner'

import Message from '@ttn-lw/lib/components/message'

import useRequest from '@ttn-lw/lib/hooks/use-request'
import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import { FullViewErrorInner } from './full-view-error'

// `<RequireRequest />` is a utility component that can wrap a component tree
// and dispatch a request action, rendering a loading spinner until the request
// has been resolved. It also takes care of rendering possible errors if wished.
const RequireRequest = ({
  requestAction,
  children,
  handleErrors,
  spinnerProps,
  errorRenderFunction: ErrorRenderFunction,
}) => {
  const [fetching, error] = useRequest(requestAction)
  if (fetching) {
    return (
      <Spinner {...spinnerProps}>
        <Message content={sharedMessages.fetching} />
      </Spinner>
    )
  }

  if (error && handleErrors) {
    return <ErrorRenderFunction error={error} />
  }

  return children
}

RequireRequest.propTypes = {
  children: PropTypes.node.isRequired,
  errorRenderFunction: PropTypes.func,
  handleErrors: PropTypes.bool,
  requestAction: PropTypes.oneOfType([
    PropTypes.shape({}),
    PropTypes.arrayOf(PropTypes.shape({})),
    PropTypes.func,
  ]).isRequired,
  spinnerProps: PropTypes.shape(Spinner.propTypes),
}

RequireRequest.defaultProps = {
  errorRenderFunction: FullViewErrorInner,
  handleErrors: true,
  spinnerProps: { center: true },
}

export default RequireRequest
