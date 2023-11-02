

import { connect } from 'react-redux'

import { listModels } from '@console/store/actions/device-repository'

import {
  selectDeviceModelsByBrandId,
  selectDeviceModelsError,
  selectDeviceModelsFetching,
} from '@console/store/selectors/device-repository'
import { selectSelectedApplicationId } from '@console/store/selectors/applications'

const mapStateToProps = (state, props) => {
  const { brandId } = props

  return {
    appId: selectSelectedApplicationId(state),
    models: selectDeviceModelsByBrandId(state, brandId),
    error: selectDeviceModelsError(state),
    fetching: selectDeviceModelsFetching(state),
  }
}

const mapDispatchToProps = { listModels }

export default ModelSelect => connect(mapStateToProps, mapDispatchToProps)(ModelSelect)
