

import autoBind from 'auto-bind'

import Marshaler from '../util/marshaler'

class Webhooks {
  constructor(registry) {
    this._api = registry
    autoBind(this)
  }

  async getAll(appId, selector) {
    const fieldMask = Marshaler.selectorToFieldMask(selector)
    const result = await this._api.List(
      {
        routeParams: { 'application_ids.application_id': appId },
      },
      fieldMask,
    )

    return Marshaler.payloadListResponse('webhooks', result)
  }

  async create(
    appId,
    webhook,
    mask = Marshaler.fieldMaskFromPatch(webhook, this._api.SetAllowedFieldMaskPaths),
  ) {
    const result = await this._api.Set(
      {
        routeParams: {
          'webhook.ids.application_ids.application_id': appId,
        },
      },
      {
        webhook,
        field_mask: Marshaler.fieldMask(mask),
      },
    )

    return Marshaler.payloadSingleResponse(result)
  }

  async getById(appId, webhookId, selector) {
    const fieldMask = Marshaler.selectorToFieldMask(selector)
    const result = await this._api.Get(
      {
        routeParams: {
          'ids.application_ids.application_id': appId,
          'ids.webhook_id': webhookId,
        },
      },
      fieldMask,
    )

    return Marshaler.payloadSingleResponse(result)
  }

  async updateById(
    appId,
    webhookId,
    patch,
    mask = Marshaler.fieldMaskFromPatch(patch, this._api.SetAllowedFieldMaskPaths),
  ) {
    const result = await this._api.Set(
      {
        routeParams: {
          'webhook.ids.application_ids.application_id': appId,
          'webhook.ids.webhook_id': webhookId,
        },
      },
      {
        webhook: patch,
        field_mask: Marshaler.fieldMask(mask),
      },
    )

    return Marshaler.payloadSingleResponse(result)
  }

  async deleteById(appId, webhookId) {
    const result = await this._api.Delete({
      routeParams: {
        'application_ids.application_id': appId,
        webhook_id: webhookId,
      },
    })

    return Marshaler.payloadSingleResponse(result)
  }

  async getFormats() {
    const result = await this._api.GetFormats()

    return Marshaler.payloadSingleResponse(result)
  }

  async listTemplates(selector) {
    const fieldMask = Marshaler.selectorToFieldMask(selector)
    const result = await this._api.ListTemplates(undefined, fieldMask)

    return Marshaler.payloadListResponse('templates', result)
  }

  async getTemplate(templateId, selector) {
    const fieldMask = Marshaler.selectorToFieldMask(selector)
    const result = await this._api.GetTemplate(
      {
        routeParams: {
          'ids.template_id': templateId,
        },
      },
      fieldMask,
    )

    return Marshaler.payloadSingleResponse(result)
  }
}

export default Webhooks
