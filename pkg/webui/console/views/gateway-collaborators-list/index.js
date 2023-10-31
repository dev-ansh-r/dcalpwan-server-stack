

import React, { useCallback } from 'react'
import { Container, Row, Col } from 'react-grid-system'
import { useParams } from 'react-router-dom'
import { createSelector } from 'reselect'

import PAGE_SIZES from '@ttn-lw/constants/page-sizes'

import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'

import CollaboratorsTable from '@console/containers/collaborators-table'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import {
  selectCollaborators,
  selectCollaboratorsTotalCount,
} from '@ttn-lw/lib/store/selectors/collaborators'
import { getCollaboratorsList } from '@ttn-lw/lib/store/actions/collaborators'

const GatewayCollaboratorsList = () => {
  const { gtwId } = useParams()

  const baseDataSelector = createSelector(
    [
      state => selectCollaborators(state, gtwId),
      state => selectCollaboratorsTotalCount(state, gtwId),
    ],
    (collaborators, totalCount) => ({
      collaborators,
      totalCount,
    }),
  )

  const getCollaborators = useCallback(
    filter => getCollaboratorsList('gateway', gtwId, filter),
    [gtwId],
  )

  return (
    <Container>
      <Row>
        <IntlHelmet title={sharedMessages.collaborators} />
        <Col>
          <CollaboratorsTable
            pageSize={PAGE_SIZES.MAX}
            baseDataSelector={baseDataSelector}
            getItemsAction={getCollaborators}
          />
        </Col>
      </Row>
    </Container>
  )
}

export default GatewayCollaboratorsList
