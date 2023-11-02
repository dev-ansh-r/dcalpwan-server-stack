
import React from 'react'
import { unionBy } from 'lodash'
import { defineMessages } from 'react-intl'

import tts from '@console/api/tts'

import toast from '@ttn-lw/components/toast'
import Select from '@ttn-lw/components/select'

import PropTypes from '@ttn-lw/lib/prop-types'

import {
  PHY_V1_0,
  PHY_V1_0_1,
  PHY_V1_0_2_REV_A,
  PHY_V1_0_2_REV_B,
  PHY_V1_0_3_REV_A,
  PHY_V1_1_REV_A,
  PHY_V1_1_REV_B,
  MAC_V1_0,
  MAC_V1_0_1,
  MAC_V1_0_2,
  MAC_V1_0_3,
  MAC_V1_0_4,
  MAC_V1_1,
  RP002_V1_0_0,
  RP002_V1_0_1,
  RP002_V1_0_2,
  RP002_V1_0_3,
  RP002_V1_0_4,
  LORAWAN_VERSIONS,
  parseLorawanMacVersion,
} from '@console/lib/device-utils'

const phyVersionsMap = {
  [PHY_V1_0.value]: [MAC_V1_0, MAC_V1_0_4],
  [PHY_V1_0_1.value]: [MAC_V1_0_1, MAC_V1_0_4],
  [PHY_V1_0_2_REV_A.value]: [MAC_V1_0_2, MAC_V1_0_4],
  [PHY_V1_0_2_REV_B.value]: [MAC_V1_0_2, MAC_V1_0_4],
  [PHY_V1_0_3_REV_A.value]: [MAC_V1_0_3, MAC_V1_0_4],
  [PHY_V1_1_REV_A.value]: [MAC_V1_1, MAC_V1_0_4],
  [PHY_V1_1_REV_B.value]: [MAC_V1_1, MAC_V1_0_4],
  [RP002_V1_0_0.value]: [MAC_V1_1, MAC_V1_0_4],
  [RP002_V1_0_1.value]: [MAC_V1_1, MAC_V1_0_4],
  [RP002_V1_0_2.value]: [MAC_V1_1, MAC_V1_0_4],
  [RP002_V1_0_3.value]: [MAC_V1_1, MAC_V1_0_4],
  [RP002_V1_0_4.value]: [MAC_V1_1, MAC_V1_0_4],
}
const m = defineMessages({
  phyVersionError: 'Failed to fetch regional parameters versions',
})

const LorawanVersionInput = props => {
  const { frequencyPlan, onChange, value, ...rest } = props

  const [phyVersions, setPhyVersions] = React.useState([])
  const [options, setOptions] = React.useState(LORAWAN_VERSIONS)

  React.useEffect(() => {
    const fetchPhyVersions = async () => {
      try {
        const { version_info } = await tts.Configuration.getPhyVersions()
        setPhyVersions(version_info)
      } catch (err) {
        toast({
          type: toast.types.ERROR,
          message: m.phyVersionError,
        })
      }
    }

    fetchPhyVersions()
  }, [])

  React.useEffect(() => {
    const currentPhyVersions = phyVersions.filter(({ band_id }) => frequencyPlan.includes(band_id))
    if (currentPhyVersions.length > 0) {
      const versions = currentPhyVersions[0].phy_versions
      const options = unionBy(
        ...versions.map(version => phyVersionsMap[version]),
        v => v.value,
      ).sort((a, b) => parseLorawanMacVersion(a.value) - parseLorawanMacVersion(b.value))
      setOptions(options)
    }
  }, [frequencyPlan, phyVersions])

  return <Select onChange={onChange} value={value} options={options} {...rest} />
}

LorawanVersionInput.propTypes = {
  frequencyPlan: PropTypes.string,
  onChange: PropTypes.func.isRequired,
  value: PropTypes.string,
}

LorawanVersionInput.defaultProps = {
  value: undefined,
  frequencyPlan: '',
}

export default LorawanVersionInput
