

import React from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { useParams } from 'react-router-dom'
import { useSelector } from 'react-redux'

import { ORGANIZATION } from '@console/constants/entities'

import PageTitle from '@ttn-lw/components/page-title'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'

import RequireRequest from '@ttn-lw/lib/components/require-request'

import ConsoleCollaboratorsForm from '@console/containers/collaborators-form'

import { selectCollaboratorById } from '@ttn-lw/lib/store/selectors/collaborators'
import sharedMessages from '@ttn-lw/lib/shared-messages'
import { getCollaborator } from '@ttn-lw/lib/store/actions/collaborators'

const OrganizationCollaboratorEditInner = () => {
  const { orgId, collaboratorId } = useParams()

  useBreadcrumbs(
    'orgs.single.collaborators.edit',
    <Breadcrumb
      path={`/organizations/${orgId}/collaborators/user/${collaboratorId}`}
      content={sharedMessages.edit}
    />,
  )

  return (
    <Container>
      <PageTitle title={sharedMessages.collaboratorEdit} values={{ collaboratorId }} />
      <Row>
        <Col lg={8} md={12}>
          <ConsoleCollaboratorsForm
            entity={ORGANIZATION}
            entityId={orgId}
            collaboratorId={collaboratorId}
            collaboratorType="user"
            update
          />
        </Col>
      </Row>
    </Container>
  )
}

const OrganizationCollaboratorEdit = () => {
  const { orgId, collaboratorId } = useParams()

  // Check if collaborator still exists after being possibly deleted.
  const collaborator = useSelector(state => selectCollaboratorById(state, collaboratorId))
  const hasCollaborator = Boolean(collaborator)

  return (
    <RequireRequest requestAction={getCollaborator('organization', orgId, collaboratorId, true)}>
      {hasCollaborator && <OrganizationCollaboratorEditInner />}
    </RequireRequest>
  )
}

export default OrganizationCollaboratorEdit
