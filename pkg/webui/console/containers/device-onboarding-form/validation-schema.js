

import Yup from '@ttn-lw/lib/yup'

import { REGISTRATION_TYPES } from './utils'
import claimValidationSchema from './provisioning-form-section/claiming-form-section/validation-schema'
import registrationValidationSchema from './provisioning-form-section/registration-form-section/validation-schema'
import repositoryValidationSchema from './type-form-section/repository-form-section/validation-schema'
import manualValidationSchema from './type-form-section/manual-form-section/validation-schema'

const validationSchema = Yup.object({
  _registration: Yup.mixed().oneOf([REGISTRATION_TYPES.SINGLE, REGISTRATION_TYPES.MULTIPLE]),
})
  .when('._claim', {
    is: true,
    then: schema => schema.concat(claimValidationSchema),
    otherwise: schema => schema.concat(registrationValidationSchema),
  })
  .when('._inputMethod', {
    is: 'device-repository',
    then: schema => schema.concat(repositoryValidationSchema),
    otherwise: schema => schema.concat(manualValidationSchema),
  })

export default validationSchema
