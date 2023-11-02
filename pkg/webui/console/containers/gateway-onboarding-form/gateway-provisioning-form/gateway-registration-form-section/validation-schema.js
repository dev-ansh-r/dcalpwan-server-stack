

import { defineMessages } from 'react-intl'

import Yup from '@ttn-lw/lib/yup'
import sharedMessages from '@ttn-lw/lib/shared-messages'
import { id as gatewayIdRegexp } from '@ttn-lw/lib/regexp'
import { selectGsConfig } from '@ttn-lw/lib/selectors/env'

const gsEnabled = selectGsConfig().enabled

const m = defineMessages({
  validateEntry: 'There must be at least one selected frequency plan ID.',
})

const hasAtLeastOneEntry = frequencyPlanIds =>
  frequencyPlanIds.length > 0 && frequencyPlanIds.some(entry => entry !== '' && entry !== undefined)

const validationSchema = Yup.object().shape({
  ids: Yup.object().shape({
    gateway_id: Yup.string()
      .min(3, Yup.passValues(sharedMessages.validateTooShort))
      .max(36, Yup.passValues(sharedMessages.validateTooLong))
      .matches(gatewayIdRegexp, Yup.passValues(sharedMessages.validateIdFormat))
      .required(sharedMessages.validateRequired),
    eui: Yup.string()
      .test(
        'has 16, 12 or 0 characters',
        Yup.passValues(sharedMessages.validateLength)({ length: 16 }),
        value => value === undefined || value.length === 12 || value.length === 16,
      )
      .test(
        "doesn't have 12 characters",
        Yup.passValues(sharedMessages.validateMacAddressEntered),
        value => value?.length !== 12,
      ),
  }),
  name: Yup.string()
    .min(2, Yup.passValues(sharedMessages.validateTooShort))
    .max(50, Yup.passValues(sharedMessages.validateTooLong)),
  require_authenticated_connection: Yup.boolean(),
  location_public: Yup.boolean(),
  status_public: Yup.boolean(),
  frequency_plan_ids: gsEnabled
    ? Yup.array()
        .of(Yup.string().max(64, Yup.passValues(sharedMessages.validateTooLong)))
        .max(8, Yup.passValues(sharedMessages.attributesValidateTooMany))
        .test('has at least one entry', m.validateEntry, hasAtLeastOneEntry)
    : Yup.array(),
})

export default validationSchema
