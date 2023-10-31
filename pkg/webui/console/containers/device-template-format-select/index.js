

import { defineMessages } from 'react-intl'

import CreateFetchSelect from '@console/containers/fetch-select'

import { getDeviceTemplateFormats } from '@console/store/actions/device-template-formats'

import {
  selectDeviceTemplateFormats,
  selectDeviceTemplateFormatsError,
  selectDeviceTemplateFormatsFetching,
} from '@console/store/selectors/device-template-formats'

const m = defineMessages({
  title: 'File format',
  warning: 'End device template formats unavailable',
})

const formatOptions = formats =>
  Object.keys(formats).map(key => ({
    value: key,
    label: formats[key].name,
    description: formats[key].description,
    fileExtensions: formats[key].file_extensions,
  }))

export default CreateFetchSelect({
  fetchOptions: getDeviceTemplateFormats,
  optionsSelector: selectDeviceTemplateFormats,
  errorSelector: selectDeviceTemplateFormatsError,
  fetchingSelector: selectDeviceTemplateFormatsFetching,
  defaultWarning: m.warning,
  defaultTitle: m.title,
  optionsFormatter: formatOptions,
})
