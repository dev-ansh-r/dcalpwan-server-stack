

import Yup from '@ttn-lw/lib/yup'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import { parseLorawanMacVersion } from '@console/lib/device-utils'

const validationSchema = Yup.object()
  .shape({
    net_id: Yup.nullableString()
      .emptyOrLength(3 * 2, Yup.passValues(sharedMessages.validateLength)) // 3 Byte hex.
      .default(null),
    root_keys: Yup.object().when(
      ['$externalJs', '$lorawanVersion', '$mayEditKeys', '$mayReadKeys'],
      ([externalJs, lorawanVersion, mayEditKeys, mayReadKeys], schema) => {
        const strippedSchema = Yup.object().strip()
        const keySchema = Yup.lazy(value =>
          !externalJs && Boolean(value) && Boolean(value.key)
            ? Yup.object().shape({
                key: Yup.string().emptyOrLength(
                  16 * 2,
                  Yup.passValues(sharedMessages.validateLength),
                ), // 16 Byte hex.
              })
            : Yup.object().strip(),
        )

        if (!mayEditKeys && !mayReadKeys) {
          return schema.strip()
        }

        if (externalJs) {
          return schema.shape({
            nwk_key: strippedSchema,
            app_key: strippedSchema,
          })
        }

        if (parseLorawanMacVersion(lorawanVersion) < 110) {
          return schema.shape({
            nwk_key: strippedSchema,
            app_key: keySchema,
          })
        }

        return schema.shape({
          nwk_key: keySchema,
          app_key: keySchema,
        })
      },
    ),
    resets_join_nonces: Yup.boolean(),
    application_server_id: Yup.string()
      .max(100, Yup.passValues(sharedMessages.validateTooLong))
      .default(''),
    application_server_kek_label: Yup.string()
      .max(2048, Yup.passValues(sharedMessages.validateTooLong))
      .default(''),
    network_server_kek_label: Yup.string()
      .max(2048, Yup.passValues(sharedMessages.validateTooLong))
      .default(''),
  })
  .noUnknown()

export default validationSchema
