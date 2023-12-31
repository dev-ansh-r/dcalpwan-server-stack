

import { mergeWith } from 'lodash'

import { getPubsubId } from '@ttn-lw/lib/selectors/id'

import {
  GET_PUBSUB,
  GET_PUBSUB_SUCCESS,
  GET_PUBSUBS_LIST_SUCCESS,
  UPDATE_PUBSUB_SUCCESS,
} from '@console/store/actions/pubsubs'

const defaultState = {
  selectedPubsub: null,
  totalCount: undefined,
  entities: {},
}

const pubsubs = (state = defaultState, { type, payload }) => {
  switch (type) {
    case GET_PUBSUB:
      return {
        ...state,
        selectedPubsub: payload.pubsubId,
      }
    case GET_PUBSUB_SUCCESS:
      return {
        ...state,
        entities: {
          ...state.entities,
          [getPubsubId(payload)]: payload,
        },
      }
    case GET_PUBSUBS_LIST_SUCCESS:
      return {
        ...state,
        entities: {
          ...payload.entities.reduce((acc, pubsub) => {
            acc[getPubsubId(pubsub)] = pubsub
            return acc
          }, {}),
        },
        totalCount: payload.totalCount,
      }
    case UPDATE_PUBSUB_SUCCESS:
      const pubsubId = getPubsubId(payload)

      return {
        ...state,
        entities: {
          ...state.entities,
          [pubsubId]: mergeWith(
            {},
            state.entities[pubsubId],
            payload,
            (_, __, key, object, source) => {
              // Handle switching between nats and pubsub. We do not want keep
              // both entries, since  there can be only one provider of a pubsub
              // integration.

              // If the payload has the `mqtt` field, then remove the `nats`
              // field from the stored object.
              if (source === payload && key === 'mqtt' && object.nats) {
                delete object.nats
              }

              // If the payload has the `nats` field, then remove the `mqtt`
              // field from the stored object.
              if (source === payload && key === 'nats' && object.mqtt) {
                delete object.mqtt
              }
            },
          ),
        },
      }
    default:
      return state
  }
}

export default pubsubs
