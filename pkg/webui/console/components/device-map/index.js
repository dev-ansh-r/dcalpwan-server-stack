

import React from 'react'

import MapWidget from '@ttn-lw/components/map/widget'

import PropTypes from '@ttn-lw/lib/prop-types'

import locationToMarkers from '@console/lib/location-to-markers'

const DeviceMap = ({ device }) => {
  const { device_id } = device.ids
  const { application_id } = device.ids.application_ids

  const markers = locationToMarkers(device.locations)

  return (
    <MapWidget
      id="device-map-widget"
      markers={markers}
      path={`/applications/${application_id}/devices/${device_id}/location`}
    />
  )
}

DeviceMap.propTypes = {
  device: PropTypes.device.isRequired,
}

export default DeviceMap
