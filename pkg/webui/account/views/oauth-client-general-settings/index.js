

import React from 'react'
import { useSelector } from 'react-redux'
import { Col, Row, Container } from 'react-grid-system'

import PageTitle from '@ttn-lw/components/page-title'

import RequireRequest from '@ttn-lw/lib/components/require-request'

import OAuthClientEdit from '@account/containers/oauth-client-edit'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { getUserRights } from '@account/store/actions/user'
import { getIsConfiguration } from '@account/store/actions/identity-server'

import {
  selectUserIsAdmin,
  selectUserId,
  selectUserRegularRights,
  selectUserPseudoRights,
} from '@account/store/selectors/user'
import { selectSelectedClient } from '@account/store/selectors/clients'

const OAuthClientGeneralSettingsInner = () => {
  const userId = useSelector(selectUserId)
  const isAdmin = useSelector(selectUserIsAdmin)
  const regularRights = useSelector(selectUserRegularRights)
  const pseudoRights = useSelector(selectUserPseudoRights)
  const oauthClient = useSelector(selectSelectedClient)

  return (
    <Container>
      <PageTitle title={sharedMessages.generalSettings} />
      <Row>
        <Col lg={8} md={12}>
          <OAuthClientEdit
            initialValues={oauthClient}
            isAdmin={isAdmin}
            userId={userId}
            rights={regularRights}
            pseudoRights={pseudoRights}
            update
          />
        </Col>
      </Row>
    </Container>
  )
}

const OAuthClientGeneralSettings = () => {
  const userId = useSelector(selectUserId)

  return (
    <RequireRequest requestAction={[getUserRights(userId), getIsConfiguration()]}>
      <OAuthClientGeneralSettingsInner />
    </RequireRequest>
  )
}

export default OAuthClientGeneralSettings
