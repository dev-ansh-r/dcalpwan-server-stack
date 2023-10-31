

import TTS, { STACK_COMPONENTS_MAP, AUTHORIZATION_MODES } from 'ttn-lw'

import toast from '@ttn-lw/components/toast'

import { selectStackConfig } from '@ttn-lw/lib/selectors/env'

import { token } from '.'

const stackConfig = selectStackConfig()

const stack = {
  [STACK_COMPONENTS_MAP.is]: stackConfig.is.enabled ? stackConfig.is.base_url : undefined,
  [STACK_COMPONENTS_MAP.gs]: stackConfig.gs.enabled ? stackConfig.gs.base_url : undefined,
  [STACK_COMPONENTS_MAP.ns]: stackConfig.ns.enabled ? stackConfig.ns.base_url : undefined,
  [STACK_COMPONENTS_MAP.as]: stackConfig.as.enabled ? stackConfig.as.base_url : undefined,
  [STACK_COMPONENTS_MAP.js]: stackConfig.js.enabled ? stackConfig.js.base_url : undefined,
  [STACK_COMPONENTS_MAP.edtc]: stackConfig.edtc.enabled ? stackConfig.edtc.base_url : undefined,
  [STACK_COMPONENTS_MAP.qrg]: stackConfig.qrg.enabled ? stackConfig.qrg.base_url : undefined,
  [STACK_COMPONENTS_MAP.gcs]: stackConfig.gcs.enabled ? stackConfig.gcs.base_url : undefined,
  [STACK_COMPONENTS_MAP.dcs]: stackConfig.dcs.enabled ? stackConfig.dcs.base_url : undefined,
}

const tts = new TTS({
  authorization: {
    mode: AUTHORIZATION_MODES.KEY,
    key: token,
  },
  stackConfig: stack,
  connectionType: 'http',
  proxy: false,
  axiosConfig: {
    timeout: 10000,
  },
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
