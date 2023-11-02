

import React from 'react'
import { Container, Col, Row } from 'react-grid-system'

import PageTitle from '@ttn-lw/components/page-title'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import InviteForm from '@console/containers/invite-user-form'

import Require from '@console/lib/components/require'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { maySendInvites } from '@console/lib/feature-checks'

const InvitationSend = () => {
  useBreadcrumbs(
    'admin-panel.user-management.invitations.send',
    <Breadcrumb
      path={`/admin-panel/user-management/invitations/send`}
      content={sharedMessages.sendInvitation}
    />,
  )

  return (
    <Require featureCheck={maySendInvites} otherwise={{ redirect: '/' }}>
      <Container>
        <PageTitle title={sharedMessages.invite} />
        <Row>
          <Col>
            <InviteForm />
          </Col>
        </Row>
      </Container>
    </Require>
  )
}

export default InvitationSend
