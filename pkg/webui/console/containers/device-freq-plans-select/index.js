
import React from 'react'

import { CreateFrequencyPlansSelect } from '@console/containers/freq-plans-select'

import PropTypes from '@ttn-lw/lib/prop-types'

const getRegionFromBandId = (bandId = '') => bandId.split('_')[0]

const FreqPlansSelect = React.memo(props => {
  const { bandId, ...rest } = props

  return React.createElement(
    CreateFrequencyPlansSelect('ns', {
      optionsFormatter: plans => {
        const region = getRegionFromBandId(bandId)
        if (!Boolean(region)) {
          return plans.map(plan => ({ value: plan.id, label: plan.name }))
        }

        return plans
          .filter(plan => plan.id.startsWith(region))
          .map(plan => ({ value: plan.id, label: plan.name }))
      },
    }),
    rest,
  )
})

FreqPlansSelect.propTypes = {
  bandId: PropTypes.string,
}

FreqPlansSelect.defaultProps = {
  bandId: undefined,
}

export default FreqPlansSelect
