

import { createFetchingSelector } from '@ttn-lw/lib/store/selectors/fetching'
import { createErrorSelector } from '@ttn-lw/lib/store/selectors/error'

import { GET_DEVICE_TEMPLATE_FORMATS_BASE } from '@console/store/actions/device-template-formats'

const selectDeviceTemplateFormatsStore = state => state.deviceTemplateFormats

export const selectDeviceTemplateFormats = state => {
  const store = selectDeviceTemplateFormatsStore(state)

  return store.formats || {}
}

export const selectDeviceTemplateFormatsError = createErrorSelector(
  GET_DEVICE_TEMPLATE_FORMATS_BASE,
)
export const selectDeviceTemplateFormatsFetching = createFetchingSelector(
  GET_DEVICE_TEMPLATE_FORMATS_BASE,
)
