

import React from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { defineMessages } from 'react-intl'

import PAGE_SIZES from '@ttn-lw/constants/page-sizes'

import PageTitle from '@ttn-lw/components/page-title'

import UserSessionsTable from '@account/containers/sessions-table'

const m = defineMessages({
  sessionManagement: 'Session management',
})

const SessionManagement = () => (
  <Container>
    <Row>
      <Col>
        <PageTitle title={m.sessionManagement} hideHeading />
        <UserSessionsTable pageSize={PAGE_SIZES.REGULAR} />
      </Col>
    </Row>
  </Container>
)

export default SessionManagement
