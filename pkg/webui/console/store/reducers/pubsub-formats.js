

import { GET_PUBSUB_FORMATS_SUCCESS } from '@console/store/actions/pubsub-formats'

const defaultState = {
  formats: undefined,
}

const pubsubs = (state = defaultState, { type, payload }) => {
  switch (type) {
    case GET_PUBSUB_FORMATS_SUCCESS:
      return {
        ...state,
        formats: payload,
      }
    default:
      return state
  }
}

export default pubsubs
