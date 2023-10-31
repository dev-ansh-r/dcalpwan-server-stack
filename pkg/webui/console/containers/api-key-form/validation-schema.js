

import Yup from '@ttn-lw/lib/yup'
import sharedMessages from '@ttn-lw/lib/shared-messages'

const pastDateCheck = value => !value || new Date(value) > new Date()

const validationSchema = Yup.object().shape({
  name: Yup.string()
    .min(2, Yup.passValues(sharedMessages.validateTooShort))
    .max(50, Yup.passValues(sharedMessages.validateTooLong)),
  rights: Yup.array().min(1, sharedMessages.validateRights),
  expires_at: Yup.string()
    .nullable()
    .test('past date', sharedMessages.validateDateInPast, pastDateCheck),
})

export default validationSchema
