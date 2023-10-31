

import React from 'react'
import { Col, Row, Container } from 'react-grid-system'

import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'

import DevicesTable from '@console/containers/devices-table'

import sharedMessages from '@ttn-lw/lib/shared-messages'

const ApplicationDeviceList = () => (
  <Container>
    <Row>
      <IntlHelmet title={sharedMessages.devices} />
      <Col>
        <DevicesTable />
      </Col>
    </Row>
  </Container>
)

export default ApplicationDeviceList
