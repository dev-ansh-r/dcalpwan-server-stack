

import { parseLorawanMacVersion } from './device-utils'

describe('Parsing LoRaWAN Mac Version', () => {
  it.each([
    ['MAC_V1_0', 100],
    ['MAC_V1_0_1', 101],
    ['MAC_V1_0_2', 102],
    ['MAC_V1_0_3', 103],
    ['MAC_V1_0_4', 104],
    ['MAC_V1_1_0', 110],
    ['100', 0],
    ['101', 0],
    ['102', 0],
    ['103', 0],
    ['104', 0],
    ['110', 0],
    [null, 0],
    [undefined, 0],
    ['invalid', 0],
    ['', 0],
  ])('yields parseLorawanVersion(%p) = %i', (actual, expected) => {
    expect(parseLorawanMacVersion(actual)).toBe(expected)
  })
})
