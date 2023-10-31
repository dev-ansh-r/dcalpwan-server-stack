

import React, { useCallback } from 'react'
import { Container, Row, Col } from 'react-grid-system'
import { useParams } from 'react-router-dom'
import { createSelector } from 'reselect'

import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'

import ApiKeysTable from '@console/containers/api-keys-table'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { getApiKeysList } from '@console/store/actions/api-keys'

import { selectApiKeys, selectApiKeysTotalCount } from '@console/store/selectors/api-keys'

const ApplicationApiKeysList = () => {
  const { appId } = useParams()

  const getApiKeys = React.useCallback(
    filters => getApiKeysList('application', appId, filters),
    [appId],
  )

  const baseDataSelectors = createSelector(
    [selectApiKeys, selectApiKeysTotalCount],
    (keys, totalCount) => ({
      keys,
      totalCount,
    }),
  )

  const baseDataSelector = useCallback(
    state => baseDataSelectors(state, appId),
    [appId, baseDataSelectors],
  )

  return (
    <Container>
      <Row>
        <IntlHelmet title={sharedMessages.apiKeys} />
        <Col>
          <ApiKeysTable baseDataSelector={baseDataSelector} getItemsAction={getApiKeys} />
        </Col>
      </Row>
    </Container>
  )
}

export default ApplicationApiKeysList
