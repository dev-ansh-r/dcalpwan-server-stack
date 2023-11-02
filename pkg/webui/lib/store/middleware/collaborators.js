

import { entitySdkServiceMap } from '@console/constants/entities'

import createRequestLogic from '@ttn-lw/lib/store/logics/create-request-logic'
import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'
import * as collaborators from '@ttn-lw/lib/store/actions/collaborators'

import { getUser } from '@account/store/actions/user'

const validParentTypes = Object.keys(entitySdkServiceMap)

const parentTypeValidator = ({ action }, allow) => {
  if (!validParentTypes.includes(action.payload.parentType)) {
    // Do not reject the action but throw an error, as this is an implementation
    // error.
    throw new Error(`Invalid parent entity type ${action.payload.parentType}`)
  }
  allow(action)
}

export default tts => {
  const getCollaboratorLogic = createRequestLogic({
    type: collaborators.GET_COLLABORATOR,
    validate: parentTypeValidator,
    process: async ({ action }, dispatch) => {
      const { parentType, parentId, collaboratorId, isUser } = action.payload

      if (isUser) {
        // Also retrieve the user to be able to determine whether
        // it is an admin user.
        try {
          await dispatch(attachPromise(getUser(collaboratorId, 'admin')))
        } catch {
          // Do not fail the action if the user could not be fetched.
        }
      }

      return isUser
        ? tts[entitySdkServiceMap[parentType]].Collaborators.getByUserId(parentId, collaboratorId)
        : tts[entitySdkServiceMap[parentType]].Collaborators.getByOrganizationId(
            parentId,
            collaboratorId,
          )
    },
  })

  const getCollaboratorsLogic = createRequestLogic({
    type: collaborators.GET_COLLABORATORS_LIST,
    process: async ({ action }) => {
      const { parentType, parentId, params } = action.payload
      const result = await tts[entitySdkServiceMap[parentType]].Collaborators.getAll(
        parentId,
        params,
      )

      return { entities: result.collaborators, totalCount: result.totalCount }
    },
  })

  const deleteCollaboratorLogic = createRequestLogic({
    type: collaborators.DELETE_COLLABORATOR,
    process: async ({ action }) => {
      const { parentType, parentId, collaboratorId, isUser } = action.payload
      return await tts[entitySdkServiceMap[parentType]].Collaborators.remove(
        isUser,
        parentId,
        collaboratorId,
      )
    },
  })

  return [getCollaboratorLogic, getCollaboratorsLogic, deleteCollaboratorLogic]
}
