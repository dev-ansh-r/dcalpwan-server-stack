

import createReducer from '../events'
import { getEventMessageSuccess } from '../../actions/events'

describe('Events reducer', () => {
  const testId = 'test'
  const reducer = createReducer(testId)
  const successActionCreator = getEventMessageSuccess(testId)

  describe('when adding new events', () => {
    it('keeps events sorted by `time`', () => {
      const testId = 'test'
      const testTime = '2019-03-28T13:18:13.376022Z'
      const testDate = new Date(testTime)
      const testName = 'as.up.data.forward'

      const initialState = reducer(undefined, {})

      const event1 = { time: testTime, name: testName }
      let newState = reducer(initialState, successActionCreator(testId, event1))

      expect(newState[testId].events).toHaveLength(1)
      expect(newState[testId].events[0]).toEqual(event1)

      const event2 = { time: new Date(testDate.getTime() + 1000).toISOString(), name: testName }
      newState = reducer(newState, successActionCreator(testId, event2))

      expect(newState[testId].events).toHaveLength(2)
      expect(newState[testId].events[0]).toEqual(event2)
      expect(newState[testId].events[1]).toEqual(event1)

      const event3 = { time: new Date(testDate.getTime() - 1000).toISOString(), name: testName }
      newState = reducer(newState, successActionCreator(testId, event3))

      expect(newState[testId].events).toHaveLength(3)
      expect(newState[testId].events[0]).toEqual(event2)
      expect(newState[testId].events[1]).toEqual(event1)
      expect(newState[testId].events[2]).toEqual(event3)

      const event4 = { time: new Date(testDate.getTime() + 2000).toISOString(), name: testName }
      newState = reducer(newState, successActionCreator(testId, event4))

      expect(newState[testId].events).toHaveLength(4)
      expect(newState[testId].events[0]).toEqual(event4)
      expect(newState[testId].events[1]).toEqual(event2)
      expect(newState[testId].events[2]).toEqual(event1)
      expect(newState[testId].events[3]).toEqual(event3)

      const event5 = { time: new Date(testDate.getTime()).toISOString(), name: testName }
      newState = reducer(newState, successActionCreator(testId, event5))

      expect(newState[testId].events).toHaveLength(5)
      expect(newState[testId].events[0]).toEqual(event4)
      expect(newState[testId].events[1]).toEqual(event2)
      expect(newState[testId].events[2]).toEqual(event5)
      expect(newState[testId].events[3]).toEqual(event1)
      expect(newState[testId].events[4]).toEqual(event3)
    })
  })
})
