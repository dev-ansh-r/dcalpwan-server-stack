

import React from 'react'
import { useSelector } from 'react-redux'
import { Container, Col, Row } from 'react-grid-system'

import PageTitle from '@ttn-lw/components/page-title'

import RequireRequest from '@ttn-lw/lib/components/require-request'

import ClientAdd from '@account/containers/oauth-client-add'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { getUserRights } from '@account/store/actions/user'

import {
  selectUserIsAdmin,
  selectUserId,
  selectUserRegularRights,
  selectUserPseudoRights,
} from '@account/store/selectors/user'

const OAuthClientAddInner = () => {
  const userId = useSelector(selectUserId)
  const isAdmin = useSelector(selectUserIsAdmin)
  const regularRights = useSelector(selectUserRegularRights)
  const pseudoRights = useSelector(selectUserPseudoRights)

  return (
    <Container>
      <PageTitle tall title={sharedMessages.addOAuthClient} />
      <Row>
        <Col lg={8} md={12}>
          <ClientAdd
            isAdmin={isAdmin}
            userId={userId}
            rights={regularRights}
            pseudoRights={pseudoRights}
          />
        </Col>
      </Row>
    </Container>
  )
}

const OAuthClientAdd = () => {
  const userId = useSelector(selectUserId)

  return (
    <RequireRequest requestAction={getUserRights(userId)}>
      <OAuthClientAddInner />
    </RequireRequest>
  )
}

export default OAuthClientAdd
