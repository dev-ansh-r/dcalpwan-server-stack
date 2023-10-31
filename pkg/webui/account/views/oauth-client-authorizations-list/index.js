

import React from 'react'
import { Row, Col, Container } from 'react-grid-system'

import PageTitle from '@ttn-lw/components/page-title'

import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'

import OAuthClientAuthorizationsTable from '@account/containers/authorizations-table'

import sharedMessages from '@ttn-lw/lib/shared-messages'

const AuthorizationsList = () => (
  <Container>
    <Row>
      <IntlHelmet title={sharedMessages.oauthClientAuthorizations} />
      <Col>
        <PageTitle title={sharedMessages.oauthClientAuthorizations} hideHeading />
        <OAuthClientAuthorizationsTable />
      </Col>
    </Row>
  </Container>
)

export default AuthorizationsList
