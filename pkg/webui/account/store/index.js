

import * as Sentry from '@sentry/browser'
import { createStore, applyMiddleware, compose } from 'redux'
import { createLogicMiddleware } from 'redux-logic'
import createSentryMiddleware from 'redux-sentry-middleware'
import { createBrowserHistory } from 'history'

import sensitiveFields from '@ttn-lw/constants/sensitive-data'

import { selectApplicationRootPath } from '@ttn-lw/lib/selectors/env'
import omitDeep from '@ttn-lw/lib/omit'
import dev from '@ttn-lw/lib/dev'
import env from '@ttn-lw/lib/env'
import requestPromiseMiddleware from '@ttn-lw/lib/store/middleware/request-promise-middleware'

import { selectUserId } from '@account/store/selectors/user'

import rootReducer from './reducers'
import logic from './middleware'

const composeEnhancers = (dev && window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__) || compose
let middlewares = [requestPromiseMiddleware, createLogicMiddleware(logic)]

if (env.sentryDsn) {
  middlewares = [
    createSentryMiddleware(Sentry, {
      actionTransformer: action => omitDeep(action, sensitiveFields),
      stateTransformer: state => omitDeep(state, sensitiveFields),
      getUserContext: state => ({ user_id: selectUserId(state) }),
    }),
    ...middlewares,
  ]
}

export const history = createBrowserHistory({ basename: `${selectApplicationRootPath()}/` })

const middleware = applyMiddleware(...middlewares)
const store = createStore(rootReducer, composeEnhancers(middleware))
if (dev && module.hot) {
  module.hot.accept('./reducers', () => {
    store.replaceReducer(rootReducer)
  })
}

export default store
