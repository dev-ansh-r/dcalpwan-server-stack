

import React from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { useParams } from 'react-router-dom'
import { useSelector } from 'react-redux'

import { GATEWAY } from '@console/constants/entities'

import PageTitle from '@ttn-lw/components/page-title'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import RequireRequest from '@ttn-lw/lib/components/require-request'

import ConsoleCollaboratorsForm from '@console/containers/collaborators-form'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import { getCollaborator, getCollaboratorsList } from '@ttn-lw/lib/store/actions/collaborators'
import { selectCollaboratorById } from '@ttn-lw/lib/store/selectors/collaborators'

const GatewayCollaboratorEditInner = () => {
  const { gtwId, collaboratorId, collaboratorType } = useParams()

  useBreadcrumbs(
    'gtws.single.collaborators.edit',
    <Breadcrumb
      path={`/gateways/${gtwId}/collaborators/${collaboratorType}/${collaboratorId}`}
      content={sharedMessages.edit}
    />,
  )

  return (
    <Container>
      <PageTitle title={sharedMessages.collaboratorEdit} values={{ collaboratorId }} />
      <Row>
        <Col lg={8} md={12}>
          <ConsoleCollaboratorsForm
            entity={GATEWAY}
            entityId={gtwId}
            collaboratorId={collaboratorId}
            collaboratorType={collaboratorType}
            update
          />
        </Col>
      </Row>
    </Container>
  )
}

const GatewayCollaboratorEdit = () => {
  const { gtwId, collaboratorId, collaboratorType } = useParams()

  const isUser = collaboratorType === 'user'

  // Check if collaborator still exists after being possibly deleted.
  const collaborator = useSelector(state => selectCollaboratorById(state, collaboratorId))
  const hasCollaborator = Boolean(collaborator)

  return (
    <RequireRequest
      requestAction={[
        getCollaborator('gateway', gtwId, collaboratorId, isUser),
        getCollaboratorsList('gateway', gtwId),
      ]}
    >
      {hasCollaborator && <GatewayCollaboratorEditInner collaboratorType={collaboratorType} />}
    </RequireRequest>
  )
}

export default GatewayCollaboratorEdit
