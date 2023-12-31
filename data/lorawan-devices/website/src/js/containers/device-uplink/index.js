
import React from 'react'
import PropTypes from 'prop-types'

import { generateUplinkEvents } from '../../components/events/utils/generate-events'
import Events from '../../components/events'

function DeviceUplink(props) {
  const {
    truncated,
    brand_id,
    model_id,
    device,
  } = props
  const [events, setEvents] = React.useState([])
  const [paused, setPaused] = React.useState(false)
  const [displayEvents, setDisplayEvents] = React.useState(false)
  const [payloadFormatter, setPayloadFormatter] = React.useState({})

  React.useEffect(() => {
    let firm
    let band
    device.firmwareversions.forEach((fw) => {
      for (const profile in fw.profiles) {
        if ('codec' in fw.profiles[profile]) {
          firm = fw.version
          band = profile
        }
      }
    })
    if (Boolean(firm) && Boolean(band)) {
      if (device.codecs[band.toLowerCase()]) {
        setPayloadFormatter(device.codecs[band.toLowerCase()].uplinkdecoder)
      }
    }
  }, [brand_id, model_id])

  React.useEffect(() => {
    if (Boolean(payloadFormatter) && 'examples' in payloadFormatter) {
      setDisplayEvents(true)
    }
  }, [payloadFormatter])
  React.useEffect(() => {
    const interval = setInterval(() => {
      if (!paused && events.length < 10 && displayEvents) {
        const generatedEvents = generateUplinkEvents(payloadFormatter)
        setEvents([...events, ...generatedEvents])
      }
    }, 5000)
    return () => clearInterval(interval)
  }, [displayEvents, events, paused, payloadFormatter])
  const onPauseToggle = () => {
    setPaused(!paused)
  }
  const onClear = () => {
    setEvents([])
  }
  return (
    displayEvents && (
      <Events
        events={events}
        paused={paused}
        onClear={onClear}
        onPauseToggle={onPauseToggle}
        truncated={truncated}
        scoped
        widget
      />
    )
  )
}

DeviceUplink.propTypes = {
  brand_id: PropTypes.string.isRequired,
  model_id: PropTypes.string.isRequired,
  truncated: PropTypes.bool,
}

DeviceUplink.defaultProps = {
  truncated: false,
}

export default DeviceUplink
