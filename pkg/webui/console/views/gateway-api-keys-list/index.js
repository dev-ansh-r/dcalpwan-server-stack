
import React from 'react'
import { Container, Row, Col } from 'react-grid-system'
import { useParams } from 'react-router-dom'
import { createSelector } from 'reselect'

import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'

import ApiKeysTable from '@console/containers/api-keys-table'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { getApiKeysList } from '@console/store/actions/api-keys'

import { selectApiKeys, selectApiKeysTotalCount } from '@console/store/selectors/api-keys'

const GatewayApiKeysList = () => {
  const { gtwId } = useParams()

  const getApiKeys = React.useCallback(
    filters => getApiKeysList('gateway', gtwId, filters),
    [gtwId],
  )

  const baseDataSelector = createSelector(
    [
      // These are the input selectors
      state => selectApiKeys(state, gtwId),
      state => selectApiKeysTotalCount(state, gtwId),
    ],
    // This is the result function
    (keys, totalCount) => ({
      keys,
      totalCount,
    }),
  )

  return (
    <Container>
      <Row>
        <IntlHelmet title={sharedMessages.apiKeys} />
        <Col>
          <ApiKeysTable
            entityId={gtwId}
            baseDataSelector={baseDataSelector}
            getItemsAction={getApiKeys}
          />
        </Col>
      </Row>
    </Container>
  )
}

export default GatewayApiKeysList
