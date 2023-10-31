

import React from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { useParams } from 'react-router-dom'

import { CLIENT } from '@console/constants/entities'

import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import PageTitle from '@ttn-lw/components/page-title'

import RequireRequest from '@ttn-lw/lib/components/require-request'

import AccountCollaboratorsForm from '@account/containers/collaborators-form'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { getClientRights } from '@account/store/actions/clients'

const OAuthClientCollaboratorAddInner = () => {
  const { clientId } = useParams()

  useBreadcrumbs(
    'clients.single.collaborators.add',
    <Breadcrumb
      path={`/oauth-clients/${clientId}/collaborators/add`}
      content={sharedMessages.add}
    />,
  )

  return (
    <Container>
      <PageTitle title={sharedMessages.addCollaborator} />
      <Row>
        <Col lg={8} md={12}>
          <AccountCollaboratorsForm entity={CLIENT} entityId={clientId} />
        </Col>
      </Row>
    </Container>
  )
}

const OAuthClientCollaboratorAdd = () => {
  const { clientId } = useParams()

  return (
    <RequireRequest requestAction={getClientRights(clientId)}>
      <OAuthClientCollaboratorAddInner />
    </RequireRequest>
  )
}

export default OAuthClientCollaboratorAdd
