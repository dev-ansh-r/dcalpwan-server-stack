

import StackConfiguration from '../util/stack-configuration'

import Api from '.'

jest.mock('../../generated/api-definition.json', () => ({
  ApplicationRegistry: {
    Create: {
      file: 'lorawan-stack/api/application_services.proto',
      http: [
        {
          method: 'post',
          pattern: '/users/{collaborator.user_ids.user_id}/applications',
          body: '*',
          parameters: ['collaborator.user_ids.user_id'],
        },
        {
          method: 'post',
          pattern: '/organizations/{collaborator.organization_ids.organization_id}/applications',
          body: '*',
          parameters: ['collaborator.organization_ids.organization_id'],
        },
      ],
    },
    List: {
      file: 'lorawan-stack/api/application_services.proto',
      http: [
        {
          method: 'get',
          pattern: '/applications',
          parameters: [],
        },
      ],
    },
    Events: {
      file: 'lorawan-stack/api/application_services.proto',
      http: [
        {
          method: 'get',
          pattern: '/events',
          parameters: [],
          stream: true,
        },
      ],
    },
  },
}))

describe('API class', () => {
  let api
  beforeEach(() => {
    api = new Api(
      'http',
      { mode: 'key', key: 'ABCDEFG' },
      new StackConfiguration({
        is: 'http://localhost:1885',
        as: 'http://localhost:1885',
        ns: 'http://localhost:1885',
        js: 'http://localhost:1885',
      }),
    )
    api._connector.handleRequest = jest.fn()
  })

  it('applies api definitions correctly', () => {
    expect(api.ApplicationRegistry.Create).toBeDefined()
    expect(typeof api.ApplicationRegistry.Create).toBe('function')
  })

  it('applies parameters correctly', () => {
    api.ApplicationRegistry.Create(
      { routeParams: { 'collaborator.user_ids.user_id': 'test' } },
      { name: 'test-name' },
    )

    expect(api._connector.handleRequest).toHaveBeenCalledTimes(1)
    expect(api._connector.handleRequest).toHaveBeenCalledWith(
      'post',
      '/users/test/applications',
      undefined,
      { name: 'test-name' },
      false,
    )
  })

  it('throws when parameters mismatch', () => {
    expect(() => {
      api.ApplicationRegistry.Create({ 'some.other.param': 'test' })
    }).toThrow()
  })

  it('respects the search query', () => {
    api.ApplicationRegistry.List(undefined, { limit: 2, page: 1 })

    expect(api._connector.handleRequest).toHaveBeenCalledTimes(1)
    expect(api._connector.handleRequest).toHaveBeenCalledWith(
      'get',
      '/applications',
      undefined,
      { limit: 2, page: 1 },
      false,
    )
  })

  it('sets stream value to true for streaming endpoints', () => {
    api.ApplicationRegistry.Events()

    expect(api._connector.handleRequest).toHaveBeenCalledTimes(1)
    expect(api._connector.handleRequest).toHaveBeenCalledWith(
      'get',
      '/events',
      undefined,
      undefined,
      true,
    )
  })
})
