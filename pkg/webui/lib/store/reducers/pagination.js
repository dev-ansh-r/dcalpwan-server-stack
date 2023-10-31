

import {
  createPaginationByIdRequestActions,
  createPaginationRequestActions,
  createPaginationDeleteActions,
  createPaginationByIdDeleteActions,
  createPaginationRestoreActions,
} from '@ttn-lw/lib/store/actions/pagination'

const defaultState = {
  ids: [],
  totalCount: undefined,
}

export const createNamedPaginationReducer = (reducerName, entityIdSelector) => {
  const [{ success: GET_PAGINATION_SUCCESS }] = createPaginationRequestActions(reducerName)
  const [{ success: DELETE_PAGINATION_SUCCESS }] = createPaginationDeleteActions(reducerName)
  const [{ success: RESTORE_PAGINATION_SUCCESS }] = createPaginationRestoreActions(reducerName)

  return (state = defaultState, { type, payload }) => {
    switch (type) {
      case GET_PAGINATION_SUCCESS:
        return {
          ...state,
          totalCount: payload.totalCount,
          ids: payload.entities.map(entityIdSelector),
        }
      case DELETE_PAGINATION_SUCCESS:
      case RESTORE_PAGINATION_SUCCESS:
        return {
          ...state,
          totalCount: state.totalCount - 1,
          ids: state.ids.filter(id => id !== payload.id),
        }
      default:
        return state
    }
  }
}

export const createNamedPaginationReducerById = (reducerName, entityIdSelector) => {
  const [{ success: GET_PAGINATION_SUCCESS }] = createPaginationByIdRequestActions(reducerName)
  const [{ success: DELETE_PAGINATION_SUCCESS }] = createPaginationByIdDeleteActions(reducerName)
  const [, { success }] = createPaginationDeleteActions(reducerName)
  const paginationReducer = createNamedPaginationReducer(reducerName, entityIdSelector)

  return (state = {}, action) => {
    const { id } = action.payload

    if (!id) {
      return state
    }

    switch (action.type) {
      case GET_PAGINATION_SUCCESS:
        return {
          ...state,
          [id]: paginationReducer(state[id], action),
        }
      case DELETE_PAGINATION_SUCCESS:
        return {
          ...state,
          [id]: paginationReducer(state[id], success({ id: action.payload.targetId })),
        }
      default:
        return state
    }
  }
}
