

import React, { useCallback } from 'react'
import { Container, Row, Col } from 'react-grid-system'
import { useParams } from 'react-router-dom'
import { createSelector } from 'reselect'

import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'

import CollaboratorsTable from '@console/containers/collaborators-table'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import {
  selectCollaborators,
  selectCollaboratorsTotalCount,
} from '@ttn-lw/lib/store/selectors/collaborators'
import { getCollaboratorsList } from '@ttn-lw/lib/store/actions/collaborators'

const ApplicationCollaboratorsList = () => {
  const { appId } = useParams()

  const getItemsAction = useCallback(
    filters => getCollaboratorsList('application', appId, filters),
    [appId],
  )

  const baseDataSelectors = createSelector(
    [selectCollaborators, selectCollaboratorsTotalCount],
    (collaborators, totalCount) => ({
      collaborators,
      totalCount,
    }),
  )

  const baseDataSelector = useCallback(
    state => baseDataSelectors(state, appId),
    [baseDataSelectors, appId],
  )

  return (
    <Container>
      <Row>
        <IntlHelmet title={sharedMessages.collaborators} />
        <Col>
          <CollaboratorsTable baseDataSelector={baseDataSelector} getItemsAction={getItemsAction} />
        </Col>
      </Row>
    </Container>
  )
}

export default ApplicationCollaboratorsList
