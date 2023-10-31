

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

export const LIST_WEBHOOK_TEMPLATES_BASE = 'LIST_WEBHOOK_TEMPLATES'
export const [
  {
    request: LIST_WEBHOOK_TEMPLATES,
    success: LIST_WEBHOOK_TEMPLATES_SUCCESS,
    failure: LIST_WEBHOOK_TEMPLATES_FAILURE,
  },
  {
    request: listWebhookTemplates,
    success: listWebhookTemplatesSuccess,
    failure: listWebhookTemplatesFailure,
  },
] = createRequestActions(LIST_WEBHOOK_TEMPLATES_BASE, undefined, selector => ({ selector }))

export const GET_WEBHOOK_TEMPLATE_BASE = 'GET_WEBHOOK_TEMPLATE'
export const [
  {
    request: GET_WEBHOOK_TEMPLATE,
    success: GET_WEBHOOK_TEMPLATE_SUCCESS,
    failure: GET_WEBHOOK_TEMPLATE_FAILURE,
  },
  {
    request: getWebhookTemplate,
    success: getWebhookTemplateSuccess,
    failure: getWebhookTemplateFailure,
  },
] = createRequestActions(
  GET_WEBHOOK_TEMPLATE_BASE,
  id => ({ id }),
  (id, selector) => ({ selector }),
)
