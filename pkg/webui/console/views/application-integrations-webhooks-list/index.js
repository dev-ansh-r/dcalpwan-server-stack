
import React from 'react'
import { Container, Row, Col } from 'react-grid-system'
import { useParams } from 'react-router-dom'

import PageTitle from '@ttn-lw/components/page-title'

import WebhooksTable from '@console/containers/webhooks-table'

import sharedMessages from '@ttn-lw/lib/shared-messages'

const ApplicationWebhooksList = () => {
  const { appId } = useParams()

  return (
    <Container>
      <PageTitle title={sharedMessages.webhooks} hideHeading />
      <Row>
        <Col>
          <WebhooksTable appId={appId} />
        </Col>
      </Row>
    </Container>
  )
}

export default ApplicationWebhooksList
