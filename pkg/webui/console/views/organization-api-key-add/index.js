

import React from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { useParams } from 'react-router-dom'

import { ORGANIZATION } from '@console/constants/entities'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import PageTitle from '@ttn-lw/components/page-title'

import { ApiKeyCreateForm } from '@console/containers/api-key-form'

import sharedMessages from '@ttn-lw/lib/shared-messages'

const OrganizationApiKeyAdd = () => {
  const { orgId } = useParams()

  useBreadcrumbs(
    'orgs.single.api-keys.add',
    <Breadcrumb path={`/organizations/${orgId}/api-keys/add`} content={sharedMessages.add} />,
  )

  return (
    <Container>
      <PageTitle title={sharedMessages.addApiKey} />
      <Row>
        <Col lg={8} md={12}>
          <ApiKeyCreateForm entity={ORGANIZATION} entityId={orgId} />
        </Col>
      </Row>
    </Container>
  )
}

export default OrganizationApiKeyAdd
