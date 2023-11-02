

import React from 'react'

import Input from '@ttn-lw/components/input'
import Form, { useFormContext } from '@ttn-lw/components/form'

import DevEUIComponent from '@console/containers/dev-eui-component'

import tooltipIds from '@ttn-lw/lib/constants/tooltip-ids'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import m from '../../messages'

const initialValues = {
  authenticated_identifiers: {
    dev_eui: '',
    authentication_code: '',
  },
  target_device_id: '',
}

const DeviceClaimingFormSection = () => {
  const idInputRef = React.useRef(null)

  const {
    values: { _withQRdata },
  } = useFormContext()

  return (
    <>
      <DevEUIComponent name="authenticated_identifiers.dev_eui" disabled={_withQRdata} />
      <Form.Field
        title={sharedMessages.claimAuthCode}
        name="authenticated_identifiers.authentication_code"
        disabled={_withQRdata}
        component={Input}
        sensitive
      />
      <Form.Field
        required
        title={sharedMessages.devID}
        name="target_device_id"
        placeholder={sharedMessages.deviceIdPlaceholder}
        component={Input}
        inputRef={idInputRef}
        tooltipId={tooltipIds.DEVICE_ID}
        description={m.deviceIdDescription}
      />
    </>
  )
}

export { DeviceClaimingFormSection as default, initialValues }
