

import { createAction } from 'redux-actions'

export default (baseType, requestPayloadCreator, requestMetaCreator) => {
  const requestType = `${baseType}_REQUEST`
  const successType = `${baseType}_SUCCESS`
  const failureType = `${baseType}_FAILURE`
  const abortType = `${baseType}_ABORT`

  return [
    {
      request: requestType,
      success: successType,
      failure: failureType,
      aborted: abortType,
    },
    {
      request: createAction(requestType, requestPayloadCreator, requestMetaCreator),
      success: createAction(successType),
      failure: createAction(failureType),
      aborted: createAction(abortType),
    },
  ]
}
