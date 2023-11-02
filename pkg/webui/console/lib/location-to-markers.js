
import React from 'react'
import { isArray, isEmpty, isPlainObject } from 'lodash'
import { Tooltip } from 'react-leaflet'

import Message from '@ttn-lw/lib/components/message'

import sharedMessages from '@ttn-lw/lib/shared-messages'

const sourceMessages = {
  SOURCE_GPS: sharedMessages.locationSourceGps,
  SOURCE_REGISTRY: sharedMessages.locationSourceRegistry,
  SOURCE_UNKNOWN: sharedMessages.locationSourceUnknown,
  SOURCE_IP_GEOLOCATION: sharedMessages.locationSourceIpGeolocation,
  SOURCE_WIFI_RSSI_GEOLOCATION: sharedMessages.locationSourceWifiRssi,
  SOURCE_BT_RSSI_GEOLOCATION: sharedMessages.locationSourceBtRssi,
  SOURCE_LORA_RSSI_GEOLOCATION: sharedMessages.locationSourceLoraRssi,
  SOURCE_LORA_TDOA_GEOLOCATION: sharedMessages.locationSourceLoraTdoa,
  SOURCE_COMBINED_GEOLOCATION: sharedMessages.locationSourceCombined,
}

const createLocationObject = (location, key) => ({
  position: {
    latitude: location.latitude || 0,
    longitude: location.longitude || 0,
  },
  accuracy: location.accuracy,
  children: (
    <Tooltip direction="top" offset={[-15, -10]} opacity={1}>
      <Message
        component="strong"
        content={location?.source ? sourceMessages[location.source] : sourceMessages.SOURCE_UNKNOWN}
      />
      <br />
      <Message
        content={
          key === 'user' || location?.source === 'SOURCE_REGISTRY'
            ? sharedMessages.locationMarkerDescriptionUser
            : location.trusted === false
            ? sharedMessages.locationMarkerDescriptionUntrusted
            : sharedMessages.locationMarkerDescriptionNonUser
        }
      />
      <br />
      Long: {location.longitude} / Lat: {location.latitude}
    </Tooltip>
  ),
})

export default locations => {
  if (isPlainObject(locations)) {
    return Object.keys(locations).map(key => createLocationObject(locations[key], key))
  }

  if (isArray(locations)) {
    return locations
      .filter(l => Boolean(l) && isPlainObject(l) && !isEmpty(l))
      .map(location => createLocationObject(location))
  }

  return []
}
