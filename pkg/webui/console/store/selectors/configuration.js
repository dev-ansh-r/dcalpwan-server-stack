

import { createFetchingSelector } from '@ttn-lw/lib/store/selectors/fetching'
import { createErrorSelector } from '@ttn-lw/lib/store/selectors/error'

import {
  GET_NS_FREQUENCY_PLANS_BASE,
  GET_GS_FREQUENCY_PLANS_BASE,
} from '@console/store/actions/configuration'

const selectConfigurationStore = state => state.configuration

export const selectGsFrequencyPlans = state => {
  const store = selectConfigurationStore(state)

  return store.gsFrequencyPlans || []
}

export const selectNsFrequencyPlans = state => {
  const store = selectConfigurationStore(state)

  return store.nsFrequencyPlans || []
}

export const selectFrequencyPlansError = createErrorSelector([
  GET_NS_FREQUENCY_PLANS_BASE,
  GET_GS_FREQUENCY_PLANS_BASE,
])

export const selectFrequencyPlansFetching = createFetchingSelector([
  GET_NS_FREQUENCY_PLANS_BASE,
  GET_GS_FREQUENCY_PLANS_BASE,
])
