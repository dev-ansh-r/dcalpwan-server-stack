

import { connect } from 'react-redux'

import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'

import { updateGateway } from '@console/store/actions/gateways'

import { selectSelectedGateway, selectSelectedGatewayId } from '@console/store/selectors/gateways'

const mapStateToProps = state => ({
  gateway: selectSelectedGateway(state),
  gatewayId: selectSelectedGatewayId(state),
})

const mapDispatchToProps = {
  updateGateway: attachPromise(updateGateway),
}

export default Component => connect(mapStateToProps, mapDispatchToProps)(Component)
