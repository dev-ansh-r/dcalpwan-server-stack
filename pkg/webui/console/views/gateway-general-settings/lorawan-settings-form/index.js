
import React from 'react'

import delay from '@console/constants/delays'
import frequencyPlans from '@console/constants/frequency-plans'

import SubmitButton from '@ttn-lw/components/submit-button'
import SubmitBar from '@ttn-lw/components/submit-bar'
import Form from '@ttn-lw/components/form'
import Checkbox from '@ttn-lw/components/checkbox'
import UnitInput from '@ttn-lw/components/unit-input'

import GatewayFrequencyPlansSelect from '@console/containers/freq-plans-select/gs-frequency-plan-select'

import tooltipIds from '@ttn-lw/lib/constants/tooltip-ids'
import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import { unit as unitRegexp, emptyDuration as emptyDurationRegexp } from '@console/lib/regexp'

import validationSchema from './validation-schema'

const decodeDelayValue = value => {
  if (emptyDurationRegexp.test(value)) {
    return {
      duration: undefined,
      unit: value,
    }
  }
  const duration = value.split(unitRegexp)[0]
  const unit = value.split(duration)[1]
  return {
    duration: duration ? Number(duration) : undefined,
    unit,
  }
}

const isEmptyFrequencyPlan = value => value?.includes(frequencyPlans.EMPTY_FREQ_PLAN)

const isNotValidDuration = value => {
  const { duration, unit } = decodeDelayValue(value)

  switch (unit) {
    case 'ms':
      return duration < delay.MINIMUM_GATEWAY_SCHEDULE_ANYTIME_DELAY
    case 's':
      return duration < delay.MINIMUM_GATEWAY_SCHEDULE_ANYTIME_DELAY / 1000
    case 'm':
      return duration < delay.MINIMUM_GATEWAY_SCHEDULE_ANYTIME_DELAY / 60000
    case 'h':
      return duration < delay.MINIMUM_GATEWAY_SCHEDULE_ANYTIME_DELAY / 3600000
  }
}

const LorawanSettingsForm = React.memo(props => {
  const { gateway, onSubmit } = props

  const [error, setError] = React.useState(undefined)

  const [shouldDisplayWarning, setShouldDisplayWarning] = React.useState(
    isNotValidDuration(gateway.schedule_anytime_delay),
  )

  const onScheduleAnytimeDelayChange = React.useCallback(value => {
    setShouldDisplayWarning(isNotValidDuration(value))
  }, [])

  const initialValues = React.useMemo(
    () => ({
      ...validationSchema.cast(gateway),
      frequency_plan_ids: gateway.frequency_plan_ids || [frequencyPlans.EMPTY_FREQ_PLAN],
    }),
    [gateway],
  )

  const onFormSubmit = React.useCallback(
    async (values, { resetForm, setSubmitting }) => {
      const castedValues = validationSchema.cast(
        isEmptyFrequencyPlan(values.frequency_plan_ids)
          ? { ...values, frequency_plan_ids: [''] }
          : values,
      )

      setError(undefined)
      try {
        await onSubmit(castedValues)
        resetForm({ values: castedValues })
      } catch (err) {
        setSubmitting(false)
        setError(err)
      }
    },
    [onSubmit],
  )

  return (
    <Form
      validationSchema={validationSchema}
      initialValues={initialValues}
      onSubmit={onFormSubmit}
      error={error}
      enableReinitialize
    >
      <GatewayFrequencyPlansSelect />
      <Form.Field
        title={sharedMessages.gatewayScheduleDownlinkLate}
        name="schedule_downlink_late"
        component={Checkbox}
        description={sharedMessages.scheduleDownlinkLateDescription}
        tooltipId={tooltipIds.SCHEDULE_DOWNLINK_LATE}
      />
      <Form.Field
        name="enforce_duty_cycle"
        component={Checkbox}
        label={sharedMessages.enforceDutyCycle}
        description={sharedMessages.enforceDutyCycleDescription}
        tooltipId={tooltipIds.ENFORCE_DUTY_CYCLE}
      />
      <Form.Field
        title={sharedMessages.scheduleAnyTimeDelay}
        name="schedule_anytime_delay"
        component={UnitInput.Duration}
        unitSelector={['ms', 's']}
        description={{
          ...sharedMessages.scheduleAnyTimeDescription,
          values: {
            minimumValue: delay.MINIMUM_GATEWAY_SCHEDULE_ANYTIME_DELAY,
            defaultValue: delay.DEFAULT_GATEWAY_SCHEDULE_ANYTIME_DELAY,
          },
        }}
        units={[
          { label: sharedMessages.milliseconds, value: 'ms' },
          { label: sharedMessages.seconds, value: 's' },
          { label: sharedMessages.minutes, value: 'm' },
          { label: sharedMessages.hours, value: 'h' },
        ]}
        onChange={onScheduleAnytimeDelayChange}
        warning={
          shouldDisplayWarning
            ? {
                ...sharedMessages.delayWarning,
                values: { minimumValue: delay.MINIMUM_GATEWAY_SCHEDULE_ANYTIME_DELAY },
              }
            : undefined
        }
        required
        tooltipId={tooltipIds.SCHEDULE_ANYTIME_DELAY}
      />
      <SubmitBar>
        <Form.Submit component={SubmitButton} message={sharedMessages.saveChanges} />
      </SubmitBar>
    </Form>
  )
})

LorawanSettingsForm.propTypes = {
  gateway: PropTypes.gateway.isRequired,
  onSubmit: PropTypes.func.isRequired,
}

export default LorawanSettingsForm
