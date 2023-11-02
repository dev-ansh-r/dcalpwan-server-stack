

import React, { useCallback } from 'react'
import { Col, Row } from 'react-grid-system'
import { defineMessages } from 'react-intl'

import Form from '@ttn-lw/components/form'
import Checkbox from '@ttn-lw/components/checkbox'
import SubmitBar from '@ttn-lw/components/submit-bar'
import SubmitButton from '@ttn-lw/components/submit-button'

import PropTypes from '@ttn-lw/lib/prop-types'
import Yup from '@ttn-lw/lib/yup'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import gatewayVisibilityMessages from '@console/lib/packet-broker/messages'

const m = defineMessages({
  saveDefaultGatewayVisibility: 'Save default gateway visibility',
})

const validationSchema = Yup.object({
  visibility: Yup.object({}),
})

const GatewayVisibilityForm = ({ onSubmit, initialValues, error }) => {
  const handleSubmit = useCallback(values => onSubmit(validationSchema.cast(values)), [onSubmit])

  return (
    <Form
      onSubmit={handleSubmit}
      initialValues={initialValues}
      error={error}
      validationSchema={validationSchema}
    >
      <Row>
        <Col md={6} xs={12}>
          <Form.Field
            name="visibility.location"
            component={Checkbox}
            label={sharedMessages.location}
          />
          <Form.Field
            name="visibility.antenna_placement"
            component={Checkbox}
            label={gatewayVisibilityMessages.gatewayAntennaPlacementLabel}
            description={gatewayVisibilityMessages.gatewayAntennaPlacementDescription}
          />
          <Form.Field
            name="visibility.antenna_count"
            component={Checkbox}
            label={gatewayVisibilityMessages.gatewayAntennaCountLabel}
          />
          <Form.Field
            name="visibility.fine_timestamps"
            component={Checkbox}
            label={gatewayVisibilityMessages.gatewayFineTimestampsLabel}
            description={gatewayVisibilityMessages.gatewayFineTimestampsDescription}
          />
        </Col>
        <Col sm={6} xs={12}>
          <Form.Field
            name="visibility.contact_info"
            component={Checkbox}
            label={gatewayVisibilityMessages.gatewayContactInfoLabel}
            description={gatewayVisibilityMessages.gatewayContactInfoDescription}
          />
          <Form.Field
            name="visibility.status"
            component={Checkbox}
            label={gatewayVisibilityMessages.gatewayStatusLabel}
            description={gatewayVisibilityMessages.gatewayStatusDescription}
          />
          <Form.Field
            name="visibility.frequency_plan"
            component={Checkbox}
            label={gatewayVisibilityMessages.gatewayFreqPlanLabel}
          />
          <Form.Field
            name="visibility.packet_rates"
            component={Checkbox}
            label={gatewayVisibilityMessages.gatewayPacketRatesLabel}
            description={gatewayVisibilityMessages.gatewayPacketRatesDescription}
          />
        </Col>
      </Row>
      <SubmitBar>
        <Form.Submit component={SubmitButton} message={m.saveDefaultGatewayVisibility} />
      </SubmitBar>
    </Form>
  )
}

GatewayVisibilityForm.propTypes = {
  error: PropTypes.error,
  initialValues: PropTypes.shape({
    visibility: PropTypes.shape({}),
  }),
  onSubmit: PropTypes.func.isRequired,
}

GatewayVisibilityForm.defaultProps = {
  error: undefined,
  initialValues: {
    visibility: {
      location: false,
      antenna_placement: false,
      antenna_count: false,
      fine_timestamps: false,
      contact_info: false,
      status: false,
      frequency_plan: false,
      packet_rates: false,
    },
  },
}

export default GatewayVisibilityForm
