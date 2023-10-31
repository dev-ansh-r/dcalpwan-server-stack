

/* eslint-env jest */
/* eslint-disable arrow-body-style */

import getByPath from './get-by-path'

describe('Get by path', () => {
  it('retrieves leaves as expected when using valid paths', () => {
    const testData = {
      a: {
        b: {
          c: 'foo',
        },
        d: 'bar',
      },
      e: 'baz',
      f: undefined,
      g: null,
      h: [1, 2, 3],
      i: {
        k: [3, 4, 5],
      },
    }

    expect(getByPath(testData, 'a.b.c')).toBe('foo')
    expect(getByPath(testData, 'a.d')).toBe('bar')
    expect(getByPath(testData, 'e')).toBe('baz')
    expect(getByPath(testData, 'f')).toBe(undefined)
    expect(getByPath(testData, 'g')).toBe(null)
    expect(getByPath(testData, 'h')).toMatchObject([1, 2, 3])
    expect(getByPath(testData, 'i.k')).toMatchObject([3, 4, 5])
    expect(getByPath(testData, 'i.e')).toBe(undefined)
    expect(getByPath(testData, 'a.b')).toMatchObject({ c: 'foo' })
  })
})
