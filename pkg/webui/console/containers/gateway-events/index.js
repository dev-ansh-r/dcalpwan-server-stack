

import React from 'react'
import { connect } from 'react-redux'

import Events from '@console/components/events'

import withFeatureRequirement from '@console/lib/components/with-feature-requirement'

import PropTypes from '@ttn-lw/lib/prop-types'

import { mayViewGatewayEvents } from '@console/lib/feature-checks'

import {
  clearGatewayEventsStream,
  pauseGatewayEventsStream,
  resumeGatewayEventsStream,
  setGatewayEventsFilter,
} from '@console/store/actions/gateways'

import {
  selectGatewayEvents,
  selectGatewayEventsPaused,
  selectGatewayEventsTruncated,
  selectGatewayEventsFilter,
} from '@console/store/selectors/gateways'

const GatewayEvents = props => {
  const {
    gtwId,
    events,
    widget,
    paused,
    onPauseToggle,
    onClear,
    onFilterChange,
    truncated,
    filter,
  } = props

  if (widget) {
    return (
      <Events.Widget events={events} entityId={gtwId} toAllUrl={`/gateways/${gtwId}/data`} scoped />
    )
  }

  return (
    <Events
      events={events}
      entityId={gtwId}
      paused={paused}
      onClear={onClear}
      onPauseToggle={onPauseToggle}
      onFilterChange={onFilterChange}
      truncated={truncated}
      filter={filter}
      scoped
    />
  )
}

GatewayEvents.propTypes = {
  events: PropTypes.events,
  filter: PropTypes.eventFilter,
  gtwId: PropTypes.string.isRequired,
  onClear: PropTypes.func.isRequired,
  onFilterChange: PropTypes.func.isRequired,
  onPauseToggle: PropTypes.func.isRequired,
  paused: PropTypes.bool.isRequired,
  truncated: PropTypes.bool.isRequired,
  widget: PropTypes.bool,
}

GatewayEvents.defaultProps = {
  widget: false,
  events: [],
  filter: undefined,
}

export default withFeatureRequirement(mayViewGatewayEvents)(
  connect(
    (state, props) => {
      const { gtwId } = props

      return {
        events: selectGatewayEvents(state, gtwId),
        paused: selectGatewayEventsPaused(state, gtwId),
        truncated: selectGatewayEventsTruncated(state, gtwId),
        filter: selectGatewayEventsFilter(state, gtwId),
      }
    },
    (dispatch, ownProps) => ({
      onClear: () => dispatch(clearGatewayEventsStream(ownProps.gtwId)),
      onPauseToggle: paused =>
        paused
          ? dispatch(resumeGatewayEventsStream(ownProps.gtwId))
          : dispatch(pauseGatewayEventsStream(ownProps.gtwId)),
      onFilterChange: filterId => dispatch(setGatewayEventsFilter(ownProps.gtwId, filterId)),
    }),
  )(GatewayEvents),
)
