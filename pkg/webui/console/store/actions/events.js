

export const createStartEventsStreamActionType = name => `START_${name}_EVENT_STREAM`

export const createStartEventsStreamSuccessActionType = name => `START_${name}_EVENT_STREAM_SUCCESS`

export const createStartEventsStreamFailureActionType = name => `START_${name}_EVENT_STREAM_FAILURE`

export const createStopEventsStreamActionType = name => `STOP_${name}_EVENT_STREAM`

export const createPauseEventsStreamActionType = name => `PAUSE_${name}_EVENT_STREAM`

export const createResumeEventsStreamActionType = name => `RESUME_${name}_EVENT_STREAM`

export const createEventStreamClosedActionType = name => `${name}_EVENT_STREAM_CLOSED`

export const createGetEventMessageSuccessActionType = name => `GET_${name}_EVENT_MESSAGE_SUCCESS`

export const createGetEventMessageFailureActionType = name => `GET_${name}_EVENT_MESSAGE_FAILURE`

export const createClearEventsActionType = name => `CLEAR_${name}_EVENTS`

export const createSetEventsFilterActionType = name => `SET_${name}_EVENTS_FILTER`

export const startEventsStream =
  name =>
  (id, { silent } = {}) => ({
    type: createStartEventsStreamActionType(name),
    id,
    silent: silent !== undefined ? silent : false,
  })

export const startEventsStreamSuccess =
  name =>
  (id, { silent } = {}) => ({
    type: createStartEventsStreamSuccessActionType(name),
    id,
    silent: silent !== undefined ? silent : false,
  })

export const startEventsStreamFailure = name => (id, error) => ({
  type: createStartEventsStreamFailureActionType(name),
  id,
  error,
})

export const stopEventsStream = name => id => ({ type: createStopEventsStreamActionType(name), id })

export const pauseEventsStream = name => id => ({
  type: createPauseEventsStreamActionType(name),
  id,
})

export const resumeEventsStream = name => id => ({
  type: createResumeEventsStreamActionType(name),
  id,
})

export const eventStreamClosed =
  name =>
  (id, { silent } = {}) => ({
    type: createEventStreamClosedActionType(name),
    id,
    silent: silent !== undefined ? silent : false,
  })

export const getEventMessageSuccess = name => (id, event) => ({
  type: createGetEventMessageSuccessActionType(name),
  id,
  event,
})

export const getEventMessageFailure = name => (id, error) => ({
  type: createGetEventMessageFailureActionType(name),
  id,
  error,
})

export const clearEvents = name => id => ({ type: createClearEventsActionType(name), id })

export const setEventsFilter = name => (id, filterId) => ({
  type: createSetEventsFilterActionType(name),
  id,
  filterId,
})
