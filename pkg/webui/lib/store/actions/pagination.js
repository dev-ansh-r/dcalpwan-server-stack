

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

export const createPaginationBaseActionType = name => `GET_${name}_LIST`

export const createPaginationDeleteBaseActionType = name => `DELETE_${name}`

export const createPaginationRestoreBaseActionType = name => `RESTORE_${name}`

export const createPaginationByParentRequestActions = name =>
  createRequestActions(
    createPaginationBaseActionType(name),
    (parentType, parentId, { page, limit, query, order } = {}) => ({
      parentType,
      parentId,
      params: { page, limit, query, order },
    }),
    (parentType, parentId, params, selectors = []) => ({ selectors }),
  )

export const createPaginationByIdRequestActions = (
  name,
  requestPayloadCreator = (id, { page, limit, query, order } = {}) => ({
    id,
    params: { page, limit, query, order },
  }),
  requestMetaCreator = (id, params, selectors = [], options = {}) => ({ selectors, options }),
) =>
  createRequestActions(
    createPaginationBaseActionType(name),
    requestPayloadCreator,
    requestMetaCreator,
  )

export const createPaginationRequestActions = (
  name,
  requestPayloadCreator = ({ page, limit, query, order, deleted } = {}) => ({
    params: { page, limit, query, order, deleted },
  }),
  requestMetaCreator = (params, selectors = [], options = {}) => ({ selectors, options }),
) =>
  createRequestActions(
    createPaginationBaseActionType(name),
    requestPayloadCreator,
    requestMetaCreator,
  )

export const createPaginationDeleteActions = name =>
  createRequestActions(
    createPaginationDeleteBaseActionType(name),
    id => ({ id }),
    (id, options = {}) => ({ options }),
  )

export const createPaginationByIdDeleteActions = name =>
  createRequestActions(createPaginationDeleteBaseActionType(name), (id, targetId) => ({
    id,
    targetId,
  }))

export const createPaginationRestoreActions = name =>
  createRequestActions(
    createPaginationRestoreBaseActionType(name),
    id => ({ id }),
    (id, options = {}) => ({ options }),
  )
