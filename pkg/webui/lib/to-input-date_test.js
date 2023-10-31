

import toInputDate from './to-input-date'

describe('To input date converter', () => {
  describe('when using a correct argument', () => {
    const date = new Date('2020-09-24T12:00:00Z')

    it('returns a string length equal to 10', () => {
      expect(toInputDate(date)).toHaveLength(10)
    })

    it('returns a string of yyyy-mm--dd format', () => {
      expect(toInputDate(date)).toMatch(/([12]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01]))/)
    })
  })

  describe('when unsing an incorrect argument', () => {
    const string = 'ABC123'
    const date = new Date('ABC123')

    it('returns undefined when using a random string', () => {
      expect(toInputDate(string)).toBe(undefined)
    })

    it('returns undefined when using an invalid date', () => {
      expect(toInputDate(date)).toBe(undefined)
    })
  })
})
