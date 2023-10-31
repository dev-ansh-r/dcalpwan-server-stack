

import { INITIALIZE_SUCCESS } from '@ttn-lw/lib/store/actions/init'

const defaultState = {
  initialized: false,
}

const app = (state = defaultState, action) => {
  switch (action.type) {
    case INITIALIZE_SUCCESS:
      return {
        initialized: true,
      }
    default:
      return state
  }
}

export default app
