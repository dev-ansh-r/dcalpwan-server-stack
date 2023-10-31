

import tts from '@console/api/tts'

import createRequestLogic from '@ttn-lw/lib/store/logics/create-request-logic'

import * as pubsubs from '@console/store/actions/pubsubs'
import * as pubsubFormats from '@console/store/actions/pubsub-formats'

const createPubsubLogic = createRequestLogic({
  type: pubsubs.CREATE_PUBSUB,
  process: async ({ action }) => {
    const { appId, pubsub } = action.payload

    return await tts.Applications.PubSubs.create(appId, pubsub)
  },
})

const getPubsubLogic = createRequestLogic({
  type: pubsubs.GET_PUBSUB,
  process: async ({ action }) => {
    const {
      payload: { appId, pubsubId },
      meta: { selector },
    } = action

    return await tts.Applications.PubSubs.getById(appId, pubsubId, selector)
  },
})

const getPubsubsLogic = createRequestLogic({
  type: pubsubs.GET_PUBSUBS_LIST,
  process: async ({ action }) => {
    const { appId } = action.payload
    const res = await tts.Applications.PubSubs.getAll(appId)

    return { entities: res.pubsubs, totalCount: res.totalCount }
  },
})

const updatePubsubsLogic = createRequestLogic({
  type: pubsubs.UPDATE_PUBSUB,
  process: async ({ action }) => {
    const { appId, pubsubId, patch } = action.payload

    return await tts.Applications.PubSubs.updateById(appId, pubsubId, patch)
  },
})

const getPubsubFormatsLogic = createRequestLogic({
  type: pubsubFormats.GET_PUBSUB_FORMATS,
  process: async () => {
    const { formats } = await tts.Applications.PubSubs.getFormats()

    return formats
  },
})

const deletePubsub = createRequestLogic({
  type: pubsubs.DELETE_PUBSUB,
  process: async ({ action }) => {
    const { appId, pubsubId } = action.payload

    return await tts.Applications.PubSubs.deleteById(appId, pubsubId)
  },
})

export default [
  createPubsubLogic,
  getPubsubLogic,
  getPubsubsLogic,
  updatePubsubsLogic,
  getPubsubFormatsLogic,
  deletePubsub,
]
