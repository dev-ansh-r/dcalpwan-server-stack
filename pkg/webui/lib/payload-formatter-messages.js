

import { defineMessages } from 'react-intl'

import TYPES from '@console/constants/formatter-types'

import sharedMessages from './shared-messages'

const m = defineMessages({
  repository: 'Repository',
  javascript: 'Javascript',
  cayennelpp: 'CayenneLPP',
})

export default Object.freeze({
  [TYPES.JAVASCRIPT]: m.javascript,
  [TYPES.REPOSITORY]: m.repository,
  [TYPES.NONE]: sharedMessages.none,
  [TYPES.GRPC]: m.grpc,
  [TYPES.CAYENNELPP]: m.cayennelpp,
})
