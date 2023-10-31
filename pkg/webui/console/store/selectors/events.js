

import CONNECTION_STATUS from '@console/constants/connection-status'

import { EVENT_STATUS_CLEARED } from '@console/lib/events/definitions'

const selectEventsStore = (state, entityId) => state[entityId] || {}

export const createEventsSelector = entity => (state, entityId) => {
  const store = selectEventsStore(state.events[entity], entityId)

  return store.events || []
}

export const createEventsStatusSelector = entity => (state, entityId) => {
  const store = selectEventsStore(state.events[entity], entityId)

  return store.status || CONNECTION_STATUS.UNKNOWN
}

export const createEventsPausedSelector = entity => (state, entityId) => {
  const store = selectEventsStore(state.events[entity], entityId)

  return Boolean(store.paused)
}

export const createEventsInterruptedSelector = entity => (state, entityId) => {
  const store = selectEventsStore(state.events[entity], entityId)

  return Boolean(store.interrupted)
}

export const createEventsErrorSelector = entity => (state, entityId) => {
  const store = selectEventsStore(state.events[entity], entityId)

  return store.error
}

export const createEventsTruncatedSelector = entity => (state, entityId) => {
  const store = selectEventsStore(state.events[entity], entityId)

  return Boolean(store.truncated)
}

export const createLatestEventSelector = entity => {
  const eventsSelector = createEventsSelector(entity)

  const selectLatestEvent = (state, entityId, includeSynthetic = false) => {
    const events = eventsSelector(state, entityId)

    return includeSynthetic ? events[0] : events.find(e => !e.isSynthetic)
  }

  return selectLatestEvent
}

export const createLatestClearedEventSelector = entity => {
  const eventsSelector = createEventsSelector(entity)

  const selectLatestClearedEvent = (state, entityId) => {
    const events = eventsSelector(state, entityId)

    return events.find(e => e.name === EVENT_STATUS_CLEARED)
  }

  return selectLatestClearedEvent
}

export const createInterruptedStreamsSelector = entity => state => {
  const eventsStore = state.events[entity]

  return Object.keys(eventsStore).reduce((acc, id) => {
    if (eventsStore[id].interrupted) {
      acc[id] = eventsStore[id]
    }

    return acc
  }, {})
}

export const createEventsFilterSelector = entity => (state, entityId) => {
  const store = selectEventsStore(state.events[entity], entityId)

  return store.filter
}
