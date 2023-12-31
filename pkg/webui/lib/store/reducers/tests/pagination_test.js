

import { createNamedPaginationReducerById, createNamedPaginationReducer } from '../pagination'
import {
  createPaginationRequestActions,
  createPaginationByIdRequestActions,
  createPaginationDeleteActions,
  createPaginationByIdDeleteActions,
} from '../../actions/pagination'

describe('Pagination reducers', () => {
  const NAME = 'ENTITY'
  const entityIdSelector = entity => entity.id

  describe('flat', () => {
    const reducer = createNamedPaginationReducer(NAME, entityIdSelector)
    const defaultState = { ids: [], totalCount: undefined }

    const { request, success, failure } = createPaginationRequestActions(NAME)[1]
    const {
      request: requestDelete,
      success: successDelete,
      failure: failureDelete,
    } = createPaginationDeleteActions(NAME)[1]

    it('return the initial state', () => {
      expect(reducer(undefined, { payload: {} })).toEqual(defaultState)
    })

    it('ignore the get `request` action', () => {
      expect(reducer(defaultState, request())).toEqual(defaultState)
    })

    it('ignore the get `failure` action', () => {
      expect(reducer(defaultState, failure())).toEqual(defaultState)
    })

    it('ignore the delete `request` action', () => {
      expect(reducer(defaultState, requestDelete('test-id'))).toEqual(defaultState)
    })

    it('ignore the delete `failure` action', () => {
      expect(reducer(defaultState, failureDelete())).toEqual(defaultState)
    })

    describe('when receiving the `success` action', () => {
      const entities = [
        { id: '1', name: 'name1' },
        { id: '2', name: 'name2' },
      ]
      const totalCount = entities.length
      const action = success({ entities, totalCount })

      let newState = null

      beforeAll(() => {
        newState = reducer(defaultState, action)
      })

      it('update the state', () => {
        expect(newState).not.toEqual(defaultState)
      })

      it('store ids only', () => {
        const { ids } = newState

        expect(ids).toEqual(entities.map(e => e.id))
      })

      it('store `totalCount`', () => {
        const { totalCount: newTotalCount } = newState

        expect(newTotalCount).toEqual(totalCount)
      })

      describe('when deleting an entity', () => {
        beforeAll(() => {
          newState = reducer(newState, successDelete({ id: entities[0].id }))
        })

        it('decrease `totalCount`', () => {
          expect(newState.totalCount).toEqual(entities.length - 1)
        })

        it('remove deleted id from `ids`', () => {
          expect(newState.ids).toHaveLength(entities.length - 1)
          expect(newState.ids).toEqual(expect.not.arrayContaining(entities))
        })
      })
    })
  })

  describe('when querying by id', () => {
    const reducer = createNamedPaginationReducerById(NAME, entityIdSelector)
    const defaultState = {}
    const entityId = 'parent-id'

    const { request, success, failure } = createPaginationByIdRequestActions(NAME)[1]
    const {
      request: requestDelete,
      success: successDelete,
      failure: failureDelete,
    } = createPaginationByIdDeleteActions(NAME)[1]

    it('return the initial state', () => {
      expect(reducer(undefined, { payload: {} })).toEqual(defaultState)
    })

    it('ignore the `request` action', () => {
      expect(reducer(defaultState, request(entityId))).toEqual(defaultState)
    })

    it('ignore the `failure` action', () => {
      expect(reducer(defaultState, failure({}))).toEqual(defaultState)
    })

    it('ignore the delete `request` action', () => {
      expect(reducer(defaultState, requestDelete('test-id', 'target-test-id'))).toEqual(
        defaultState,
      )
    })

    it('ignore the delete `failure` action', () => {
      expect(reducer(defaultState, failureDelete('test-id'))).toEqual(defaultState)
    })

    describe('when receiving the `success` action', () => {
      const entities = [
        { id: '1', name: 'name1' },
        { id: '2', name: 'name2' },
      ]
      const totalCount = entities.length
      const action = success({ id: entityId, entities, totalCount })

      let newState = null

      beforeAll(() => {
        newState = reducer(defaultState, action)
      })

      it('ignore without `id` in the payload', () => {
        const updatedState = reducer(defaultState, success({ entities: [], totalCount }))

        expect(updatedState).toEqual(defaultState)
      })

      it('update the state', () => {
        expect(newState).not.toEqual(defaultState)
      })

      it('store results per entity id', () => {
        const { [entityId]: results } = newState

        expect(results).not.toBeUndefined()
      })

      it('only store ids', () => {
        const { [entityId]: results } = newState

        expect(results.ids).toEqual(entities.map(entityIdSelector))
      })

      it('store `totalCount`', () => {
        const { [entityId]: results } = newState

        expect(results.totalCount).toEqual(totalCount)
      })

      describe('deletes entity', () => {
        let updatedState

        beforeAll(() => {
          updatedState = reducer(
            newState,
            successDelete({ id: entityId, targetId: entities[0].id }),
          )
        })

        it('decrease `totalCount`', () => {
          const { [entityId]: results } = updatedState
          expect(results.totalCount).toEqual(entities.length - 1)
        })

        it('remove deleted id from `ids`', () => {
          const { [entityId]: results } = updatedState
          expect(results.ids).toHaveLength(entities.length - 1)
          expect(results.ids).toEqual(expect.not.arrayContaining(entities))
        })
      })

      describe('when receiving the `success` action for another entity', () => {
        const otherEntityId = 'other-entity-id'
        const otherEntities = [
          { id: '3', name: 'name3' },
          { id: '4', name: 'name4' },
          { id: '5', name: 'name5' },
        ]
        const otherTotalCount = otherEntities.length
        const action = success({
          id: otherEntityId,
          entities: otherEntities,
          totalCount: otherTotalCount,
        })

        let otherNewState = null

        beforeAll(() => {
          otherNewState = reducer(newState, action)
        })

        it('update the state', () => {
          expect(otherNewState).not.toEqual(newState)
        })

        it('preserve previous entries', () => {
          const { [entityId]: results } = otherNewState

          expect(results).not.toBeUndefined()
          expect(results.ids).toEqual(entities.map(e => e.id))
          expect(results.totalCount).toEqual(totalCount)
        })

        it('store results per entity id', () => {
          const { [otherEntityId]: results } = otherNewState

          expect(results).not.toBeUndefined()
        })

        it('only store ids', () => {
          const { [otherEntityId]: results } = otherNewState

          expect(results.ids).toEqual(otherEntities.map(e => e.id))
        })

        it('store `totalCount`', () => {
          const { [otherEntityId]: results } = otherNewState

          expect(results.totalCount).toEqual(otherTotalCount)
        })

        describe('when deleting an entity', () => {
          let updatedState

          beforeAll(() => {
            updatedState = reducer(
              otherNewState,
              successDelete({ id: otherEntityId, targetId: otherEntities[0].id }),
            )
          })

          it('decrease `totalCount`', () => {
            const { [entityId]: results } = updatedState
            expect(results.totalCount).toEqual(otherEntities.length - 1)
          })

          it('remove deleted id from `ids`', () => {
            const { [entityId]: results } = updatedState
            expect(results.ids).toHaveLength(otherEntities.length - 1)
            expect(results.ids).toEqual(expect.not.arrayContaining(otherEntities))
          })
        })
      })
    })
  })
})
