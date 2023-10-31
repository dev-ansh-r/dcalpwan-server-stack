

import * as Sentry from '@sentry/browser'
import { createStore, applyMiddleware, compose } from 'redux'
import { createLogicMiddleware } from 'redux-logic'
import createSentryMiddleware from 'redux-sentry-middleware'
import { cloneDeepWith } from 'lodash'

import sensitiveFields from '@ttn-lw/constants/sensitive-data'

import omitDeep from '@ttn-lw/lib/omit'
import env from '@ttn-lw/lib/env'
import dev from '@ttn-lw/lib/dev'
import requestPromiseMiddleware from '@ttn-lw/lib/store/middleware/request-promise-middleware'

import { selectUserId } from '@console/store/selectors/logout'

import rootReducer from './reducers'
import logics from './middleware/logics'

const composeEnhancers = (dev && window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__) || compose
let middlewares = [requestPromiseMiddleware, createLogicMiddleware(logics)]

if (env.sentryDsn) {
  const trimEvents = state => ({
    ...state,
    events: cloneDeepWith(state.events, (value, key) => {
      if (key === 'events' && value instanceof Array) {
        // Only transfer the last 5 events to Sentry to avoid
        // `Payload too large` errors.
        return value.slice(0, 5)
      }
    }),
  })

  middlewares = [
    createSentryMiddleware(Sentry, {
      actionTransformer: action => omitDeep(action, sensitiveFields),
      stateTransformer: state => omitDeep(trimEvents(state), sensitiveFields),
      getUserContext: state => ({ user_id: selectUserId(state) }),
    }),
    ...middlewares,
  ]
}

export default () => {
  const middleware = applyMiddleware(...middlewares)

  const store = createStore(rootReducer, composeEnhancers(middleware))
  if (dev && module.hot) {
    module.hot.accept('./reducers', () => {
      store.replaceReducer(rootReducer)
    })
  }

  return store
}
