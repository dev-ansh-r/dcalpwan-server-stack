

import { hexToBase64, base64ToHex } from './bytes'

describe('Bytes utils', () => {
  const base64 = [
    'AQ==',
    'Ag==',
    '/w==',
    'vyUs3Q==',
    'MEijVA==',
    'EjRWeJCrze8=',
    'CZxkIN2eINE=',
    'AQIDBA==',
  ]
  const hex = [
    '01',
    '02',
    'ff',
    'bf252cdd',
    '3048a354',
    '1234567890abcdef',
    '099c6420dd9e20d1',
    '01 02 03 04 ',
  ]

  describe('when using base64ToHex', () => {
    const testTable = base64.map((value, index) => [value, hex[index]])
    it.each(testTable)('yields base64ToHex(%s) = %s', (base64Str, expectedHex) => {
      expect(base64ToHex(base64Str)).toBe(expectedHex.replace(/ /g, ''))
    })
  })

  describe('when using hexToBase64', () => {
    const testTable = hex.map((value, index) => [value, base64[index]])
    it.each(testTable)('yields hexToBase64(%s) = %s', (hexStr, expectedBase64) => {
      expect(hexToBase64(hexStr)).toBe(expectedBase64)
    })
  })
})
