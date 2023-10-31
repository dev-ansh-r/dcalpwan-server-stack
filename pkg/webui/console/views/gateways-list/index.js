

import React from 'react'
import { Container, Row, Col } from 'react-grid-system'

import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'

import GatewaysTable from '@console/containers/gateways-table'

import sharedMessages from '@ttn-lw/lib/shared-messages'

const GatewaysList = () => (
  <Container>
    <Row>
      <IntlHelmet title={sharedMessages.gateways} />
      <Col>
        <GatewaysTable />
      </Col>
    </Row>
  </Container>
)

export default GatewaysList
