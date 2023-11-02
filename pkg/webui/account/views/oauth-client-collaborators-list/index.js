

import React from 'react'
import { Container, Row, Col } from 'react-grid-system'
import { useParams } from 'react-router-dom'

import PAGE_SIZES from '@ttn-lw/constants/page-sizes'

import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'

import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'

import CollaboratorsTable from '@account/containers/collaborators-table'

import sharedMessages from '@ttn-lw/lib/shared-messages'

const OAuthClientCollaboratorsList = () => {
  const { clientId } = useParams()

  useBreadcrumbs(
    'clients.single.collaborators',
    <Breadcrumb
      path={`/oauth-clients/${clientId}/collaborators`}
      content={sharedMessages.collaborators}
    />,
  )

  return (
    <Container>
      <Row>
        <IntlHelmet title={sharedMessages.collaborators} />
        <Col>
          <CollaboratorsTable pageSize={PAGE_SIZES.REGULAR} />
        </Col>
      </Row>
    </Container>
  )
}

export default OAuthClientCollaboratorsList
