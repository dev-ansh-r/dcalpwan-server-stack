

import Yup from '@ttn-lw/lib/yup'
import sharedMessages from '@ttn-lw/lib/shared-messages'
import { hasSpecial, hasUpper, hasDigit, hasMinLength, hasMaxLength } from '@ttn-lw/lib/password'

export default requirements => {
  const passwordValidation = Yup.string()
    .default('')
    .required(sharedMessages.validateRequired)
    .test(
      'min-length',
      { message: sharedMessages.validateTooShort, values: { min: requirements.min_length } },
      password => hasMinLength(password, requirements.min_length),
    )
    .test(
      'max-length',
      { message: sharedMessages.validateTooLong, values: { max: requirements.max_length } },
      password => hasMaxLength(password, requirements.max_length),
    )
    .test(
      'min-special',
      { message: sharedMessages.validateSpecial, values: { special: requirements.min_special } },
      password => hasSpecial(password, requirements.min_special),
    )
    .test(
      'min-upper',
      { message: sharedMessages.validateUppercase, values: { upper: requirements.min_uppercase } },
      password => hasUpper(password, requirements.min_uppercase),
    )
    .test(
      'min-digit',
      { message: sharedMessages.validateDigit, values: { digit: requirements.min_digits } },
      password => hasDigit(password, requirements.min_digits),
    )

  return Yup.object().shape({
    password: passwordValidation,
    confirmPassword: Yup.string()
      .default('')
      .required(sharedMessages.validateRequired)
      .oneOf([Yup.ref('password'), null], sharedMessages.validatePasswordMatch),
  })
}
