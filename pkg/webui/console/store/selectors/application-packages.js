

import { createErrorSelector } from '@ttn-lw/lib/store/selectors/error'

import { GET_APP_PKG_DEFAULT_ASSOC_BASE } from '@console/store/actions/application-packages'

const selectApplicationPackagesStore = state => state.applicationPackages

export const selectApplicationPackageDefaultAssociations = state => {
  const store = selectApplicationPackagesStore(state)

  return store.default || {}
}
export const selectApplicationPackageDefaultAssociation = (state, fPort) =>
  selectApplicationPackageDefaultAssociations(state)[fPort]

export const selectGetApplicationPackagesError = createErrorSelector(GET_APP_PKG_DEFAULT_ASSOC_BASE)
