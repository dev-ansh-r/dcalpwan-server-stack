
import { isObject } from 'lodash'

import { isBadRequestError, getBackendErrorDetails } from '@ttn-lw/lib/errors/utils'

export const isValidPolicy = policy => Boolean(isObject(policy) && policy.uplink && policy.downlink)

export const isNotEnabledError = error =>
  isBadRequestError(error) &&
  'details' in error &&
  getBackendErrorDetails(error).name === 'not_enabled'
