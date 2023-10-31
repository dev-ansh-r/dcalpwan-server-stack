

import { getOrganizationId } from '@ttn-lw/lib/selectors/id'

import {
  GET_ORGS_LIST_SUCCESS,
  CREATE_ORG_SUCCESS,
  GET_ORG,
  GET_ORG_SUCCESS,
  UPDATE_ORG_SUCCESS,
  DELETE_ORG_SUCCESS,
  GET_ORG_COLLABORATOR_COUNT_SUCCESS,
} from '@console/store/actions/organizations'

const organization = (state = {}, organization) => ({
  ...state,
  ...organization,
})

const defaultState = {
  entities: {},
  selectedOrganization: null,
  collaboratorCounts: {},
}

const organizations = (state = defaultState, { type, payload }) => {
  switch (type) {
    case GET_ORG:
      return {
        ...state,
        selectedOrganization: payload.id,
      }
    case GET_ORGS_LIST_SUCCESS:
      const entities = payload.entities.reduce(
        (acc, org) => {
          const id = getOrganizationId(org)

          acc[id] = organization(acc[id], org)
          return acc
        },
        { ...state.entities },
      )

      return {
        ...state,
        entities,
      }
    case CREATE_ORG_SUCCESS:
    case GET_ORG_SUCCESS:
    case UPDATE_ORG_SUCCESS:
      const id = getOrganizationId(payload)

      return {
        ...state,
        entities: {
          ...state.entities,
          [id]: organization(state.entities[id], payload),
        },
      }
    case GET_ORG_COLLABORATOR_COUNT_SUCCESS:
      return {
        ...state,
        collaboratorCounts: {
          ...state.collaboratorCounts,
          [payload.id]: payload.collaboratorCount,
        },
      }
    case DELETE_ORG_SUCCESS:
      const { [payload.id]: deleted, ...rest } = state.entities

      return {
        selectedOrganization: null,
        entities: rest,
      }
    default:
      return state
  }
}

export default organizations
