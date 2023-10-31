

import React from 'react'

import MapWidget from '@ttn-lw/components/map/widget'

import PropTypes from '@ttn-lw/lib/prop-types'

import locationToMarkers from '@console/lib/location-to-markers'

const GatewayMap = ({ gateway }) => {
  const { gateway_id } = gateway.ids

  const markers = gateway.antennas
    ? locationToMarkers(gateway.antennas.map(antenna => antenna.location))
    : []

  return (
    <MapWidget
      id="gateway-map-widget"
      markers={markers}
      path={`/gateways/${gateway_id}/location`}
    />
  )
}

GatewayMap.propTypes = {
  gateway: PropTypes.gateway.isRequired,
}

export default GatewayMap
