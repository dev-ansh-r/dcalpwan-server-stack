

import React from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { useParams } from 'react-router-dom'
import { useSelector } from 'react-redux'

import { ORGANIZATION } from '@console/constants/entities'

import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import PageTitle from '@ttn-lw/components/page-title'

import RequireRequest from '@ttn-lw/lib/components/require-request'

import { ApiKeyEditForm } from '@console/containers/api-key-form'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { getApiKey } from '@console/store/actions/api-keys'

import { selectApiKeyById } from '@console/store/selectors/api-keys'

const OrganizationApiKeyEditInner = () => {
  const { orgId, apiKeyId } = useParams()

  useBreadcrumbs(
    'orgs.single.api-keys.edit',
    <Breadcrumb
      path={`/organizations/${orgId}/api-keys/${apiKeyId}`}
      content={sharedMessages.edit}
    />,
  )

  return (
    <Container>
      <PageTitle title={sharedMessages.keyEdit} />
      <Row>
        <Col lg={8} md={12}>
          <ApiKeyEditForm entity={ORGANIZATION} entityId={orgId} />
        </Col>
      </Row>
    </Container>
  )
}

const OrganizationApiKeyEdit = () => {
  const { orgId, apiKeyId } = useParams()

  // Check if API key still exists after possibly being deleted.
  const apiKey = useSelector(state => selectApiKeyById(state, apiKeyId))
  const hasApiKey = Boolean(apiKey)

  return (
    <RequireRequest requestAction={getApiKey('organization', orgId, apiKeyId)}>
      {hasApiKey && <OrganizationApiKeyEditInner />}
    </RequireRequest>
  )
}

export default OrganizationApiKeyEdit
