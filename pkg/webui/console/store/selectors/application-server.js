

import { get } from 'lodash'

const selectAsStore = state => state.as

export const selectAsConfiguration = state => selectAsStore(state).configuration

export const selectPubSubProviders = state =>
  get(selectAsConfiguration(state), 'pubsub.providers') || {}
export const selectNatsProviderDisabled = state =>
  selectPubSubProviders(state).nats === 'DISABLED' || false
export const selectMqttProviderDisabled = state =>
  selectPubSubProviders(state).mqtt === 'DISABLED' || false

export const selectWebhooksConfiguration = state => selectAsConfiguration(state).webhooks
export const selectWebhooksHealthStatusEnabled = state => {
  const webhooksConfig = selectWebhooksConfiguration(state)

  if (!webhooksConfig) {
    return false
  }

  return (
    webhooksConfig.unhealthy_retry_interval !== '0s' ||
    'unhealthy_attempts_threshold' in webhooksConfig
  )
}

export const selectWebhookHasUnhealthyConfig = state => {
  const webhooksConfig = selectWebhooksConfiguration(state)

  if (!webhooksConfig) {
    return false
  }

  return (
    webhooksConfig.unhealthy_retry_interval !== '0s' &&
    'unhealthy_attempts_threshold' in webhooksConfig
  )
}

export const selectWebhookRetryInterval = state => {
  const webhooksConfig = selectWebhooksConfiguration(state)

  if (!webhooksConfig) {
    return false
  }

  return 'unhealthy_attempts_threshold' in webhooksConfig
    ? selectWebhooksConfiguration(state).unhealthy_retry_interval
    : null
}
