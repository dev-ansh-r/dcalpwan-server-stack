

import { useSelector } from 'react-redux'

import { APPLICATION, GATEWAY, ORGANIZATION, CLIENT } from '@console/constants/entities'

import { selectCollaboratorById } from '@ttn-lw/lib/store/selectors/collaborators'

import { selectUserId } from '@account/store/selectors/user'
import { selectUserById } from '@console/store/selectors/users'
import {
  selectGatewayRights,
  selectGatewayPseudoRights,
  selectGatewayRightsError,
} from '@console/store/selectors/gateways'
import {
  selectApplicationPseudoRights,
  selectApplicationRights,
  selectApplicationRightsError,
} from '@console/store/selectors/applications'
import {
  selectOrganizationPseudoRights,
  selectOrganizationRights,
  selectOrganizationRightsError,
} from '@console/store/selectors/organizations'
import { selectClientPseudoRights, selectClientRights } from '@account/store/selectors/clients'

const sdkServices = {
  [APPLICATION]: 'Applications',
  [GATEWAY]: 'Gateways',
  [ORGANIZATION]: 'Organizations',
  [CLIENT]: 'Clients',
}

const isCollaboratorUser = collaborator => collaborator.ids && 'user_ids' in collaborator.ids

const useCollaboratorData = (entity, entityId, collaboratorId, tts) => {
  const rightsSelector = {
    [GATEWAY]: selectGatewayRights,
    [APPLICATION]: selectApplicationRights,
    [ORGANIZATION]: selectOrganizationRights,
    [CLIENT]: selectClientRights,
  }
  const pseudoRightsSelector = {
    [GATEWAY]: selectGatewayPseudoRights,
    [APPLICATION]: selectApplicationPseudoRights,
    [ORGANIZATION]: selectOrganizationPseudoRights,
    [CLIENT]: selectClientPseudoRights,
  }
  const righsErrorSelector = {
    [GATEWAY]: selectGatewayRightsError,
    [APPLICATION]: selectApplicationRightsError,
    [ORGANIZATION]: selectOrganizationRightsError,
    [CLIENT]: selectOrganizationRightsError,
  }

  const collaborator = useSelector(state => selectCollaboratorById(state, collaboratorId))
  const rights = useSelector(rightsSelector[entity])
  const pseudoRights = useSelector(pseudoRightsSelector[entity])
  const error = useSelector(righsErrorSelector[entity])
  const update = Boolean(collaborator)
  const isUser = update && isCollaboratorUser(collaborator)
  const admin = useSelector(state => selectUserById(state, collaboratorId))?.admin
  const isAdmin = isUser && admin
  const currentUserId = useSelector(selectUserId)
  const isCurrentUser = isUser && currentUserId === collaboratorId
  const updateCollaborator = patch => tts[sdkServices[entity]].Collaborators.update(entityId, patch)
  const removeCollaborator = (isUser, collaboratorIds) => {
    tts[sdkServices[entity]].Collaborators.remove(isUser, entityId, collaboratorIds)
  }

  return {
    collaborator,
    isCollaboratorUser: isUser,
    isCollaboratorAdmin: isAdmin,
    isCollaboratorCurrentUser: isCurrentUser,
    currentUserId,
    rights,
    pseudoRights,
    error,
    updateCollaborator,
    removeCollaborator,
  }
}

export default useCollaboratorData
