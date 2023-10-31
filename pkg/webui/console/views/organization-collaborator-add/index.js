

import React from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { useParams } from 'react-router-dom'

import { ORGANIZATION } from '@console/constants/entities'

import PageTitle from '@ttn-lw/components/page-title'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import ConsoleCollaboratorsForm from '@console/containers/collaborators-form'

import sharedMessages from '@ttn-lw/lib/shared-messages'

const OrganizationCollaboratorAdd = () => {
  const { orgId } = useParams()

  useBreadcrumbs(
    'orgs.single.collaborators.add',
    <Breadcrumb path={`/organizations/${orgId}/collaborators/add`} content={sharedMessages.add} />,
  )

  return (
    <Container>
      <PageTitle title={sharedMessages.addCollaborator} />
      <Row>
        <Col lg={8} md={12}>
          <ConsoleCollaboratorsForm entity={ORGANIZATION} entityId={orgId} />
        </Col>
      </Row>
    </Container>
  )
}

export default OrganizationCollaboratorAdd
