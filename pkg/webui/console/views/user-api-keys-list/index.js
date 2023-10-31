

import React, { useCallback } from 'react'
import { Container, Row, Col } from 'react-grid-system'
import { useSelector } from 'react-redux'
import { createSelector } from 'reselect'

import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'

import ApiKeysTable from '@console/containers/api-keys-table'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { getApiKeysList } from '@console/store/actions/api-keys'

import { selectApiKeys, selectApiKeysTotalCount } from '@console/store/selectors/api-keys'
import { selectUserId } from '@account/store/selectors/user'

const UserApiKeysList = () => {
  const userId = useSelector(selectUserId)

  const baseDataSelectors = createSelector(
    [selectApiKeys, selectApiKeysTotalCount],
    (keys, totalCount) => ({
      keys,
      totalCount,
    }),
  )

  const baseDataSelector = useCallback(
    state => baseDataSelectors(state, userId),
    [baseDataSelectors, userId],
  )

  const getApiKeys = React.useCallback(
    filters => getApiKeysList('users', userId, filters),
    [userId],
  )

  return (
    <Container>
      <Row>
        <IntlHelmet title={sharedMessages.personalApiKeys} />
        <Col>
          <ApiKeysTable baseDataSelector={baseDataSelector} getItemsAction={getApiKeys} />
        </Col>
      </Row>
    </Container>
  )
}

export default UserApiKeysList
