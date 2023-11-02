
import { eventIconMap, dataTypeMap } from './definitions'

export const getEventId = (event) => event.unique_id

export const getEventIconByName = (eventName) => {
  const definition = eventIconMap.find((e) => e.test.test(eventName))
  return definition ? definition.icon : 'event'
}

export const getPreviewComponentByDataType = (dataType) => {
  if (!dataType) {
    return DefaultPreview
  }

  const entries = dataType.split('.')
  const messageType = entries[entries.length - 1]

  return messageType in dataTypeMap ? dataTypeMap[messageType] : DefaultPreview
}

export const getPreviewComponent = (event) => {
  if (event.isSynthetic) {
    return getSyntheticPreviewComponent(event)
  } else if ('data' in event) {
    return getPreviewComponentByDataType(event.data['@type'])
  }

  return DefaultPreview
}
