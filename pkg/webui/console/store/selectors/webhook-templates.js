

import { createSelector } from 'reselect'

import { createFetchingSelector } from '@ttn-lw/lib/store/selectors/fetching'
import { createErrorSelector } from '@ttn-lw/lib/store/selectors/error'

import {
  LIST_WEBHOOK_TEMPLATES_BASE,
  GET_WEBHOOK_TEMPLATE_BASE,
} from '@console/store/actions/webhook-templates'

const selectWebhookTemplatesStore = state => state.webhookTemplates
const selectWebhookTemplatesEntitiesStore = state => selectWebhookTemplatesStore(state).entities

export const selectWebhookTemplateById = (state, id) => {
  const entities = selectWebhookTemplatesEntitiesStore(state)
  if (!Boolean(entities)) return undefined

  return entities[id]
}
export const selectWebhookTemplateError = createErrorSelector(GET_WEBHOOK_TEMPLATE_BASE)
export const selectWebhookTemplateFetching = createFetchingSelector(GET_WEBHOOK_TEMPLATE_BASE)

export const selectWebhookTemplates = createSelector([selectWebhookTemplatesStore], state => {
  const { entities } = state

  if (!Boolean(entities)) return undefined

  return Object.keys(entities).map(key => entities[key])
})
export const selectWebhookTemplatesError = createErrorSelector(LIST_WEBHOOK_TEMPLATES_BASE)
export const selectWebhookTemplatesFetching = createFetchingSelector(LIST_WEBHOOK_TEMPLATES_BASE)
