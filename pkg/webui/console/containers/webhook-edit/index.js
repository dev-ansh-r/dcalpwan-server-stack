

import React, { useState } from 'react'
import { useDispatch } from 'react-redux'
import { useNavigate } from 'react-router-dom'

import tts from '@console/api/tts'

import toast from '@ttn-lw/components/toast'

import WebhookForm from '@console/components/webhook-form'

import diff from '@ttn-lw/lib/diff'
import PropTypes from '@ttn-lw/lib/prop-types'
import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import { updateWebhook } from '@console/store/actions/webhooks'

const WebhookEdit = props => {
  const {
    selectedWebhook,
    appId,
    webhookTemplate,
    healthStatusEnabled,
    webhookId,
    hasUnhealthyWebhookConfig,
    webhookRetryInterval,
  } = props

  const dispatch = useDispatch()
  const navigate = useNavigate()

  const handleUpdateWebhook = React.useCallback(
    async updatedWebhook => {
      const patch = diff(selectedWebhook, updatedWebhook, {
        exclude: ['ids'],
        patchArraysItems: false,
        patchInFull: ['headers'],
      })

      if (Object.keys(patch).length === 0) {
        await dispatch(attachPromise(updateWebhook(appId, webhookId, updateWebhook)))
      } else {
        await dispatch(attachPromise(updateWebhook(appId, webhookId, patch)))
      }
    },
    [selectedWebhook, dispatch, appId, webhookId],
  )
  const showSuccessToast = React.useCallback(() => {
    toast({
      message: sharedMessages.webhookUpdated,
      type: toast.types.SUCCESS,
    })
  }, [])

  const [error, setError] = useState()
  const handleWebhookSubmit = React.useCallback(
    async (values, newWebhook, { resetForm }) => {
      setError(undefined)
      try {
        const result = await handleUpdateWebhook(newWebhook)
        resetForm({ values })
        showSuccessToast(result)
      } catch (error) {
        resetForm({ values })
        setError(error)
      }
    },
    [handleUpdateWebhook, showSuccessToast],
  )

  const handleDelete = React.useCallback(async () => {
    await tts.Applications.Webhooks.deleteById(appId, webhookId)
  }, [appId, webhookId])
  const handleDeleteSuccess = React.useCallback(() => {
    toast({
      message: sharedMessages.webhookDeleted,
      type: toast.types.SUCCESS,
    })

    navigate(`/applications/${appId}/integrations/webhooks`)
  }, [appId, navigate])

  const handleReactivateSuccess = React.useCallback(() => {
    toast({
      message: sharedMessages.reactivateSuccess,
      type: toast.types.SUCCESS,
    })
  }, [])

  const handleReactivate = React.useCallback(
    async updatedHealthStatus => {
      await dispatch(
        attachPromise(updateWebhook(appId, webhookId, updatedHealthStatus, ['health_status'])),
      )
    },
    [appId, dispatch, webhookId],
  )

  return (
    <WebhookForm
      update
      appId={appId}
      initialWebhookValue={selectedWebhook}
      webhookTemplate={webhookTemplate}
      onSubmit={handleWebhookSubmit}
      onDelete={handleDelete}
      onDeleteSuccess={handleDeleteSuccess}
      onReactivate={handleReactivate}
      onReactivateSuccess={handleReactivateSuccess}
      healthStatusEnabled={healthStatusEnabled}
      webhookRetryInterval={webhookRetryInterval}
      hasUnhealthyWebhookConfig={hasUnhealthyWebhookConfig}
      error={error}
    />
  )
}

WebhookEdit.propTypes = {
  appId: PropTypes.string.isRequired,
  hasUnhealthyWebhookConfig: PropTypes.bool.isRequired,
  healthStatusEnabled: PropTypes.bool.isRequired,
  selectedWebhook: PropTypes.webhook.isRequired,
  webhookId: PropTypes.string.isRequired,
  webhookRetryInterval: PropTypes.string,
  webhookTemplate: PropTypes.webhookTemplate,
}

WebhookEdit.defaultProps = {
  webhookTemplate: undefined,
  webhookRetryInterval: null,
}

export default WebhookEdit
