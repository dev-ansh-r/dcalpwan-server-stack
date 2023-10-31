

import TTS, { STACK_COMPONENTS_MAP, AUTHORIZATION_MODES } from 'ttn-lw'

import toast from '@ttn-lw/components/toast'

import { selectStackConfig, selectCSRFToken } from '@ttn-lw/lib/selectors/env'

const stackConfig = selectStackConfig()
const csrfToken = selectCSRFToken()

const stack = {
  [STACK_COMPONENTS_MAP.is]: stackConfig.is.enabled ? stackConfig.is.base_url : undefined,
}

const tts = new TTS({
  authorization: {
    mode: AUTHORIZATION_MODES.SESSION,
    csrfToken,
  },
  stackConfig: stack,
  connectionType: 'http',
  proxy: false,
})

// Forward header warnings to the toast message queue.
tts.subscribe('warning', payload => {
  toast({
    title: 'Warning',
    type: toast.types.WARNING,
    message: payload,
    preventConsecutive: true,
    messageGroup: 'header-warning',
  })
})

export default tts
