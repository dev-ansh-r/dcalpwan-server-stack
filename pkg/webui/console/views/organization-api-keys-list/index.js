
import React from 'react'
import { Container, Row, Col } from 'react-grid-system'
import { useParams } from 'react-router-dom'

import PAGE_SIZES from '@ttn-lw/constants/page-sizes'

import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'

import ApiKeysTable from '@console/containers/api-keys-table'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { getApiKeysList } from '@console/store/actions/api-keys'

import { selectApiKeys, selectApiKeysTotalCount } from '@console/store/selectors/api-keys'

const OrganizationApiKeysList = () => {
  const { orgId } = useParams()

  const getApiKeys = React.useCallback(
    filters => getApiKeysList('organization', orgId, filters),
    [orgId],
  )
  const baseDataSelector = React.useCallback(
    state => {
      const id = { id: orgId }
      return {
        keys: selectApiKeys(state, id),
        totalCount: selectApiKeysTotalCount(state, id),
      }
    },
    [orgId],
  )

  return (
    <Container>
      <Row>
        <IntlHelmet title={sharedMessages.apiKeys} />
        <Col>
          <ApiKeysTable
            pageSize={PAGE_SIZES.REGULAR}
            baseDataSelector={baseDataSelector}
            getItemsAction={getApiKeys}
          />
        </Col>
      </Row>
    </Container>
  )
}

export default OrganizationApiKeysList
