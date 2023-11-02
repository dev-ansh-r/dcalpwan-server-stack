
import React from 'react'
import { Navigate } from 'react-router-dom'
import { useSelector } from 'react-redux'

import ApplicationWebhookAddForm from '@console/views/application-integrations-webhook-add-form'

import { selectWebhookTemplates } from '@console/store/selectors/webhook-templates'

const ApplicationWebhookAdd = () => {
  const hasTemplates = useSelector(state => selectWebhookTemplates(state).length > 0)

  // Forward to the template chooser, when templates have been configured.
  if (hasTemplates) {
    return <Navigate to="template" />
  }

  return <ApplicationWebhookAddForm />
}

export default ApplicationWebhookAdd
