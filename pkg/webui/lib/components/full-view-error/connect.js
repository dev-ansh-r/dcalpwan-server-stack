

import { connect } from 'react-redux'

import { selectOnlineStatus } from '@ttn-lw/lib/store/selectors/status'

export default ErrorView =>
  connect(state => ({
    onlineStatus: selectOnlineStatus(state),
  }))(ErrorView)
