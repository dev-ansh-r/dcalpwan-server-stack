
/* eslint-disable react/prop-types */

import React from 'react'
import { withInfo } from '@storybook/addon-info'

import Input from '@ttn-lw/components/input'
import Form from '@ttn-lw/components/form'

import { FieldsWrapperExample, info } from './shared'

export default {
  title: 'Fields/TextArea',
  decorators: [withInfo(info)],
}

export const Default = () => (
  <FieldsWrapperExample
    initialValues={{
      default: 'something...',
      description: 'something...',
      warning: 'something...',
      error: 'something...',
      disabled: 'something...',
    }}
  >
    <Form.Field name="default" title="Default" type="textarea" component={Input} />
    <Form.Field
      name="description"
      title="With Description"
      description="A select field."
      type="textarea"
      component={Input}
    />
    <Form.Field name="error" title="With Error" type="textarea" component={Input} />
    <Form.Field
      name="warning"
      title="With Warning"
      warning="A select field."
      type="textarea"
      component={Input}
    />
    <Form.Field name="disabled" title="Disabled" disabled type="textarea" component={Input} />
  </FieldsWrapperExample>
)
