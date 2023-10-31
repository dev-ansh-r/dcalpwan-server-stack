

import { createStore, applyMiddleware, compose, combineReducers } from 'redux'
import { routerMiddleware, connectRouter } from 'connected-react-router'

const createRootReducer = history =>
  combineReducers({
    router: connectRouter(history),
  })

export default history => {
  const middleware = applyMiddleware(routerMiddleware(history))

  return createStore(createRootReducer(history), compose(middleware))
}
