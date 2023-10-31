

import getCookieValue from './cookie'

describe('Cookie utils', () => {
  describe('when `document.cookie` is empty', () => {
    it('returns `undefined`', () => {
      const value = getCookieValue('missingKey')

      expect(value).toBeUndefined()
    })
  })

  describe('when `document.cookie` has a single entry', () => {
    const key = 'testKey'
    const value = 'testValue'

    beforeEach(() => {
      document.cookie = `${key}=${value}`
    })

    afterEach(() => {
      document.cookie = `${key}=; expires=Thu, 01 Jan 1970 00:00:00 GMT`
    })

    it('extracts value for existing key', () => {
      expect(getCookieValue(key)).toBe(value)
    })

    it('returns `undefined` for non existing key', () => {
      expect(getCookieValue('nonExistingKey')).toBeUndefined()
    })
  })

  describe('when `document.cookie` has multiple entries', () => {
    const key1 = 'testKey1'
    const key2 = 'testKey2'
    const value1 = 'testValue1'
    const value2 = 'testValue2'

    beforeEach(() => {
      document.cookie = `${key1}=${value1}`
      document.cookie = `${key2}=${value2}`
    })

    afterEach(() => {
      document.cookie = `${key1}=; expires=Thu, 01 Jan 1970 00:00:00 GMT`
      document.cookie = `${key2}=; expires=Thu, 01 Jan 1970 00:00:00 GMT`
    })

    it('extracts the value of the first entry', () => {
      expect(getCookieValue(key1)).toBe(value1)
    })

    it('extracts the value of the second entry', () => {
      expect(getCookieValue(key2)).toBe(value2)
    })
  })
})
