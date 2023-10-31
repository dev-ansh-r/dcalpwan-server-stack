

import { defineMessages } from 'react-intl'

import CreateFetchSelect from '@console/containers/fetch-select'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { getPubsubFormats } from '@console/store/actions/pubsub-formats'

import {
  selectPubsubFormats,
  selectPubsubFormatsError,
  selectPubsubFormatsFetching,
} from '@console/store/selectors/pubsub-formats'

const m = defineMessages({
  warning: 'Pub/Sub formats unavailable',
})

export default CreateFetchSelect({
  fetchOptions: getPubsubFormats,
  optionsSelector: selectPubsubFormats,
  errorSelector: selectPubsubFormatsError,
  fetchingSelector: selectPubsubFormatsFetching,
  defaultWarning: m.warning,
  defaultTitle: sharedMessages.pubsubFormat,
})
