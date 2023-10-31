

import React from 'react'
import { connect } from 'react-redux'

import Events from '@console/components/events'

import withFeatureRequirement from '@console/lib/components/with-feature-requirement'

import PropTypes from '@ttn-lw/lib/prop-types'

import { mayViewApplicationEvents } from '@console/lib/feature-checks'

import {
  clearApplicationEventsStream,
  pauseApplicationEventsStream,
  resumeApplicationEventsStream,
  setApplicationEventsFilter,
} from '@console/store/actions/applications'

import {
  selectApplicationEvents,
  selectApplicationEventsPaused,
  selectApplicationEventsTruncated,
  selectApplicationEventsFilter,
} from '@console/store/selectors/applications'

const ApplicationEvents = props => {
  const {
    appId,
    events,
    widget,
    paused,
    onClear,
    onPauseToggle,
    truncated,
    onFilterChange,
    filter,
  } = props

  if (widget) {
    return (
      <Events.Widget entityId={appId} events={events} toAllUrl={`/applications/${appId}/data`} />
    )
  }

  return (
    <Events
      entityId={appId}
      events={events}
      paused={paused}
      onClear={onClear}
      truncated={truncated}
      filter={filter}
      onPauseToggle={onPauseToggle}
      onFilterChange={onFilterChange}
    />
  )
}

ApplicationEvents.propTypes = {
  appId: PropTypes.string.isRequired,
  events: PropTypes.events,
  filter: PropTypes.eventFilter,
  onClear: PropTypes.func.isRequired,
  onFilterChange: PropTypes.func.isRequired,
  onPauseToggle: PropTypes.func.isRequired,
  paused: PropTypes.bool.isRequired,
  truncated: PropTypes.bool.isRequired,
  widget: PropTypes.bool,
}

ApplicationEvents.defaultProps = {
  widget: false,
  events: [],
  filter: undefined,
}

export default withFeatureRequirement(mayViewApplicationEvents)(
  connect(
    (state, props) => {
      const { appId } = props

      return {
        events: selectApplicationEvents(state, appId),
        paused: selectApplicationEventsPaused(state, appId),
        truncated: selectApplicationEventsTruncated(state, appId),
        filter: selectApplicationEventsFilter(state, appId),
      }
    },
    (dispatch, ownProps) => ({
      onClear: () => dispatch(clearApplicationEventsStream(ownProps.appId)),
      onPauseToggle: paused =>
        paused
          ? dispatch(resumeApplicationEventsStream(ownProps.appId))
          : dispatch(pauseApplicationEventsStream(ownProps.appId)),
      onFilterChange: filterId => dispatch(setApplicationEventsFilter(ownProps.appId, filterId)),
    }),
  )(ApplicationEvents),
)
