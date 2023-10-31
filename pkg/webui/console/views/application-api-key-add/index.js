

import React from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { useParams } from 'react-router-dom'

import { APPLICATION } from '@console/constants/entities'

import PageTitle from '@ttn-lw/components/page-title'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import { ApiKeyCreateForm } from '@console/containers/api-key-form'

import sharedMessages from '@ttn-lw/lib/shared-messages'

const ApplicationApiKeyAdd = () => {
  const { appId } = useParams()

  useBreadcrumbs(
    'apps.single.api-keys.add',
    <Breadcrumb path={`/applications/${appId}/api-keys/add`} content={sharedMessages.add} />,
  )

  return (
    <Container>
      <PageTitle title={sharedMessages.addApiKey} />
      <Row>
        <Col lg={8} md={12}>
          <ApiKeyCreateForm entityId={appId} entity={APPLICATION} />
        </Col>
      </Row>
    </Container>
  )
}

export default ApplicationApiKeyAdd
