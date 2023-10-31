

import { connect } from 'react-redux'

import {
  clearOrganizationEventsStream,
  pauseOrganizationEventsStream,
  resumeOrganizationEventsStream,
} from '@console/store/actions/organizations'

import {
  selectOrganizationEvents,
  selectOrganizationEventsPaused,
  selectOrganizationEventsTruncated,
} from '@console/store/selectors/organizations'

const mapStateToProps = (state, props) => {
  const { orgId } = props

  return {
    events: selectOrganizationEvents(state, orgId),
    paused: selectOrganizationEventsPaused(state, orgId),
    truncated: selectOrganizationEventsTruncated(state, orgId),
  }
}

const mapDispatchToProps = (dispatch, ownProps) => ({
  onClear: () => dispatch(clearOrganizationEventsStream(ownProps.orgId)),
  onPauseToggle: paused =>
    paused
      ? dispatch(resumeOrganizationEventsStream(ownProps.orgId))
      : dispatch(pauseOrganizationEventsStream(ownProps.orgId)),
})

export default Events => connect(mapStateToProps, mapDispatchToProps)(Events)
