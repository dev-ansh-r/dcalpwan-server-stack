

import Yup from '@ttn-lw/lib/yup'
import sharedMessages from '@ttn-lw/lib/shared-messages'

// Validation schema of the device type repository form section.
// Please observe the following rules to keep the validation schemas maintainable:
// 1. DO NOT USE ANY TYPE CONVERSIONS HERE. Use decocer/encoder on field level instead.
//    Consider all values as backend values. Exceptions may apply in consideration.
// 2. Comment each individual validation prop and use whitespace to structure visually.
// 3. Do not use ternary assignments but use plain if statements to ensure clarity.
const validationSchema = Yup.object({
  version_ids: Yup.object({
    brand_id: Yup.string(),
    model_id: Yup.string(),
    hardware_version: Yup.string(),
    firmware_version: Yup.string(),
    band_id: Yup.string(),
  }),
  frequency_plan_id: Yup.string().required(sharedMessages.validateRequired),
})

export default validationSchema
