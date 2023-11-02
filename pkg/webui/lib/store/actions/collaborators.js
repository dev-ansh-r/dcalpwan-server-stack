

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'
import { createPaginationByParentRequestActions } from '@ttn-lw/lib/store/actions/pagination'

export const SHARED_NAME = 'COLLABORATORS'

export const GET_COLLABORATOR_BASE = 'GET_COLLABORATOR'
export const [
  {
    request: GET_COLLABORATOR,
    success: GET_COLLABORATOR_SUCCESS,
    failure: GET_COLLABORATOR_FAILURE,
  },
  { request: getCollaborator, success: getCollaboratorSuccess, failure: getCollaboratorFailure },
] = createRequestActions(
  GET_COLLABORATOR_BASE,
  (parentType, parentId, collaboratorId, isUser) => ({
    parentType,
    parentId,
    collaboratorId,
    isUser,
  }),
  (parentType, parentId, collaboratorId, isUser, selector) => ({ selector }),
)

export const GET_COLLABORATORS_LIST_BASE = 'GET_COLLABORATORS_LIST'
export const [
  {
    request: GET_COLLABORATORS_LIST,
    success: GET_COLLABORATORS_LIST_SUCCESS,
    failure: GET_COLLABORATORS_LIST_FAILURE,
  },
  {
    request: getCollaboratorsList,
    success: getCollaboratorsListSuccess,
    failure: getCollaboratorsListFailure,
  },
] = createPaginationByParentRequestActions(SHARED_NAME)

export const DELETE_COLLABORATOR_BASE = 'DELETE_COLLABORATOR'
export const [
  {
    request: DELETE_COLLABORATOR,
    success: DELETE_COLLABORATOR_SUCCESS,
    failure: DELETE_COLLABORATOR_FAILURE,
  },
  {
    request: deleteCollaborator,
    success: deleteCollaboratorSuccess,
    failure: deleteCollaboratorFailure,
  },
] = createRequestActions(
  DELETE_COLLABORATOR_BASE,
  (parentType, parentId, collaboratorId, isUser) => ({
    parentType,
    parentId,
    collaboratorId,
    isUser,
  }),
  (parentType, parentId, collaboratorId, isUser, selector) => ({ selector }),
)
