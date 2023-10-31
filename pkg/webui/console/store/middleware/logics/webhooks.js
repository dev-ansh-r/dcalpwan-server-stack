

import tts from '@console/api/tts'

import createRequestLogic from '@ttn-lw/lib/store/logics/create-request-logic'

import * as webhooks from '@console/store/actions/webhooks'
import * as webhookFormats from '@console/store/actions/webhook-formats'
import * as webhookTemplates from '@console/store/actions/webhook-templates'

const getWebhookLogic = createRequestLogic({
  type: webhooks.GET_WEBHOOK,
  process: async ({ action }) => {
    const {
      payload: { appId, webhookId },
      meta: { selector },
    } = action

    return await tts.Applications.Webhooks.getById(appId, webhookId, selector)
  },
})

const getWebhooksLogic = createRequestLogic({
  type: webhooks.GET_WEBHOOKS_LIST,
  process: async ({ action }) => {
    const {
      payload: { appId },
      meta: { selector },
    } = action
    const res = await tts.Applications.Webhooks.getAll(appId, selector)

    return { entities: res.webhooks, totalCount: res.totalCount }
  },
})

const updateWebhookLogic = createRequestLogic({
  type: webhooks.UPDATE_WEBHOOK,
  process: async ({ action }) => {
    const { appId, webhookId, patch } = action.payload

    return await tts.Applications.Webhooks.updateById(appId, webhookId, patch)
  },
})

const getWebhookFormatsLogic = createRequestLogic({
  type: webhookFormats.GET_WEBHOOK_FORMATS,
  process: async () => {
    const { formats } = await tts.Applications.Webhooks.getFormats()

    return formats
  },
})

const getWebhookTemplateLogic = createRequestLogic({
  type: webhookTemplates.GET_WEBHOOK_TEMPLATE,
  process: async ({ action }) => {
    const { id } = action.payload
    const { selector } = action.meta

    return await tts.Applications.Webhooks.getTemplate(id, selector)
  },
})

const getWebhookTemplatesLogic = createRequestLogic({
  type: webhookTemplates.LIST_WEBHOOK_TEMPLATES,
  process: async ({ action }) => {
    const { selector } = action.meta
    const { templates } = await tts.Applications.Webhooks.listTemplates(selector)

    return templates
  },
})

const createWebhookLogic = createRequestLogic({
  type: webhooks.CREATE_WEBHOOK,
  process: async ({ action }) => {
    const { appId, webhook } = action.payload

    return await tts.Applications.Webhooks.create(appId, webhook)
  },
})

export default [
  getWebhookLogic,
  getWebhooksLogic,
  updateWebhookLogic,
  getWebhookFormatsLogic,
  getWebhookTemplateLogic,
  getWebhookTemplatesLogic,
  createWebhookLogic,
]
