

import { createAction } from 'redux-actions'

export const createRequestActions = (baseType, requestPayloadCreator, requestMetaCreator) => {
  const requestType = `${baseType}_REQUEST`
  const successType = `${baseType}_SUCCESS`
  const failureType = `${baseType}_FAILURE`
  const abortType = `${baseType}_ABORT`

  return [
    {
      request: requestType,
      success: successType,
      failure: failureType,
      abort: abortType,
    },
    {
      request: createAction(requestType, requestPayloadCreator, requestMetaCreator),
      success: createAction(successType),
      failure: createAction(failureType),
      abort: createAction(abortType),
    },
  ]
}

/**
 * The attachPromise function extends an action creator or action to include a
 * flag which results in a promise being attached to the action by the promise
 * middleware.
 *
 * @param {object|Function} actionOrActionCreator - The original action or
 * action creator.
 * @returns {object|Function} - The modified action or action creator.
 */
export default actionOrActionCreator => {
  const decorateAction = action => ({
    ...action,
    meta: {
      ...action.meta,
      _attachPromise: true,
    },
  })

  if (typeof actionOrActionCreator === 'object') {
    return decorateAction(actionOrActionCreator)
  }

  return (...args) => {
    const action = actionOrActionCreator(...args)
    return decorateAction(action)
  }
}

/**
 * Helper function to retrieve the result action types based
 * on the request action type.
 *
 * @param {string} typeString - The request action type.
 * @param {string} status - The result type, either `SUCCESS`, `FAILURE` or `ABORT`.
 * @returns {string} - The result action type.
 */
export const getResultActionFromType = (typeString, status) => typeString.replace('REQUEST', status)
