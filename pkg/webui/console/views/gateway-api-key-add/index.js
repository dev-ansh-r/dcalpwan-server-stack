

import React from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { useParams } from 'react-router-dom'

import { GATEWAY } from '@console/constants/entities'

import PageTitle from '@ttn-lw/components/page-title'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import { ApiKeyCreateForm } from '@console/containers/api-key-form'

import sharedMessages from '@ttn-lw/lib/shared-messages'

const GatewayApiKeyAdd = () => {
  const { gtwId } = useParams()

  useBreadcrumbs(
    'gtws.single.api-keys.add',
    <Breadcrumb path={`/gateways/${gtwId}/api-keys/add`} content={sharedMessages.add} />,
  )

  return (
    <Container>
      <PageTitle title={sharedMessages.addApiKey} />
      <Row>
        <Col lg={8} md={12}>
          <ApiKeyCreateForm entityId={gtwId} entity={GATEWAY} />
        </Col>
      </Row>
    </Container>
  )
}

export default GatewayApiKeyAdd
