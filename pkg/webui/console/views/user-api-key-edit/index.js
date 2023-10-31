

import React from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { useSelector } from 'react-redux'
import { useParams } from 'react-router-dom'

import { USER } from '@console/constants/entities'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import PageTitle from '@ttn-lw/components/page-title'

import RequireRequest from '@ttn-lw/lib/components/require-request'

import { ApiKeyEditForm } from '@console/containers/api-key-form'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { getUsersRightsList } from '@console/store/actions/users'
import { getApiKey } from '@console/store/actions/api-keys'

import { selectUserId } from '@account/store/selectors/user'
import { selectApiKeyById } from '@console/store/selectors/api-keys'

const UserApiKeyEditInner = () => {
  const userId = useSelector(selectUserId)
  const { apiKeyId } = useParams()

  useBreadcrumbs(
    'usr.single.api-keys.edit',
    <Breadcrumb path={`/users/api-keys/edit/${apiKeyId}`} content={sharedMessages.edit} />,
  )

  return (
    <Container>
      <PageTitle title={sharedMessages.keyEdit} />
      <Row>
        <Col lg={8} md={12}>
          <ApiKeyEditForm entity={USER} entityId={userId} />
        </Col>
      </Row>
    </Container>
  )
}

const UserApiKeyEdit = () => {
  const userId = useSelector(selectUserId)
  const { apiKeyId } = useParams()

  // Check if API key still exists after possibly being deleted.
  const apiKey = useSelector(state => selectApiKeyById(state, apiKeyId))
  const hasApiKey = Boolean(apiKey)

  return (
    <RequireRequest
      requestAction={[getUsersRightsList(userId), getApiKey('users', userId, apiKeyId)]}
    >
      {hasApiKey && <UserApiKeyEditInner />}
    </RequireRequest>
  )
}

export default UserApiKeyEdit
