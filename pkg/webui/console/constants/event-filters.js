
import { APPLICATION, END_DEVICE, GATEWAY } from '@console/constants/entities'

export const END_DEVICE_EVENTS_VERBOSE_FILTERS = [
  'as.*.drop',
  'as.down.data.receive',
  'as.up.*.forward',
  'js.join.accept',
  'js.join.reject',
  'ns.up.data.process',
  'ns.up.join.process',
  'ns.down.data.schedule.attempt',
  'ns.mac.*.answer.reject',
  '*.warning',
  '*.fail',
  'end_device.*',
]

export const APPLICATION_EVENTS_VERBOSE_FILTERS = [
  ...END_DEVICE_EVENTS_VERBOSE_FILTERS,
  'application.*',
]

export const GATEWAY_EVENTS_VERBOSE_FILTERS = [
  'gs.down.send',
  'gs.gateway.connect',
  'gs.gateway.disconnect',
  'gs.status.receive',
  'gs.up.receive',
  '*.warning',
  '*.fail',
  'gateway.*',
]

// Regex for matching heartbeat events that trigger an update of the
// last activity display.
export const EVENT_END_DEVICE_HEARTBEAT_FILTERS_REGEXP = /^as.up\..*\.forward$/

// Utility function to convert filter arrays to Regular Expressions strings
// that the backend accepts for applying filters.
const filterListToRegExpList = array =>
  array.map(f => `/^${f.replace(/\./g, '\\.').replace(/\*/g, '.*')}$/`)

export const EVENT_FILTERS = {
  [APPLICATION]: [
    {
      id: 'default',
      filter: APPLICATION_EVENTS_VERBOSE_FILTERS,
      filterRegExp: filterListToRegExpList(APPLICATION_EVENTS_VERBOSE_FILTERS),
    },
  ],
  [END_DEVICE]: [
    {
      id: 'default',
      filter: END_DEVICE_EVENTS_VERBOSE_FILTERS,
      filterRegExp: filterListToRegExpList(END_DEVICE_EVENTS_VERBOSE_FILTERS),
    },
  ],
  [GATEWAY]: [
    {
      id: 'default',
      filter: GATEWAY_EVENTS_VERBOSE_FILTERS,
      filterRegExp: filterListToRegExpList(GATEWAY_EVENTS_VERBOSE_FILTERS),
    },
  ],
}
