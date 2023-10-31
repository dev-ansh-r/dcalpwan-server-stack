

import React from 'react'
import { Col, Row } from 'react-grid-system'

import DeviceImportForm from '@console/components/device-import-form'

import PropTypes from '@ttn-lw/lib/prop-types'

import m from './messages'

const Form = ({ handleSubmit, jsEnabled }) => {
  const initialValues = {
    format_id: '',
    data: '',
    set_claim_auth_code: false,
    _inputMethod: 'no-fallback',
    frequency_plan_id: '',
    lorawan_version: '',
    lorawan_phy_version: '',
    version_ids: {
      brand_id: '',
      model_id: '',
      firmware_version: '',
      hardware_version: '',
      band_id: '',
    },
  }
  const largeFile = 10 * 1024 * 1024

  return (
    <Row>
      <Col md={8}>
        <DeviceImportForm
          jsEnabled={jsEnabled}
          largeFileWarningMessage={m.largeFileWarningMessage}
          warningSize={largeFile}
          initialValues={initialValues}
          onSubmit={handleSubmit}
        />
      </Col>
    </Row>
  )
}

Form.propTypes = {
  handleSubmit: PropTypes.func.isRequired,
  jsEnabled: PropTypes.bool.isRequired,
}

export default Form
