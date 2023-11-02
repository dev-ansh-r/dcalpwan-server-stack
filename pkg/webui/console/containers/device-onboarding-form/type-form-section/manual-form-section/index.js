

import React from 'react'

import Form, { useFormContext } from '@ttn-lw/components/form'

import Message from '@ttn-lw/lib/components/message'

import LorawanVersionInput from '@console/components/lorawan-version-input'
import PhyVersionInput from '@console/components/phy-version-input'

import { NsFrequencyPlansSelect } from '@console/containers/freq-plans-select'

import tooltipIds from '@ttn-lw/lib/constants/tooltip-ids'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import { frequencyPlanValueSetter, lorawanVersionValueSetter } from '@console/lib/device-utils'

import m from '../../messages'

import AdvancedSettingsSection, {
  initialValues as advancedSettingsInitialValues,
} from './advanced-settings-section'

const initialValues = {
  lorawan_version: '',
  lorawan_phy_version: '',
  frequency_plan_id: '',
  ...advancedSettingsInitialValues,
}

const DeviceTypeManualFormSection = () => {
  const {
    values: { frequency_plan_id, lorawan_version, lorawan_phy_version },
  } = useFormContext()

  const hasCompleted = Boolean(frequency_plan_id && lorawan_version && lorawan_phy_version)

  return (
    <>
      <NsFrequencyPlansSelect
        required
        tooltipId={tooltipIds.FREQUENCY_PLAN}
        name="frequency_plan_id"
        valueSetter={frequencyPlanValueSetter}
      />
      <Form.Field
        required
        title={sharedMessages.macVersion}
        name="lorawan_version"
        component={LorawanVersionInput}
        tooltipId={tooltipIds.LORAWAN_VERSION}
        frequencyPlan={frequency_plan_id}
        valueSetter={lorawanVersionValueSetter}
      />
      <Form.Field
        required
        title={sharedMessages.phyVersion}
        name="lorawan_phy_version"
        component={PhyVersionInput}
        tooltipId={tooltipIds.REGIONAL_PARAMETERS}
        lorawanVersion={lorawan_version}
      />
      {!hasCompleted && <Message content={m.continueManual} className="mb-ls-m" component="div" />}
      <AdvancedSettingsSection />
    </>
  )
}

export { DeviceTypeManualFormSection as default, initialValues }
