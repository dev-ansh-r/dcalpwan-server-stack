

import DownlinkMessage from '../previews/downlink-message'
import GatewayUplinkMessage from '../previews/gateway-uplink-message'
import ApplicationDownlink from '../previews/application-downlink'
import ApplicationUplink from '../previews/application-uplink'
import ApplicationUplinkNormalized from '../previews/application-uplink-normalized'
import ApplicationUp from '../previews/application-up'
import ApplicationLocation from '../previews/application-location'
import JoinRequest from '../previews/join-request'
import JoinResponse from '../previews/join-response'
import GatewayStatus from '../previews/gateway-status'
import ErrorDetails from '../previews/error-details'
import Value from '../previews/value'

export const eventIconMap = [
  {
    test: /^(ns\.|as\.|js\.)?([a-z0-9](?:[-_]?[a-z0-9]){2,}\.)+create$/,
    icon: 'event_create',
  },
  {
    test: /^(ns\.|as\.|js\.)?([a-z0-9](?:[-_]?[a-z0-9]){2,}\.)+update$/,
    icon: 'event_update',
  },
  {
    test: /^(ns\.|as\.|js\.)?([a-z0-9](?:[-_]?[a-z0-9]){2,}\.)+delete$/,
    icon: 'event_delete',
  },
  {
    test: /^(ns|as)\.up(\.[a-z0-9](?:[-_]?[a-z0-9]){2,})+$/,
    icon: 'event_uplink',
  },
  {
    test: /^(ns|as)\.down(\.[a-z0-9](?:[-_]?[a-z0-9]){2,})+$/,
    icon: 'event_downlink',
  },
  {
    test: /^(js|ns|as)(\.up|\.down)?\.(join|rejoin)(\.[a-z0-9](?:[-_]?[a-z0-9]){2,})+$/,
    icon: 'event_join',
  },
  {
    test: /^gs\.up(\.[a-z0-9](?:[-_]?[a-z0-9]){2,})+$/,
    icon: 'event_uplink',
  },
  {
    test: /^gs\.down(\.[a-z0-9](?:[-_]?[a-z0-9]){2,})+$/,
    icon: 'event_downlink',
  },
  {
    test: /^gs.gateway.connect$/,
    icon: 'event_gateway_connect',
  },
  {
    test: /^gs.gateway.disconnect$/,
    icon: 'event_gateway_disconnect',
  },
  {
    test: /^ns\.mac\.rekey\..*$/,
    icon: 'event_rekey',
  },
  {
    test: /^ns\.mac\.device_mode\..*$/,
    icon: 'event_mode',
  },
  {
    test: /^ns\..*\.switch\..*$/,
    icon: 'event_switch',
  },
  {
    test: /^ns(\.[a-z0-9](?:[-_]?[a-z0-9]){2,})+$/,
    icon: 'event_connection',
  },
  {
    test: /^gs\.status\..*$/,
    icon: 'event_status',
  },
  {
    test: /^synthetic\.status\.cleared$/,
    icon: 'event_clear_all',
  },
  {
    test: /^synthetic\.error\..*$/,
    icon: 'event_error',
  },
]

export const dataTypeMap = {
  ApplicationDownlink,
  ApplicationUplink,
  ApplicationUplinkNormalized,
  ApplicationUp,
  ApplicationLocation,
  DownlinkMessage,
  GatewayUplinkMessage,
  JoinRequest,
  JoinResponse,
  ErrorDetails,
  GatewayStatus,
  Value,
}

export const applicationUpMessages = [
  'uplink_message',
  'uplink_normalized',
  'join_accept',
  'downlink_ack',
  'downlink_nack',
  'downlink_sent',
  'downlink_failed',
  'downlink_queued',
  'downlink_queue_invalidated',
  'location_solved',
  'service_data',
]
