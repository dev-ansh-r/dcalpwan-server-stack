
import {
  createPaginationIdsSelectorByEntity,
  createPaginationTotalCountSelectorByEntity,
} from '@ttn-lw/lib/store/selectors/pagination'
import { createFetchingSelector } from '@ttn-lw/lib/store/selectors/fetching'
import { createErrorSelector } from '@ttn-lw/lib/store/selectors/error'
import { combinePacketBrokerIds } from '@ttn-lw/lib/selectors/id'

import { GET_PACKET_BROKER_INFO_BASE } from '@console/store/actions/packet-broker'

const ENTITY = 'packetBrokerNetworks'

const selectPacketBrokerStore = state => state.packetBroker

// General.
export const selectInfo = state => selectPacketBrokerStore(state).info
export const selectRegistration = state => selectInfo(state).registration || {}
export const selectPacketBrokerOwnCombinedId = state =>
  combinePacketBrokerIds(selectRegistration(state).id)
export const selectInfoFetching = createFetchingSelector(GET_PACKET_BROKER_INFO_BASE)
export const selectInfoError = createErrorSelector(GET_PACKET_BROKER_INFO_BASE)

export const selectRegistered = state => selectPacketBrokerStore(state).registered
export const selectRegisterEnabled = state => selectPacketBrokerStore(state).registerEnabled
export const selectEnabled = state => selectPacketBrokerStore(state).enabled
export const selectListed = state => selectPacketBrokerStore(state).listed

export const selectHomeNetworkDefaultRoutingPolicy = state =>
  selectPacketBrokerStore(state).defaultHomeNetworkRoutingPolicy

export const selectHomeNetworkDefaultGatewayVisibility = state =>
  selectPacketBrokerStore(state).defaultHomeNetworkGatewayVisibility

// Network.
export const selectPacketBrokerNetworkEntitiesStore = state =>
  selectPacketBrokerStore(state).networks.entities
export const selectPacketBrokerNetworkById = (state, combinedId) =>
  selectPacketBrokerNetworkEntitiesStore(state)[combinedId]

// Networks.
const selectPBNetworksIds = createPaginationIdsSelectorByEntity(ENTITY)
const selectPBNetworksTotalCount = createPaginationTotalCountSelectorByEntity(ENTITY)

export const selectPacketBrokerNetworks = state =>
  selectPBNetworksIds(state).map(netId => selectPacketBrokerNetworkById(state, netId))
export const selectPacketBrokerNetworksTotalCount = state => selectPBNetworksTotalCount(state)

// Policies.
export const selectPacketBrokerPoliciesStore = state => selectPacketBrokerStore(state).policies
export const selectPacketBrokerForwarderPoliciesStore = state =>
  selectPacketBrokerPoliciesStore(state).forwarders
export const selectPacketBrokerForwarderPolicyById = (state, combinedId) =>
  selectPacketBrokerForwarderPoliciesStore(state)[combinedId]
export const selectPacketBrokerHomeNetworkPoliciesStore = state =>
  selectPacketBrokerPoliciesStore(state).homeNetworks
export const selectPacketBrokerHomeNetworkPolicyById = (state, combinedId) =>
  selectPacketBrokerHomeNetworkPoliciesStore(state)[combinedId]
