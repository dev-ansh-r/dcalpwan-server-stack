

import reducer from '../ui/fetching'

describe('Fetching reducers', () => {
  const BASE = 'BASE_ACTION'
  const REQUEST = `${BASE}_REQUEST`
  const SUCCESS = `${BASE}_SUCCESS`
  const FAILURE = `${BASE}_FAILURE`

  it('return the initial state', () => {
    expect(reducer(undefined, {})).toEqual({})
  })

  describe('when dispatching invalid action type', () => {
    it('ignore the action', () => {
      expect(reducer({}, { type: BASE })).toEqual({})
    })
  })

  describe('when dispatching the `request` action', () => {
    it('set fetching to `true`', () => {
      expect(reducer({}, { type: REQUEST })).toEqual({
        [BASE]: true,
      })
    })

    describe('when dispatching the `success` action', () => {
      it('set fetching to `false`', () => {
        expect(reducer({}, { type: SUCCESS })).toEqual({
          [BASE]: false,
        })
      })
    })

    describe('when dispatching the `failure` action', () => {
      it('set fetching to `false`', () => {
        expect(reducer({}, { type: FAILURE })).toEqual({
          [BASE]: false,
        })
      })
    })
  })
})
