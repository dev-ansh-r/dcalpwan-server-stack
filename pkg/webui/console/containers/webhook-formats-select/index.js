

import { defineMessages } from 'react-intl'

import CreateFetchSelect from '@console/containers/fetch-select'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { getWebhookFormats } from '@console/store/actions/webhook-formats'

import {
  selectWebhookFormats,
  selectWebhookFormatsError,
  selectWebhookFormatsFetching,
} from '@console/store/selectors/webhook-formats'

const m = defineMessages({
  warning: 'Webhook formats unavailable',
})

export default CreateFetchSelect({
  fetchOptions: getWebhookFormats,
  optionsSelector: selectWebhookFormats,
  errorSelector: selectWebhookFormatsError,
  fetchingSelector: selectWebhookFormatsFetching,
  defaultWarning: m.warning,
  defaultTitle: sharedMessages.webhookFormat,
})
