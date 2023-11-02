

import React from 'react'
import { Row, Col, Container } from 'react-grid-system'
import { useSelector } from 'react-redux'

import PAGE_SIZES from '@ttn-lw/constants/page-sizes'

import PageTitle from '@ttn-lw/components/page-title'

import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'
import RequireRequest from '@ttn-lw/lib/components/require-request'

import ClientsTable from '@account/containers/clients-table'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { getUserRights } from '@account/store/actions/user'

import { selectUserId } from '@account/store/selectors/user'

const ClientsList = () => {
  const userId = useSelector(selectUserId)
  return (
    <RequireRequest requestAction={getUserRights(userId)}>
      <Container>
        <Row>
          <IntlHelmet title={sharedMessages.oauthClients} />
          <Col>
            <PageTitle title={sharedMessages.oauthClients} hideHeading />
            <ClientsTable pageSize={PAGE_SIZES.REGULAR} />
          </Col>
        </Row>
      </Container>
    </RequireRequest>
  )
}

export default ClientsList
