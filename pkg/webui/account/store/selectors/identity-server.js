

import { createFetchingSelector } from '@ttn-lw/lib/store/selectors/fetching'
import { createErrorSelector } from '@ttn-lw/lib/store/selectors/error'

import { GET_IS_CONFIGURATION_BASE } from '@console/store/actions/identity-server'

const selectIsStore = state => state.is

export const selectIsConfiguration = state => selectIsStore(state).configuration
export const selectIsConfigurationFetching = createFetchingSelector(GET_IS_CONFIGURATION_BASE)
export const selectIsConfigurationError = createErrorSelector(GET_IS_CONFIGURATION_BASE)

export const selectUserRegistration = state => selectIsConfiguration(state).user_registration || {}
export const selectPasswordRequirements = state =>
  selectUserRegistration(state).password_requirements || {}

export const selectProfilePictureConfiguration = state =>
  selectIsConfiguration(state).profile_picture || {}
export const selectUseGravatarConfiguration = state =>
  selectProfilePictureConfiguration(state).use_gravatar
export const selectDisableUploadConfiguration = state =>
  selectProfilePictureConfiguration(state).disable_upload
