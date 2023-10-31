

import Yup from '@ttn-lw/lib/yup'
import sharedMessages from '@ttn-lw/lib/shared-messages'
import { userId as contactIdRegex } from '@ttn-lw/lib/regexp'

const organizationSchema = Yup.object().shape({
  organization_id: Yup.string().matches(contactIdRegex, sharedMessages.validateAlphanum),
})

const userSchema = Yup.object().shape({
  user_id: Yup.string().matches(contactIdRegex, sharedMessages.validateAlphanum),
})

export const contactSchema = Yup.object().shape({
  administrative_contact: Yup.object()
    .when(['organization_ids'], {
      is: organizationIds => Boolean(organizationIds),
      then: schema => schema.concat(organizationSchema),
      otherwise: schema => schema.concat(userSchema),
    })
    .nullable()
    .required(sharedMessages.validateRequired),
  technical_contact: Yup.object()
    .when(['organization_ids'], {
      is: organizationIds => Boolean(organizationIds),
      then: schema => schema.concat(organizationSchema),
      otherwise: schema => schema.concat(userSchema),
    })
    .nullable()
    .required(sharedMessages.validateRequired),
})

export default contactSchema
