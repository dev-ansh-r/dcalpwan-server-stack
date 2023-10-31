

import { createFetchingSelector } from '@ttn-lw/lib/store/selectors/fetching'

describe('Fetching selectors', () => {
  const BASE_ACTION_TYPE = 'BASE_ACTION'
  let initialState = null

  describe('when created with a single base action type', () => {
    const selector = createFetchingSelector(BASE_ACTION_TYPE)

    beforeAll(() => {
      initialState = { ui: { fetching: {} } }
    })

    describe('when it has no fetching entries', () => {
      it('returns `false`', () => {
        expect(selector(initialState)).toBe(false)
      })
    })

    describe('when it has fetching entry', () => {
      beforeAll(() => {
        initialState.ui.fetching[BASE_ACTION_TYPE] = true
      })

      it('return `true`', () => {
        expect(selector(initialState)).toBe(true)
      })
    })
  })

  describe('when created with two base action types', () => {
    const BASE_ACTION_TYPE_OTHER = 'BASE_ACTION_OTHER'
    const selector = createFetchingSelector([BASE_ACTION_TYPE, BASE_ACTION_TYPE_OTHER])

    beforeAll(() => {
      initialState = { ui: { fetching: {} } }
    })

    describe('when it has no fetching entries', () => {
      it('return `false`', () => {
        expect(selector(initialState)).toBe(false)
      })
    })

    describe('when it has a fetching entry', () => {
      beforeAll(() => {
        initialState.ui.fetching[BASE_ACTION_TYPE] = true
      })

      it('return `true`', () => {
        expect(selector(initialState)).toBe(true)
      })

      describe('when it has two fetching entries', () => {
        beforeAll(() => {
          initialState.ui.fetching[BASE_ACTION_TYPE_OTHER] = true
        })

        it('return `true`', () => {
          expect(selector(initialState)).toBe(true)
        })
      })
    })
  })
})
