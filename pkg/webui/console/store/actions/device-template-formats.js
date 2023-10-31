

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'
import { createErrorSelector } from '@ttn-lw/lib/store/selectors/error'
import { createFetchingSelector } from '@ttn-lw/lib/store/selectors/fetching'

export const GET_DEVICE_TEMPLATE_FORMATS_BASE = 'GET_DEVICE_TEMPLATE_FORMATS'
export const [
  {
    request: GET_DEVICE_TEMPLATE_FORMATS,
    success: GET_DEVICE_TEMPLATE_FORMATS_SUCCESS,
    failure: GET_DEVICE_TEMPLATE_FORMATS_FAILURE,
  },
  {
    request: getDeviceTemplateFormats,
    success: getDeviceTemplateFormatsSuccess,
    failure: getDeviceTemplateFormatsFailure,
  },
] = createRequestActions(GET_DEVICE_TEMPLATE_FORMATS_BASE)
export const getDeviceTemplateFormatsFetching = createFetchingSelector(
  GET_DEVICE_TEMPLATE_FORMATS_BASE,
)
export const getDeviceTemplateFormatsError = createErrorSelector(GET_DEVICE_TEMPLATE_FORMATS_BASE)

export const CONVERT_TEMPLATE_BASE = 'CONVERT_TEMPLATE'
export const [
  {
    request: CONVERT_TEMPLATE,
    success: CONVERT_TEMPLATE_SUCCESS,
    failure: CONVERT_TEMPLATE_FAILURE,
  },
  { request: convertTemplate, success: convertTemplateSuccess, failure: convertTemplateFailure },
] = createRequestActions(CONVERT_TEMPLATE_BASE, (format_id, data) => ({ format_id, data }))
