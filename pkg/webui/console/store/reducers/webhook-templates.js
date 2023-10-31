

import { getWebhookTemplateId } from '@ttn-lw/lib/selectors/id'

import {
  LIST_WEBHOOK_TEMPLATES_SUCCESS,
  GET_WEBHOOK_TEMPLATE_SUCCESS,
} from '@console/store/actions/webhook-templates'

const defaultState = {
  entities: undefined,
}

const webhookTemplates = (state = defaultState, { type, payload }) => {
  switch (type) {
    case LIST_WEBHOOK_TEMPLATES_SUCCESS:
      const entities = payload.reduce(
        (acc, template) => {
          acc[getWebhookTemplateId(template)] = template

          return acc
        },
        { ...state.entities },
      )
      return {
        ...state,
        entities,
      }
    case GET_WEBHOOK_TEMPLATE_SUCCESS:
      return {
        ...state,
        entities: {
          ...state.entities,
          [getWebhookTemplateId(payload)]: payload,
        },
      }
    default:
      return state
  }
}

export default webhookTemplates
