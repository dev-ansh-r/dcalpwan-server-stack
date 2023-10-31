

import CreateFetchSelect from '@console/containers/fetch-select'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { getGsFrequencyPlans, getNsFrequencyPlans } from '@console/store/actions/configuration'

import {
  selectGsFrequencyPlans,
  selectNsFrequencyPlans,
  selectFrequencyPlansError,
  selectFrequencyPlansFetching,
} from '@console/store/selectors/configuration'

import { formatOptions, m } from './utils'

export const CreateFrequencyPlansSelect = (source, options = {}) =>
  CreateFetchSelect({
    fetchOptions: source === 'ns' ? getNsFrequencyPlans : getGsFrequencyPlans,
    optionsSelector: source === 'ns' ? selectNsFrequencyPlans : selectGsFrequencyPlans,
    fetchingSelector: selectFrequencyPlansFetching,
    errorSelector: selectFrequencyPlansError,
    defaultWarning: m.warning,
    defaultTitle: sharedMessages.frequencyPlan,
    optionsFormatter: formatOptions,
    additionalOptions: source === 'gs' ? [{ value: 'no-frequency-plan', label: m.none }] : [],
    ...options,
  })

export const GsFrequencyPlansSelect = CreateFrequencyPlansSelect('gs')
export const NsFrequencyPlansSelect = CreateFrequencyPlansSelect('ns')
