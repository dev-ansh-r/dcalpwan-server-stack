

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

export const GET_APP_LINK_BASE = 'GET_APPLICATION_LINK'
export const [
  { request: GET_APP_LINK, success: GET_APP_LINK_SUCCESS, failure: GET_APP_LINK_FAILURE },
  {
    request: getApplicationLink,
    success: getApplicationLinkSuccess,
    failure: getApplicationLinkFailure,
  },
] = createRequestActions(
  GET_APP_LINK_BASE,
  id => ({ id }),
  (id, selector) => ({ selector }),
)

export const UPDATE_APP_LINK_BASE = 'UPDATE_APPLICATION_LINK'
export const [
  { request: UPDATE_APP_LINK, success: UPDATE_APP_LINK_SUCCESS, failure: UPDATE_APP_LINK_FAILURE },
  {
    request: updateApplicationLink,
    success: updateApplicationLinkSuccess,
    failure: updateApplicationLinkFailure,
  },
] = createRequestActions(UPDATE_APP_LINK_BASE, (id, link) => ({ id, link }))

export const DELETE_APP_LINK_SUCCESS = 'DELETE_APPLICATION_LINK_SUCCESS'

export const deleteApplicationLinkSuccess = () => ({ type: DELETE_APP_LINK_SUCCESS })
