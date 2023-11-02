

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

export const INITIALIZE_BASE = 'INITIALIZE'
export const [
  { request: INITIALIZE, success: INITIALIZE_SUCCESS, failure: INITIALIZE_FAILURE },
  { request: initialize, success: initializeSuccess, failure: initializeFailure },
] = createRequestActions(INITIALIZE_BASE)
