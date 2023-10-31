

import { CancelablePromise } from 'cancelable-promise'

/**
 * `promisifyDispatch` is a decorator for the dispatch function that attaches
 * a cancelable promise to the action that it will use as return value.
 * You should usually use the middleware instead.
 *
 * @param {object} dispatch - The to be decorated dispatch function.
 * @returns {Function} - The decorated dispatch function.
 */
export const promisifyDispatch = dispatch => action => {
  if (action.meta && action.meta._attachPromise && !action.meta._resolve && !action.meta._reject) {
    return new CancelablePromise((resolve, reject) => {
      action.meta = {
        ...action.meta,
        _resolve: resolve,
        _reject: reject,
      }
      dispatch(action)
    })
  }
  return dispatch(action)
}

/**
 * This middleware will check for request actions and attach a cancelable
 * promise to the action.
 *
 * @returns {object} The middleware.
 */
const requestPromiseMiddleware = () => promisifyDispatch

export default requestPromiseMiddleware
