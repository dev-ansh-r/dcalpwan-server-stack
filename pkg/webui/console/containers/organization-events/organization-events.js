

import React from 'react'

import Events from '@console/components/events'

import PropTypes from '@ttn-lw/lib/prop-types'

const OrganizationEvents = props => {
  const { orgId, events, widget, paused, onPauseToggle, onClear, truncated } = props

  if (widget) {
    return (
      <Events.Widget
        events={events}
        toAllUrl={`/organizations/${orgId}/data`}
        entityId={orgId}
        scoped
      />
    )
  }

  return (
    <Events
      events={events}
      paused={paused}
      onClear={onClear}
      onPauseToggle={onPauseToggle}
      truncated={truncated}
      entityId={orgId}
      disableFiltering
    />
  )
}

OrganizationEvents.propTypes = {
  events: PropTypes.events,
  onClear: PropTypes.func.isRequired,
  onPauseToggle: PropTypes.func.isRequired,
  orgId: PropTypes.string.isRequired,
  paused: PropTypes.bool.isRequired,
  truncated: PropTypes.bool.isRequired,
  widget: PropTypes.bool,
}

OrganizationEvents.defaultProps = {
  widget: false,
  events: [],
}

export default OrganizationEvents
