

const selectWebhookStore = state => state.webhooks

// Webhook.
export const selectWebhookEntityStore = state => selectWebhookStore(state).entities
export const selectSelectedWebhookId = state => selectWebhookStore(state).selectedWebhook
export const selectSelectedWebhook = state =>
  selectWebhookEntityStore(state)[selectSelectedWebhookId(state)]

// Webhooks.
export const selectWebhooks = state => Object.values(selectWebhookEntityStore(state))
export const selectWebhooksTotalCount = state => selectWebhookStore(state).totalCount
