

import {
  GET_AUTHORIZATIONS_LIST_SUCCESS,
  GET_ACCESS_TOKENS_LIST_SUCCESS,
  DELETE_ALL_TOKENS_SUCCESS,
  DELETE_ACCESS_TOKEN_SUCCESS,
} from '@account/store/actions/authorizations'

const defaultState = {
  authorizations: [],
  authorizationsTotalCount: undefined,
  tokens: [],
  tokensTotalCount: undefined,
}

const authorizations = (state = defaultState, { type, payload }) => {
  switch (type) {
    case GET_AUTHORIZATIONS_LIST_SUCCESS:
      return {
        ...state,
        authorizations: payload.entities,
        authorizationsTotalCount: payload.authorizationsTotalCount,
      }
    case GET_ACCESS_TOKENS_LIST_SUCCESS:
      return {
        ...state,
        tokens: payload.entities,
        tokensTotalCount: payload.tokensTotalCount,
      }
    case DELETE_ALL_TOKENS_SUCCESS:
      return {
        ...state,
        tokens: [],
        tokensTotalCount: 0,
      }
    case DELETE_ACCESS_TOKEN_SUCCESS:
      return {
        ...state,
        tokens: state.tokens.filter(token => token.id !== payload.id),
        tokensTotalCount: state.tokensTotalCount - 1,
      }
    default:
      return state
  }
}

export default authorizations
