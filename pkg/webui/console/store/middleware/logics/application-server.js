

import tts from '@console/api/tts'

import createRequestLogic from '@ttn-lw/lib/store/logics/create-request-logic'

import * as as from '@console/store/actions/application-server'

const getAsConfigurationLogic = createRequestLogic({
  type: as.GET_AS_CONFIGURATION,
  process: async () => tts.As.getConfiguration(),
})

export default [getAsConfigurationLogic]
