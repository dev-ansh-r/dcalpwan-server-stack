

import React from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { useSelector } from 'react-redux'
import { useParams } from 'react-router-dom'

import PageTitle from '@ttn-lw/components/page-title'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import RequireRequest from '@ttn-lw/lib/components/require-request'

import WebhookEdit from '@console/containers/webhook-edit'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import { getWebhookTemplateId } from '@ttn-lw/lib/selectors/id'

import { getWebhook } from '@console/store/actions/webhooks'

import { selectSelectedWebhook } from '@console/store/selectors/webhooks'
import { selectWebhookTemplateById } from '@console/store/selectors/webhook-templates'
import {
  selectWebhooksHealthStatusEnabled,
  selectWebhookRetryInterval,
  selectWebhookHasUnhealthyConfig,
} from '@console/store/selectors/application-server'

const ApplicationWebhookEditInner = () => {
  const { appId, webhookId } = useParams()
  const healthStatusEnabled = useSelector(selectWebhooksHealthStatusEnabled)
  const webhookRetryInterval = useSelector(selectWebhookRetryInterval)
  const hasUnhealthyWebhookConfig = useSelector(selectWebhookHasUnhealthyConfig)
  const webhook = useSelector(selectSelectedWebhook)
  const webhookTemplateId = getWebhookTemplateId(webhook)
  const webhookTemplate = useSelector(state => selectWebhookTemplateById(state, webhookTemplateId))
  useBreadcrumbs(
    'apps.single.integrations.webhooks.edit',
    <Breadcrumb
      path={`/applications/${appId}/integrations/${webhookId}`}
      content={sharedMessages.edit}
    />,
  )

  return (
    <Container>
      <PageTitle
        title={sharedMessages.editWebhook}
        className="mb-0"
        hideHeading={Boolean(webhookTemplate)}
      />
      <Row>
        <Col lg={8} md={12}>
          <WebhookEdit
            update
            appId={appId}
            selectedWebhook={webhook}
            webhookId={webhookId}
            webhookTemplate={webhookTemplate}
            healthStatusEnabled={healthStatusEnabled}
            webhookRetryInterval={webhookRetryInterval}
            hasUnhealthyWebhookConfig={hasUnhealthyWebhookConfig}
          />
        </Col>
      </Row>
    </Container>
  )
}

const webhookEntitySelector = [
  'base_url',
  'format',
  'headers',
  'downlink_api_key',
  'uplink_message',
  'uplink_normalized',
  'join_accept',
  'downlink_ack',
  'downlink_nack',
  'downlink_sent',
  'downlink_failed',
  'downlink_queued',
  'downlink_queue_invalidated',
  'location_solved',
  'service_data',
  'template_ids',
  'health_status',
  'field_mask',
]

const ApplicationWebhookEdit = () => {
  const { appId, webhookId } = useParams()

  return (
    <RequireRequest requestAction={getWebhook(appId, webhookId, webhookEntitySelector)}>
      <ApplicationWebhookEditInner />
    </RequireRequest>
  )
}

export default ApplicationWebhookEdit
