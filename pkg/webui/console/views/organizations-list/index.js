
import React from 'react'
import { Row, Col, Container } from 'react-grid-system'

import PAGE_SIZES from '@ttn-lw/constants/page-sizes'

import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'

import OrganizationsTable from '@console/containers/organizations-table'

import sharedMessages from '@ttn-lw/lib/shared-messages'

const List = () => (
  <Container>
    <Row>
      <IntlHelmet title={sharedMessages.organizations} />
      <Col>
        <OrganizationsTable pageSize={PAGE_SIZES.REGULAR} />
      </Col>
    </Row>
  </Container>
)

export default List
