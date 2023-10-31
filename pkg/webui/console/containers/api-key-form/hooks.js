

import { useSelector, useDispatch } from 'react-redux'
import { useParams } from 'react-router-dom'
import { useCallback } from 'react'

import { APPLICATION, GATEWAY, ORGANIZATION, USER } from '@console/constants/entities'
import tts from '@console/api/tts'

import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'

import { getApplicationsRightsList } from '@console/store/actions/applications'

import {
  selectApplicationRights,
  selectApplicationPseudoRights,
} from '@console/store/selectors/applications'
import { selectApiKeyById } from '@console/store/selectors/api-keys'
import { selectGatewayRights, selectGatewayPseudoRights } from '@console/store/selectors/gateways'
import {
  selectOrganizationRights,
  selectOrganizationPseudoRights,
} from '@console/store/selectors/organizations'
import { selectUserRights, selectUserPseudoRights } from '@console/store/selectors/users'

const sdkServices = {
  [APPLICATION]: 'Applications',
  [GATEWAY]: 'Gateways',
  [ORGANIZATION]: 'Organizations',
  [USER]: 'Users',
}

const useApiKeyData = (entity, entityId) => {
  const rightsSelector = {
    [APPLICATION]: selectApplicationRights,
    [GATEWAY]: selectGatewayRights,
    [ORGANIZATION]: selectOrganizationRights,
    [USER]: selectUserRights,
  }

  const pseudoRightsSelector = {
    [APPLICATION]: selectApplicationPseudoRights,
    [GATEWAY]: selectGatewayPseudoRights,
    [ORGANIZATION]: selectOrganizationPseudoRights,
    [USER]: selectUserPseudoRights,
  }

  const { apiKeyId } = useParams()

  const rights = useSelector(rightsSelector[entity])
  const pseudoRights = useSelector(pseudoRightsSelector[entity])
  const key = useSelector(state => selectApiKeyById(state, apiKeyId))

  const dispatch = useDispatch()
  const create = tts[sdkServices[entity]].ApiKeys.create
  const updateById = patch =>
    apiKeyId ? tts[sdkServices[entity]].ApiKeys.updateById(entityId, apiKeyId, patch) : null
  const deleteById = () =>
    apiKeyId ? tts[sdkServices[entity]].ApiKeys.deleteById(entityId, apiKeyId) : null
  const getRights = useCallback(
    () => dispatch(attachPromise(getApplicationsRightsList(entityId))),
    [dispatch, entityId],
  )

  return {
    rights,
    pseudoRights,
    id: apiKeyId,
    create,
    updateById,
    deleteById,
    getRights,
    apiKey: key,
  }
}

export default useApiKeyData
