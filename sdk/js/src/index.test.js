

import Applications from './service/applications'

import TTS from '.'

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

const mockDeviceData = {
  ids: {
    device_id: 'test-device',
    application_ids: {
      application_id: 'test',
    },
    dev_eui: 'string',
    join_eui: 'string',
    dev_addr: 'string',
  },
}

jest.mock('./api', () =>
  jest.fn().mockImplementation(() => ({
    ApplicationRegistry: {
      Get: jest.fn().mockResolvedValue({ data: mockApplicationData }),
      List: jest.fn().mockResolvedValue({
        data: { applications: [mockApplicationData] },
        headers: { 'x-total-count': 1 },
      }),
    },
    EndDeviceRegistry: {
      Get: jest.fn().mockResolvedValue({ data: mockDeviceData }),
      UpdateAllowedFieldMaskPaths: [],
    },
    NsEndDeviceRegistry: {
      Get: jest.fn().mockResolvedValue({ data: mockDeviceData }),
      SetAllowedFieldMaskPaths: [],
    },
    AsEndDeviceRegistry: {
      Get: jest.fn().mockResolvedValue({ data: mockDeviceData }),
      SetAllowedFieldMaskPaths: [],
    },
    JsEndDeviceRegistry: {
      Get: jest.fn().mockResolvedValue({ data: mockDeviceData }),
      SetAllowedFieldMaskPaths: [],
    },
  })),
)

describe('SDK class', () => {
  const token = 'faketoken'
  const tts = new TTS({
    authorization: {
      mode: 'key',
      key: token,
    },
    connectionType: 'http',
    stackConfig: { is: 'http://localhost:1885/api/v3' },
  })

  it('instanciates successfully', async () => {
    expect(tts).toBeDefined()
    expect(tts).toBeInstanceOf(TTS)
    expect(tts.Applications).toBeInstanceOf(Applications)
  })

  it('retrieves application instance correctly', async () => {
    const app = await tts.Applications.getById('test')
    expect(app).toBeDefined()
  })
})
