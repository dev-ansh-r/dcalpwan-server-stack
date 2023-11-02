
import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

export const LIST_BRANDS_BASE = 'LIST_BRANDS'
export const [
  { request: LIST_BRANDS, success: LIST_BRANDS_SUCCESS, failure: LIST_BRANDS_FAILURE },
  { request: listBrands, success: listBrandsSuccess, failure: listBrandsFailure },
] = createRequestActions(
  LIST_BRANDS_BASE,
  (appId, params = {}) => ({ appId, params }),
  (appId, params, selector) => ({ selector }),
)

export const GET_BRAND_BASE = 'GET_BRAND'
export const [
  { request: GET_BRAND, success: GET_BRAND_SUCCESS, failure: GET_BRAND_FAILURE },
  { request: getBrand, success: getBrandSuccess, failure: getBrandFailure },
] = createRequestActions(
  GET_BRAND_BASE,
  (appId, brandId) => ({ appId, brandId }),
  (appId, brandId, selector) => ({ selector }),
)

export const LIST_MODELS_BASE = 'LIST_MODELS'
export const [
  { request: LIST_MODELS, success: LIST_MODELS_SUCCESS, failure: LIST_MODELS_FAILURE },
  { request: listModels, success: listModelsSuccess, failure: listModelsFailure },
] = createRequestActions(
  LIST_MODELS_BASE,
  (appId, brandId, params) => ({ appId, brandId, params }),
  (appId, brandId, params, selector) => ({ selector }),
)

export const GET_MODEL_BASE = 'GET_MODEL'
export const [
  { request: GET_MODEL, success: GET_MODEL_SUCCESS, failure: GET_MODEL_FAILURE },
  { request: getModel, success: getModelSuccess, failure: getModelFailure },
] = createRequestActions(
  GET_MODEL_BASE,
  (appId, brandId, modelId) => ({ appId, brandId, modelId }),
  (appId, brandId, modelId, selector) => ({ selector }),
)

export const GET_TEMPLATE_BASE = 'GET_DEVICE_TEMPLATE'
export const [
  { request: GET_TEMPLATE, success: GET_TEMPLATE_SUCCESS, failure: GET_TEMPLATE_FAILURE },
  { request: getTemplate, success: getTemplateSuccess, failure: getTemplateFailure },
] = createRequestActions(GET_TEMPLATE_BASE, (appId, version) => ({ appId, version }))

export const GET_REPO_PF_BASE = 'GET_REPOSITORY_PAYLOAD_FORMATTERS'
export const [
  { request: GET_REPO_PF, success: GET_REPO_PF_SUCCESS, failure: GET_REPO_PF_FAILURE },
  {
    request: getRepositoryPayloadFormatters,
    success: getRepositoryPayloadFormattersSuccess,
    failure: getRepositoryPayloadFormattersFailure,
  },
] = createRequestActions(GET_REPO_PF_BASE, (appId, version_ids) => ({ appId, version_ids }))
