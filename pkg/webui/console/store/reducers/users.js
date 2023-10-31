

import { getUserId } from '@ttn-lw/lib/selectors/id'

import {
  GET_USERS_LIST_SUCCESS,
  GET_USER_SUCCESS,
  UPDATE_USER_SUCCESS,
  GET_USER,
  DELETE_USER_SUCCESS,
  CREATE_USER_SUCCESS,
  GET_USER_INVITATIONS_SUCCESS,
} from '@console/store/actions/users'

const initialState = {
  entities: {},
  selectedUser: null,
  invitations: [],
  invitationsTotalCount: undefined,
}

const user = (state = {}, user) => ({
  ...state,
  ...user,
})

const users = (state = initialState, { type, payload }) => {
  switch (type) {
    case GET_USER:
      return {
        ...state,
        selectedUser: payload.id,
      }
    case CREATE_USER_SUCCESS:
    case UPDATE_USER_SUCCESS:
    case GET_USER_SUCCESS:
      const id = getUserId(payload)

      return {
        ...state,
        entities: {
          ...state.entities,
          [id]: user(state.entities[id], payload),
        },
      }
    case GET_USERS_LIST_SUCCESS:
      const entities = payload.entities.reduce(
        (acc, app) => {
          const id = getUserId(app)

          acc[id] = user(acc[id], app)
          return acc
        },
        { ...state.entities },
      )

      return {
        ...state,
        entities,
      }
    case DELETE_USER_SUCCESS:
      const { [payload.id]: deleted, ...rest } = state.entities

      return {
        selectedUser: null,
        entities: rest,
      }
    case GET_USER_INVITATIONS_SUCCESS:
      return {
        ...state,
        invitations: payload.invitations,
        invitationsTotalCount: payload.totalCount,
      }
    default:
      return state
  }
}

export default users
