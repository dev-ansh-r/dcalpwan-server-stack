

import { defineMessages } from 'react-intl'

import Yup from '@ttn-lw/lib/yup'
import sharedMessages from '@ttn-lw/lib/shared-messages'

const m = defineMessages({
  validateCode: 'Claim authentication code must consist only of numbers and letters',
})

const validationSchema = Yup.object({
  target_device_id: Yup.string().required(sharedMessages.validateRequired),
  authenticated_identifiers: Yup.object().shape({
    dev_eui: Yup.string()
      .length(8 * 2, Yup.passValues(sharedMessages.validateLength))
      .required(sharedMessages.validateRequired),
    authentication_code: Yup.string().when('._claim', {
      is: true,
      then: schema =>
        schema
          .matches(/^[A-Z0-9]{1,32}$/, Yup.passValues(m.validateCode))
          .required(sharedMessages.validateRequired),
    }),
  }),
})

export default validationSchema
