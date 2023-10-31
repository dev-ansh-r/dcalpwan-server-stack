

import {
  GET_APP_LINK_SUCCESS,
  GET_APP_LINK_FAILURE,
  UPDATE_APP_LINK_SUCCESS,
  DELETE_APP_LINK_SUCCESS,
} from '@console/store/actions/link'

const defaultProps = {
  link: undefined,
}

const getLinkSuccess = (state, { payload }) => {
  const { link = {} } = payload

  return {
    ...state,
    link,
  }
}

const getLinkFailure = (state, { payload }) => ({
  ...state,
  link: payload.link || {},
})

const updateLinkSuccess = (state, link) => {
  const newLink = { ...state.link, ...link.payload }
  return {
    ...state,
    link: newLink,
  }
}

const deleteLinkSuccess = state => ({
  ...state,
  link: {},
})

const link = (state = defaultProps, action) => {
  switch (action.type) {
    case GET_APP_LINK_SUCCESS:
      return getLinkSuccess(state, action)
    case UPDATE_APP_LINK_SUCCESS:
      return updateLinkSuccess(state, action)
    case DELETE_APP_LINK_SUCCESS:
      return deleteLinkSuccess(state)
    case GET_APP_LINK_FAILURE:
      return getLinkFailure(state, action)
    default:
      return state
  }
}

export default link
