

import React from 'react'
import { Container, Row, Col } from 'react-grid-system'
import { useParams } from 'react-router-dom'

import PageTitle from '@ttn-lw/components/page-title'

import PubsubsTable from '@console/containers/pubsubs-table'

import sharedMessages from '@ttn-lw/lib/shared-messages'

const ApplicationPubsubsList = () => {
  const { appId } = useParams()

  return (
    <Container>
      <PageTitle title={sharedMessages.integrations} hideHeading />
      <Row>
        <Col>
          <PubsubsTable appId={appId} />
        </Col>
      </Row>
    </Container>
  )
}

export default ApplicationPubsubsList
