

import React, { useCallback } from 'react'
import { useSelector } from 'react-redux'

import gatewayIcon from '@assets/misc/gateway.svg'

import RequireRequest from '@ttn-lw/lib/components/require-request'

import EntityTitleSection from '@console/components/entity-title-section'

import GatewayConnection from '@console/containers/gateway-connection'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'
import { selectCollaboratorsTotalCount } from '@ttn-lw/lib/store/selectors/collaborators'
import { getCollaboratorsList } from '@ttn-lw/lib/store/actions/collaborators'

import {
  checkFromState,
  mayViewOrEditGatewayApiKeys,
  mayViewOrEditGatewayCollaborators,
} from '@console/lib/feature-checks'

import { getApiKeysList } from '@console/store/actions/api-keys'

import { selectApiKeysTotalCount } from '@console/store/selectors/api-keys'
import { selectGatewayById } from '@console/store/selectors/gateways'

const { Content } = EntityTitleSection

const GatewayTitleSection = ({ gtwId }) => {
  const apiKeysTotalCount = useSelector(selectApiKeysTotalCount)
  const collaboratorsTotalCount = useSelector(state =>
    selectCollaboratorsTotalCount(state, { id: gtwId }),
  )
  const mayViewCollaborators = useSelector(state =>
    checkFromState(mayViewOrEditGatewayCollaborators, state),
  )
  const mayViewApiKeys = useSelector(state => checkFromState(mayViewOrEditGatewayApiKeys, state))
  const gateway = useSelector(state => selectGatewayById(state, gtwId))

  const bottomBarLeft = <GatewayConnection gtwId={gtwId} />
  const bottomBarRight = (
    <>
      {mayViewCollaborators && (
        <Content.EntityCount
          icon="collaborators"
          value={collaboratorsTotalCount}
          keyMessage={sharedMessages.collaboratorCounted}
          toAllUrl={`/gateways/${gtwId}/collaborators`}
        />
      )}
      {mayViewApiKeys && (
        <Content.EntityCount
          icon="api_keys"
          value={apiKeysTotalCount}
          keyMessage={sharedMessages.apiKeyCounted}
          toAllUrl={`/gateways/${gtwId}/api-keys`}
        />
      )}
    </>
  )

  const loadData = useCallback(
    async dispatch => {
      if (mayViewCollaborators) {
        dispatch(getCollaboratorsList('gateway', gtwId))
      }

      if (mayViewApiKeys) {
        dispatch(getApiKeysList('gateway', gtwId))
      }
    },
    [gtwId, mayViewApiKeys, mayViewCollaborators],
  )

  return (
    <RequireRequest requestAction={loadData}>
      <EntityTitleSection
        id={gtwId}
        name={gateway.name}
        icon={gatewayIcon}
        iconAlt={sharedMessages.gateway}
      >
        <Content
          creationDate={gateway.created_at}
          bottomBarLeft={bottomBarLeft}
          bottomBarRight={bottomBarRight}
        />
      </EntityTitleSection>
    </RequireRequest>
  )
}

GatewayTitleSection.propTypes = {
  gtwId: PropTypes.string.isRequired,
}

export default GatewayTitleSection
