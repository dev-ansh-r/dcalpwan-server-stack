

import { isUndefined, omitBy } from 'lodash'

import {
  GET_APP_PKG_DEFAULT_ASSOC_SUCCESS,
  SET_APP_PKG_DEFAULT_ASSOC_SUCCESS,
  DELETE_APP_PKG_DEFAULT_ASSOC_SUCCESS,
} from '@console/store/actions/application-packages'

const defaultState = {
  default: {},
}

const applicationPackages = (state = defaultState, { type, payload }) => {
  switch (type) {
    case GET_APP_PKG_DEFAULT_ASSOC_SUCCESS:
    case SET_APP_PKG_DEFAULT_ASSOC_SUCCESS:
      return {
        ...state,
        default: omitBy(
          {
            ...state.default,
            [payload.ids.f_port]: 'created_at' in payload ? payload : undefined,
          },
          isUndefined,
        ),
      }
    case DELETE_APP_PKG_DEFAULT_ASSOC_SUCCESS:
      const { [payload.fPort]: deleted, ...rest } = state.default

      return {
        ...state,
        default: rest,
      }
    default:
      return state
  }
}

export default applicationPackages
