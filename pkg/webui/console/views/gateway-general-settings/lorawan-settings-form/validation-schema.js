
import Yup from '@ttn-lw/lib/yup'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import { delay as delayRegexp } from '@console/lib/regexp'

const validationSchema = Yup.object().shape({
  frequency_plan_ids: Yup.array(),
  schedule_downlink_late: Yup.boolean().default(false),
  schedule_anytime_delay: Yup.string().matches(
    delayRegexp,
    Yup.passValues(sharedMessages.validateDelayFormat),
  ),
})

export default validationSchema
