

import { connect } from 'react-redux'

import { getOrganizationsList } from '@console/store/actions/organizations'

import {
  selectOrganizationsFetching,
  selectOrganizationsError,
  selectOrganizations,
} from '@console/store/selectors/organizations'
import { selectUser } from '@console/store/selectors/logout'

export default OwnersSelect =>
  connect(
    state => ({
      user: selectUser(state),
      organizations: selectOrganizations(state),
      error: selectOrganizationsError(state),
      fetching: selectOrganizationsFetching(state),
    }),
    { getOrganizationsList },
  )(OwnersSelect)
