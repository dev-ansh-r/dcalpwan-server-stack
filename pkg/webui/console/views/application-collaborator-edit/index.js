

import React from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { useParams } from 'react-router-dom'
import { useSelector } from 'react-redux'

import { APPLICATION } from '@console/constants/entities'

import PageTitle from '@ttn-lw/components/page-title'

import RequireRequest from '@ttn-lw/lib/components/require-request'
import GenericNotFound from '@ttn-lw/lib/components/full-view-error/not-found'

import ConsoleCollaboratorsForm from '@console/containers/collaborators-form'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import { getCollaborator, getCollaboratorsList } from '@ttn-lw/lib/store/actions/collaborators'
import { selectCollaboratorById } from '@ttn-lw/lib/store/selectors/collaborators'

const ApplicationCollaboratorEditInner = () => {
  const { appId, collaboratorId } = useParams()

  return (
    <Container>
      <PageTitle title={sharedMessages.collaboratorEdit} values={{ collaboratorId }} />
      <Row>
        <Col lg={8} md={12}>
          <ConsoleCollaboratorsForm
            entity={APPLICATION}
            entityId={appId}
            collaboratorId={collaboratorId}
            collaboratorType="user"
            update
          />
        </Col>
      </Row>
    </Container>
  )
}

const ApplicationCollaboratorEdit = () => {
  const { appId, collaboratorId, collaboratorType } = useParams()

  // Check if collaborator still exists after being possibly deleted.
  const collaborator = useSelector(state => selectCollaboratorById(state, collaboratorId))
  const hasCollaborator = Boolean(collaborator)

  if (collaboratorType !== 'user' && collaboratorType !== 'organization') {
    return <GenericNotFound />
  }

  return (
    <RequireRequest
      requestAction={[
        getCollaborator('application', appId, collaboratorId, collaboratorType === 'user'),
        getCollaboratorsList('application', appId),
      ]}
    >
      {hasCollaborator && <ApplicationCollaboratorEditInner />}
    </RequireRequest>
  )
}

export default ApplicationCollaboratorEdit
