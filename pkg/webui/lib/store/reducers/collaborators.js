

import { getCollaboratorId } from '@ttn-lw/lib/selectors/id'
import {
  GET_COLLABORATORS_LIST_SUCCESS,
  GET_COLLABORATOR_SUCCESS,
  GET_COLLABORATOR,
} from '@ttn-lw/lib/store/actions/collaborators'

const defaultState = {
  entities: {},
  selectedCollaborator: null,
  totalCount: null,
}

const collaborator = (state = {}, collaborator) => ({
  ...state,
  ...collaborator,
})

const collaborators = (state = defaultState, { type, payload }) => {
  switch (type) {
    case GET_COLLABORATOR:
      return {
        ...state,
        selectedCollaborator: payload.collaboratorId,
      }
    case GET_COLLABORATOR_SUCCESS:
      const id = getCollaboratorId(payload)
      return {
        ...state,
        entities: {
          ...state.entities,
          [id]: collaborator(state.entities[id], payload),
        },
      }
    case GET_COLLABORATORS_LIST_SUCCESS:
      return {
        ...state,
        entities: {
          ...payload.entities.reduce(
            (acc, col) => {
              const id = getCollaboratorId(col)
              acc[id] = collaborator(state.entities[id], col)
              return acc
            },
            { ...state.entities },
          ),
        },
        totalCount: payload.totalCount,
      }
    default:
      return state
  }
}

export default collaborators
