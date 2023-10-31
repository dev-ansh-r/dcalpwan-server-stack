

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

export const GET_APP_PKG_DEFAULT_ASSOC_BASE = 'GET_APPLICATION_PACKAGE_DEFAULT_ASSOCIATION'
export const [
  {
    request: GET_APP_PKG_DEFAULT_ASSOC,
    success: GET_APP_PKG_DEFAULT_ASSOC_SUCCESS,
    failure: GET_APP_PKG_DEFAULT_ASSOC_FAILURE,
  },
  {
    request: getAppPkgDefaultAssoc,
    success: getAppPkgDefaultAssocSuccess,
    failure: getAppPkgDefaultAssocFailure,
  },
] = createRequestActions(
  GET_APP_PKG_DEFAULT_ASSOC_BASE,
  (appId, fPort) => ({ appId, fPort }),
  (appId, fPort, selector) => ({ selector }),
)

export const SET_APP_PKG_DEFAULT_ASSOC_BASE = 'SET_APPLICATION_PACKAGE_DEFAULT_ASSOCIATION'
export const [
  {
    request: SET_APP_PKG_DEFAULT_ASSOC,
    success: SET_APP_PKG_DEFAULT_ASSOC_SUCCESS,
    failure: SET_APP_PKG_DEFAULT_ASSOC_FAILURE,
  },
  {
    request: setAppPkgDefaultAssoc,
    success: setAppPkgDefaultAssocSuccess,
    failure: setAppPkgDefaultAssocFailure,
  },
] = createRequestActions(
  SET_APP_PKG_DEFAULT_ASSOC_BASE,
  (appId, fPort, data) => ({ appId, fPort, data }),
  (appId, fPort, data, selector) => ({ selector }),
)

export const DELETE_APP_PKG_DEFAULT_ASSOC_BASE = 'DELETE_APPLICATION_PACKAGE_DEFAULT_ASSOCIATION'
export const [
  {
    request: DELETE_APP_PKG_DEFAULT_ASSOC,
    success: DELETE_APP_PKG_DEFAULT_ASSOC_SUCCESS,
    failure: DELETE_APP_PKG_DEFAULT_ASSOC_FAILURE,
  },
  {
    request: deleteAppPkgDefaultAssoc,
    success: deleteAppPkgDefaultAssocSuccess,
    failure: deleteAppPkgDefaultAssocFailure,
  },
] = createRequestActions(DELETE_APP_PKG_DEFAULT_ASSOC_BASE, (appId, fPort) => ({
  appId,
  fPort,
}))
