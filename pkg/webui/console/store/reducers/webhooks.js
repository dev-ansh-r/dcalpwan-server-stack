

import { merge } from 'lodash'

import { getWebhookId } from '@ttn-lw/lib/selectors/id'

import {
  GET_WEBHOOK,
  GET_WEBHOOK_SUCCESS,
  GET_WEBHOOKS_LIST_SUCCESS,
  UPDATE_WEBHOOK_SUCCESS,
} from '@console/store/actions/webhooks'

const defaultState = {
  selectedWebhook: null,
  totalCount: undefined,
  entities: {},
}

const webhooks = (state = defaultState, { type, payload }) => {
  switch (type) {
    case GET_WEBHOOK:
      return {
        ...state,
        selectedWebhook: payload.webhookId,
      }
    case GET_WEBHOOK_SUCCESS:
      return {
        ...state,
        entities: {
          ...state.entities,
          [getWebhookId(payload)]: payload,
        },
      }
    case GET_WEBHOOKS_LIST_SUCCESS:
      return {
        ...state,
        entities: {
          ...payload.entities.reduce((acc, webhook) => {
            acc[getWebhookId(webhook)] = webhook
            return acc
          }, {}),
        },
        totalCount: payload.totalCount,
      }
    case UPDATE_WEBHOOK_SUCCESS:
      const webhookId = getWebhookId(payload)

      return {
        ...state,
        entities: {
          ...state.entities,
          [webhookId]: merge({}, state.entities[webhookId], payload),
        },
      }
    default:
      return state
  }
}

export default webhooks
