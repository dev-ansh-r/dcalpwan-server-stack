

import { GET_WEBHOOK_FORMATS_SUCCESS } from '@console/store/actions/webhook-formats'

const defaultState = {
  formats: undefined,
}

const webhooks = (state = defaultState, { type, payload }) => {
  switch (type) {
    case GET_WEBHOOK_FORMATS_SUCCESS:
      return {
        ...state,
        formats: payload,
      }
    default:
      return state
  }
}

export default webhooks
