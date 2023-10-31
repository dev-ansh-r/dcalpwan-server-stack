

import tts from '@console/api/tts'

import createRequestLogic from '@ttn-lw/lib/store/logics/create-request-logic'

import * as is from '@console/store/actions/identity-server'

const getIsConfigurationLogic = createRequestLogic({
  type: is.GET_IS_CONFIGURATION,
  process: async () => tts.Is.getConfiguration(),
})

export default [getIsConfigurationLogic]
