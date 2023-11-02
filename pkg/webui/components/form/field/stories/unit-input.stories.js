
/* eslint-disable react/prop-types */

import React from 'react'
import { withInfo } from '@storybook/addon-info'

import UnitInput from '@ttn-lw/components/unit-input'
import Form from '@ttn-lw/components/form'

import { FieldsWrapperExample, info } from './shared'

export default {
  title: 'Fields/UnitInput',
  decorators: [withInfo(info)],
}

export const Default = () => (
  <FieldsWrapperExample
    initialValues={{
      default: '530ms',
      description: '530ms',
      warning: '530ms',
      error: '530ms',
      disabled: '530ms',
    }}
  >
    <Form.Field name="default" title="Default" component={UnitInput.Duration} />
    <Form.Field
      name="description"
      title="Description"
      component={UnitInput.Duration}
      description="The unit input"
    />
    <Form.Field
      name="warning"
      title="Warning"
      component={UnitInput.Duration}
      warning="The unit input"
    />
    <Form.Field name="error" title="Error" component={UnitInput.Duration} />
  </FieldsWrapperExample>
)
