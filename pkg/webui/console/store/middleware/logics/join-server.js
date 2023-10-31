

import tts from '@console/api/tts'

import createRequestLogic from '@ttn-lw/lib/store/logics/create-request-logic'

import * as js from '@console/store/actions/join-server'

const getJoinEUIPrefixesLogic = createRequestLogic({
  type: js.GET_JOIN_EUI_PREFIXES,
  process: async () => {
    const { prefixes } = await tts.Js.listJoinEUIPrefixes()

    return prefixes
  },
})

export default [getJoinEUIPrefixesLogic]
