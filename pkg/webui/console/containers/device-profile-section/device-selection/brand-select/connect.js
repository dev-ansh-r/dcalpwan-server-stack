

import { connect } from 'react-redux'

import {
  selectDeviceBrands,
  selectDeviceBrandsError,
  selectDeviceBrandsFetching,
} from '@console/store/selectors/device-repository'
import { selectSelectedApplicationId } from '@console/store/selectors/applications'

const mapStateToProps = state => ({
  appId: selectSelectedApplicationId(state),
  brands: selectDeviceBrands(state),
  error: selectDeviceBrandsError(state),
  fetching: selectDeviceBrandsFetching(state),
})

const mapDispatchToProps = {}

export default BrandSelect => connect(mapStateToProps, mapDispatchToProps)(BrandSelect)
