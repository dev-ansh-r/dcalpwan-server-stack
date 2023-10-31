

import React from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { useSelector } from 'react-redux'

import { USER } from '@console/constants/entities'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import PageTitle from '@ttn-lw/components/page-title'

import RequireRequest from '@ttn-lw/lib/components/require-request'

import { ApiKeyCreateForm } from '@console/containers/api-key-form'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { getUsersRightsList } from '@console/store/actions/users'

import { selectUserId } from '@account/store/selectors/user'

const UserApiKeyAddInner = () => {
  const userId = useSelector(selectUserId)

  useBreadcrumbs(
    'usr.single.api-keys.add',
    <Breadcrumb path={`/users/api-keys/add`} content={sharedMessages.add} />,
  )

  return (
    <Container>
      <PageTitle title={sharedMessages.addApiKey} />
      <Row>
        <Col lg={8} md={12}>
          <ApiKeyCreateForm entity={USER} entityId={userId} />
        </Col>
      </Row>
    </Container>
  )
}

const UserApiKeyAdd = () => {
  const userId = useSelector(selectUserId)
  return (
    <RequireRequest requestAction={getUsersRightsList(userId)}>
      <UserApiKeyAddInner />
    </RequireRequest>
  )
}

export default UserApiKeyAdd
