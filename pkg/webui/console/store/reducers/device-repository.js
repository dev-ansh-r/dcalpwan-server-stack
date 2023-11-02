
import { handleActions } from 'redux-actions'
import { get } from 'lodash'

import {
  LIST_BRANDS_SUCCESS,
  LIST_MODELS_SUCCESS,
  GET_MODEL_SUCCESS,
  GET_TEMPLATE_SUCCESS,
  GET_REPO_PF_SUCCESS,
} from '@console/store/actions/device-repository'

export const defaultState = {
  brands: {
    list: [],
    totalCount: 0,
  },
  models: {},
  template: undefined,
  repo_payload_formatters: undefined,
}

const handleListBrands = (state, payload) => {
  const { brands: list, totalCount } = payload

  return {
    ...state,
    brands: {
      list,
      totalCount,
    },
  }
}
const handleListModels = (state, payload) => {
  const { brandId, models: list, totalCount } = payload

  return {
    ...state,
    models: {
      ...state.models,
      [brandId]: {
        list,
        totalCount,
      },
    },
  }
}
const handleGetModel = (state, payload) => {
  const { brand_id, model_id } = payload
  const models = get(state, `models.${brand_id}.list`, [])
  const modelIdx = models.findIndex(m => m.model_id === model_id)

  if (modelIdx === -1) {
    return {
      ...state,
      models: {
        ...state.models,
        [brand_id]: {
          list: models.concat(payload),
        },
      },
    }
  }

  return {
    ...state,
    models: {
      ...state.models,
      [brand_id]: {
        list: models.map(m => (m.model_id === model_id ? { ...m, ...payload } : m)),
      },
    },
  }
}
const handleGetTemplate = (state, payload) => ({ ...state, template: payload })

const handleGetRepositoryPayloadFromatters = (state, payload) => ({
  ...state,
  repo_payload_formatters: payload,
})

export default handleActions(
  {
    [LIST_BRANDS_SUCCESS]: (state, { payload }) => handleListBrands(state, payload),
    [LIST_MODELS_SUCCESS]: (state, { payload }) => handleListModels(state, payload),
    [GET_MODEL_SUCCESS]: (state, { payload }) => handleGetModel(state, payload),
    [GET_TEMPLATE_SUCCESS]: (state, { payload }) => handleGetTemplate(state, payload),
    [GET_REPO_PF_SUCCESS]: (state, { payload }) =>
      handleGetRepositoryPayloadFromatters(state, payload),
  },
  defaultState,
)
