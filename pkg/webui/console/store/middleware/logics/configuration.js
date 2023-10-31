

import tts from '@console/api/tts'

import createRequestLogic from '@ttn-lw/lib/store/logics/create-request-logic'

import * as configuration from '@console/store/actions/configuration'

import {
  selectNsFrequencyPlans,
  selectGsFrequencyPlans,
} from '@console/store/selectors/configuration'

const getNsFrequencyPlansLogic = createRequestLogic({
  type: configuration.GET_NS_FREQUENCY_PLANS,
  validate: ({ getState, action }, allow, reject) => {
    const plansNs = selectNsFrequencyPlans(getState())
    if (plansNs && plansNs.length) {
      reject()
    } else {
      allow(action)
    }
  },
  process: async () => {
    const frequencyPlans = (await tts.Configuration.listNsFrequencyPlans()).frequency_plans

    return frequencyPlans
  },
})

const getGsFrequencyPlansLogic = createRequestLogic({
  type: configuration.GET_GS_FREQUENCY_PLANS,
  validate: ({ getState, action }, allow, reject) => {
    const plansGs = selectGsFrequencyPlans(getState())
    if (plansGs && plansGs.length) {
      reject()
    } else {
      allow(action)
    }
  },
  process: async () => {
    const frequencyPlans = (await tts.Configuration.listGsFrequencyPlans()).frequency_plans

    return frequencyPlans
  },
})

export default [getNsFrequencyPlansLogic, getGsFrequencyPlansLogic]
