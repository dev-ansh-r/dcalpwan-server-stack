
import tts from '@console/api/tts'

import { isNotFoundError } from '@ttn-lw/lib/errors/utils'
import createRequestLogic from '@ttn-lw/lib/store/logics/create-request-logic'

import * as repository from '@console/store/actions/device-repository'

const ignoreNotFound = async (action, appId, version_ids) => {
  try {
    return await action(appId, version_ids)
  } catch (e) {
    if (!isNotFoundError(e)) {
      throw e
    }
  }

  return undefined
}

const listDeviceBrandsLogic = createRequestLogic({
  type: repository.LIST_BRANDS,
  process: async ({ action }) => {
    const {
      payload: { appId, params = {} },
      meta: { selector = [] },
    } = action

    return tts.Applications.Devices.Repository.listBrands(appId, params, selector)
  },
})

const getDeviceBrandLogic = createRequestLogic({
  type: repository.GET_BRAND,
  process: ({ action }) => {
    const {
      payload: { appId, brandId },
      meta: { selector = [] },
    } = action

    return tts.Applications.Devices.Repository.getBrand(appId, brandId, selector)
  },
})

const listDeviceModelsLogic = createRequestLogic({
  type: repository.LIST_MODELS,
  process: async ({ action }) => {
    const {
      payload: { appId, brandId, params = {} },
      meta: { selector = [] },
    } = action

    const { models, totalCount } = await tts.Applications.Devices.Repository.listModels(
      appId,
      brandId,
      params,
      selector,
    )

    return { models, totalCount, brandId }
  },
})

const getDeviceModelLogic = createRequestLogic({
  type: repository.GET_MODEL,
  process: ({ action }) => {
    const {
      payload: { appId, brandId, modelId },
      meta: { selector = [] },
    } = action

    return tts.Applications.Devices.Repository.getModel(appId, brandId, modelId, selector)
  },
})

const getTemplateLogic = createRequestLogic({
  type: repository.GET_TEMPLATE,
  process: ({ action }) => {
    const {
      payload: { appId, version },
    } = action

    return tts.Applications.Devices.Repository.getTemplate(appId, version)
  },
})

const getRepositoryPayloadFormattersLogic = createRequestLogic({
  type: repository.GET_REPO_PF,
  process: async ({ action }) => {
    const {
      payload: { appId, version_ids },
    } = action

    const repositoryPayloadFormatters = await Promise.all([
      ignoreNotFound(tts.Applications.Devices.Repository.getUplinkDecoder, appId, version_ids),
      ignoreNotFound(tts.Applications.Devices.Repository.getDownlinkDecoder, appId, version_ids),
      ignoreNotFound(tts.Applications.Devices.Repository.getDownlinkEncoder, appId, version_ids),
    ])

    return {
      ...repositoryPayloadFormatters[0],
      ...repositoryPayloadFormatters[1],
      ...repositoryPayloadFormatters[2],
    }
  },
})

export default [
  listDeviceBrandsLogic,
  getDeviceBrandLogic,
  listDeviceModelsLogic,
  getDeviceModelLogic,
  getTemplateLogic,
  getRepositoryPayloadFormattersLogic,
]
