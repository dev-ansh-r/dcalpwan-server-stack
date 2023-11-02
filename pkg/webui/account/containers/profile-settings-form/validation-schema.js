
import Yup from '@ttn-lw/lib/yup'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import m from './messages'

export default Yup.object().shape({
  _profile_picture_source: Yup.string()
    .oneOf(['gravatar', 'upload'])
    .when('$initialProfilePictureSource', ([ppSource], schema) => schema.default(ppSource)),
  profile_picture: Yup.object()
    .nullable()
    .when(
      ['_profile_picture_source', '$useGravatarConfig', '$disableUploadConfig'],
      ([ppSource, useGravatarConfig, uploadDisabled], schema) => {
        if (!useGravatarConfig && uploadDisabled) {
          return schema.strip()
        }
        if (ppSource === 'upload' && useGravatarConfig) {
          return schema.required(m.imageRequired)
        }

        if (ppSource === 'gravatar') {
          // To use gravatar, the profile picture value has to be `null`.
          return schema.transform(() => null)
        }

        return schema
      },
    ),
  name: Yup.string()
    .min(2, Yup.passValues(sharedMessages.validateTooShort))
    .max(50, Yup.passValues(sharedMessages.validateTooLong)),
  primary_email_address: Yup.string()
    .email(sharedMessages.validateEmail)
    .required(sharedMessages.validateRequired),
})
