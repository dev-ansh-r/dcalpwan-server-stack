

import { createErrorSelector } from '@ttn-lw/lib/store/selectors/error'

describe('Error selectors', () => {
  const BASE_ACTION_TYPE = 'BASE_ACTION'
  let initialState = null

  describe('when created with a single base action type', () => {
    const selector = createErrorSelector(BASE_ACTION_TYPE)

    beforeAll(() => {
      initialState = { ui: { error: {} } }
    })

    describe('has no errors', () => {
      it('should return `undefined`', () => {
        expect(selector(initialState)).toBeUndefined()
      })
    })

    describe('has error', () => {
      const error = { status: 404 }

      beforeAll(() => {
        initialState.ui.error[BASE_ACTION_TYPE] = error
      })

      it('return the error object', () => {
        expect(selector(initialState)).toEqual(error)
      })
    })
  })

  describe('when created with two base action types', () => {
    const BASE_ACTION_TYPE_OTHER = 'BASE_ACTION_OTHER'
    const selector = createErrorSelector([BASE_ACTION_TYPE, BASE_ACTION_TYPE_OTHER])

    beforeAll(() => {
      initialState = { ui: { error: {} } }
    })

    describe('when there is no error', () => {
      it('return `undefined`', () => {
        expect(selector(initialState)).toBeUndefined()
      })
    })

    describe('when there is an error', () => {
      const not_found = { status: 404 }
      const forbidden = { status: 403 }

      beforeAll(() => {
        initialState.ui.error[BASE_ACTION_TYPE] = not_found
      })

      it('return the error object', () => {
        expect(selector(initialState)).toEqual(not_found)
      })

      describe('when there are two errors', () => {
        beforeAll(() => {
          initialState.ui.error[BASE_ACTION_TYPE_OTHER] = forbidden
        })

        it('return the first error object', () => {
          expect(selector(initialState)).toEqual(not_found)
        })
      })
    })
  })
})
