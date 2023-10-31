

import { defineMessages } from 'react-intl'

export const formatOptions = plans => plans.map(plan => ({ value: plan.id, label: plan.name }))
export const m = defineMessages({
  warning: 'Frequency plans unavailable',
  none: 'Do not set a frequency plan',
  selectFrequencyPlan: 'Select a frequency plan...',
  addFrequencyPlan: 'Add frequency plan',
  frequencyPlanDescription:
    'Note: most gateways use a single frequency plan. Some 16 and 64 channel gateways however allow setting multiple.',
})
