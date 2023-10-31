

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

export const GET_JOIN_EUI_PREFIXES_BASE = 'GET_JOIN_EUI_PREFIXES'
export const [
  {
    request: GET_JOIN_EUI_PREFIXES,
    success: GET_JOIN_EUI_PREFIXES_SUCCESS,
    failure: GET_JOIN_EUI_PREFIXES_FAILURE,
  },
  {
    request: getJoinEUIPrefixes,
    success: getJoinEUIPrefixesSuccess,
    failure: getJoinEUIPrefixesFailure,
  },
] = createRequestActions(GET_JOIN_EUI_PREFIXES_BASE)
