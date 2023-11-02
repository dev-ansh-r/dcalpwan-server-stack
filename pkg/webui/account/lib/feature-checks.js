
import { selectUserIsAdmin, selectUserRights } from '@account/store/selectors/user'

export const checkFromState = (featureCheck, state) =>
  featureCheck.check(featureCheck.rightsSelector(state))

// Admin feature checks.
export const mayPerformAdminActions = {
  rightsSelector: selectUserIsAdmin,
  check: isAdmin => isAdmin,
}

export const mayPerformAllClientActions = {
  rightsSelector: selectUserRights,
  check: rights => rights.includes('RIGHT_CLIENT_ALL'),
}

export const mayEditBasicClientInformation = {
  rightsSelector: selectUserRights,
  check: rights => rights.includes('RIGHT_CLIENT_SETTINGS_BASIC'),
}

export const mayPurgeEntities = mayPerformAdminActions
