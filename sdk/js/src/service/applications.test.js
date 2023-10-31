

const mockApplicationData = {
  ids: {
    application_id: 'test',
  },
  created_at: '2018-08-29T14:00:20.793Z',
  updated_at: '2018-08-29T14:00:20.793Z',
  name: 'string',
  description: 'string',
  attributes: {
    additionalProp1: 'string',
    additionalProp2: 'string',
    additionalProp3: 'string',
  },
  contact_info: [
    {
      contact_type: 'CONTACT_TYPE_OTHER',
      contact_method: 'CONTACT_METHOD_OTHER',
      value: 'string',
      public: true,
      validated_at: '2018-08-29T14:00:20.793Z',
    },
  ],
}

jest.mock('../api', () =>
  jest.fn().mockImplementation(() => ({
    EndDeviceRegistry: {
      UpdateAllowedFieldMaskPaths: [],
    },
    NsEndDeviceRegistry: {
      SetAllowedFieldMaskPaths: [],
    },
    JsEndDeviceRegistry: {
      SetAllowedFieldMaskPaths: [],
    },
    AsEndDeviceRegistry: {
      SetAllowedFieldMaskPaths: [],
    },
    ApplicationRegistry: {
      Get: jest.fn().mockResolvedValue({ data: mockApplicationData }),
      List: jest.fn().mockResolvedValue({
        data: { applications: [mockApplicationData] },
        headers: { 'x-total-count': 1 },
      }),
    },
  })),
)

describe('Applications', () => {
  let applications
  beforeEach(() => {
    const Api = require('../api')

    const Applications = require('./applications').default
    applications = new Applications(new Api(), { defaultUserId: 'testuser' })
  })

  describe('when using proxied results', () => {
    it('initializes correctly', () => {
      jest.resetModules()

      expect(applications._api).toBeDefined()
    })

    it('returns an application instance on getById()', async () => {
      jest.resetModules()

      const app = await applications.getById('test')
      expect(app).toBeDefined()
      expect(app.ids.application_id).toBe('test')
    })

    it('returns an application list on getAll()', async () => {
      jest.resetModules()

      const result = await applications.getAll()
      expect(result).toBeDefined()

      const { applications: apps, totalCount } = result
      expect(apps.constructor.name).toBe('Array')
      expect(apps).toHaveLength(1)
      expect(totalCount).toBe(1)
    })
  })
})
