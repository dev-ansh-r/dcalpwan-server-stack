

import React from 'react'
import { Row, Col, Container } from 'react-grid-system'

import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'

import ApplicationsTable from '@console/containers/applications-table'

import sharedMessages from '@ttn-lw/lib/shared-messages'

const ApplicationsList = () => (
  <Container>
    <Row>
      <IntlHelmet title={sharedMessages.applications} />
      <Col>
        <ApplicationsTable />
      </Col>
    </Row>
  </Container>
)

export default ApplicationsList
