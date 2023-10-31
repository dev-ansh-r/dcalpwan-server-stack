

import React from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { useParams } from 'react-router-dom'
import { useSelector } from 'react-redux'

import { APPLICATION } from '@console/constants/entities'

import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import PageTitle from '@ttn-lw/components/page-title'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'

import RequireRequest from '@ttn-lw/lib/components/require-request'

import { ApiKeyEditForm } from '@console/containers/api-key-form'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { getApiKey } from '@console/store/actions/api-keys'

import { selectApiKeyById } from '@console/store/selectors/api-keys'

const ApplicationApiKeyEditInner = () => {
  const { apiKeyId, appId } = useParams()

  useBreadcrumbs(
    'apps.single.api-keys.edit',
    <Breadcrumb
      path={`/applications/${appId}/api-keys/${apiKeyId}`}
      content={sharedMessages.edit}
    />,
  )

  return (
    <Container>
      <PageTitle title={sharedMessages.keyEdit} />
      <Row>
        <Col lg={8} md={12}>
          <ApiKeyEditForm entity={APPLICATION} entityId={appId} />
        </Col>
      </Row>
    </Container>
  )
}

const ApplicationApiKeyEdit = () => {
  const { apiKeyId, appId } = useParams()

  // Check if API key still exists after possibly being deleted.
  const apiKey = useSelector(state => selectApiKeyById(state, apiKeyId))
  const hasApiKey = Boolean(apiKey)

  return (
    <RequireRequest requestAction={getApiKey('application', appId, apiKeyId)}>
      {hasApiKey && <ApplicationApiKeyEditInner />}
    </RequireRequest>
  )
}

export default ApplicationApiKeyEdit
