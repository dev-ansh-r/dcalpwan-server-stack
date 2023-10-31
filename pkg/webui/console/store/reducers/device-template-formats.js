

import { GET_DEVICE_TEMPLATE_FORMATS_SUCCESS } from '@console/store/actions/device-template-formats'

const defaultState = {
  formats: undefined,
}

const deviceTemplateFormats = (state = defaultState, { type, payload }) => {
  switch (type) {
    case GET_DEVICE_TEMPLATE_FORMATS_SUCCESS:
      return {
        ...state,
        formats: payload,
      }
    default:
      return state
  }
}

export default deviceTemplateFormats
