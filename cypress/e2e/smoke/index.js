

import deviceTests from './devices'
import registrationTests from './registration'
import applicationTests from './applications'
import featureToggleTests from './feature-toggles'
import gatewayTests from './gateways'
import organizationTests from './organizations'
import forgotPasswordTests from './forgot-password'
import contactInfoValidationTests from './contact-info-validation'
import authorizationTests from './authorization'
import profileSettingsTests from './profile-settings'

export default [
  ...registrationTests,
  ...applicationTests,
  ...deviceTests,
  ...featureToggleTests,
  ...gatewayTests,
  ...organizationTests,
  ...forgotPasswordTests,
  ...contactInfoValidationTests,
  ...authorizationTests,
  ...profileSettingsTests,
]
