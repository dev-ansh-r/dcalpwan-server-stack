

import Yup from '@ttn-lw/lib/yup'

import registerValidationSchema from './gateway-registration-form-section/validation-schema'

export const validationSchema = Yup.object({
  _ownerId: Yup.string(),
}).concat(registerValidationSchema)

export default validationSchema
