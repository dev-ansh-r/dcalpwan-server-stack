

import Yup from '@ttn-lw/lib/yup'
import sharedMessages from '@ttn-lw/lib/shared-messages'
import contactSchema from '@ttn-lw/lib/shared-schemas'

import { approvalStates } from './utils'

const validationSchema = Yup.object().shape({
  owner_id: Yup.string().required(sharedMessages.validateRequired),
  ids: Yup.object().shape({
    client_id: Yup.string()
      .min(2, Yup.passValues(sharedMessages.validateTooShort))
      .max(36, Yup.passValues(sharedMessages.validateTooLong))
      .required(sharedMessages.validateRequired),
  }),
  name: Yup.string()
    .min(2, Yup.passValues(sharedMessages.validateTooShort))
    .max(2000, Yup.passValues(sharedMessages.validateTooLong)),
  description: Yup.string(),
  redirect_uris: Yup.array().max(10, Yup.passValues(sharedMessages.attributesValidateTooMany)),
  logout_redirect_uris: Yup.array().max(
    10,
    Yup.passValues(sharedMessages.attributesValidateTooMany),
  ),
  skip_authorization: Yup.bool(),
  endorsed: Yup.bool(),
  grants: Yup.array().max(3, Yup.passValues(sharedMessages.attributesValidateTooMany)),
  state: Yup.lazy(value => {
    if (value === '') {
      return Yup.string().strip()
    }

    return Yup.string().oneOf(approvalStates, sharedMessages.validateRequired)
  }),
  state_description: Yup.lazy(value => {
    if (value === '') {
      return Yup.string().strip()
    }

    return Yup.string()
  }),
  rights: Yup.array().min(1, sharedMessages.validateRights),
})

validationSchema.concat(contactSchema)

export default validationSchema
