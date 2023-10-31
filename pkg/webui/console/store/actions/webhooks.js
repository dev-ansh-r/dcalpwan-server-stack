

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

export const GET_WEBHOOK_BASE = 'GET_WEBHOOK'
export const [
  { request: GET_WEBHOOK, success: GET_WEBHOOK_SUCCESS, failure: GET_WEBHOOK_FAILURE },
  { request: getWebhook, success: getWebhookSuccess, failure: getWebhookFailure },
] = createRequestActions(
  GET_WEBHOOK_BASE,
  (appId, webhookId) => ({ appId, webhookId }),
  (appId, webhookId, selector) => ({ selector }),
)

export const GET_WEBHOOKS_LIST_BASE = 'GET_WEBHOOKS_LIST'
export const [
  {
    request: GET_WEBHOOKS_LIST,
    success: GET_WEBHOOKS_LIST_SUCCESS,
    failure: GET_WEBHOOKS_LIST_FAILURE,
  },
  { request: getWebhooksList, success: getWebhooksListSuccess, failure: getWebhooksListFailure },
] = createRequestActions(
  GET_WEBHOOKS_LIST_BASE,
  appId => ({ appId }),
  (appId, selector) => ({ selector }),
)

export const UPDATE_WEBHOOK_BASE = 'UPDATE_WEBHOOK'
export const [
  { request: UPDATE_WEBHOOK, success: UPDATE_WEBHOOK_SUCCESS, failure: UPDATE_WEBHOOK_FAILURE },
  { request: updateWebhook, success: updateWebhookSuccess, failure: updateWebhookFailure },
] = createRequestActions(UPDATE_WEBHOOK_BASE, (appId, webhookId, patch) => ({
  appId,
  webhookId,
  patch,
}))

export const CREATE_WEBHOOK_BASE = 'CREATE_WEBHOOK'
export const [
  { request: CREATE_WEBHOOK, success: CREATE_WEBHOOK_SUCCESS, failure: CREATE_WEBHOOK_FAILURE },
  { request: createWebhook, success: createWebhookSuccess, failure: createWebhookFailure },
] = createRequestActions(CREATE_WEBHOOK_BASE, (appId, webhook) => ({ appId, webhook }))
