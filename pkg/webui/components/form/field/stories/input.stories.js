
/* eslint-disable react/prop-types */

import React from 'react'
import { withInfo } from '@storybook/addon-info'

import Input from '@ttn-lw/components/input'
import Form from '@ttn-lw/components/form'

import { FieldsWrapperExample, info } from './shared'

export default {
  title: 'Fields/Input',
  decorators: [withInfo(info)],
}

export const Default = () => (
  <FieldsWrapperExample
    initialValues={{
      default: 'something...',
      required: 'something...',
      description: 'something...',
      warning: 'something...',
      error: 'something...',
      disabled: 'something...',
    }}
  >
    <Form.Field name="default" title="Default" component={Input} />
    <Form.Field name="xxs-size" title="XXS Size" component={Input} inputWidth="xxs" />
    <Form.Field name="xs-size" title="XS Size" component={Input} inputWidth="xs" />
    <Form.Field name="s-size" title="S Size" component={Input} inputWidth="s" />
    <Form.Field name="m-size" title="M Size" component={Input} inputWidth="m" />
    <Form.Field name="l-size" title="L Size" component={Input} inputWidth="l" />
    <Form.Field name="required" title="Required" component={Input} required />
    <Form.Field
      name="description"
      title="With Description"
      description="A select field."
      component={Input}
    />
    <Form.Field name="error" title="With Error" component={Input} />
    <Form.Field name="warning" title="With Warning" warning="A select field." component={Input} />
    <Form.Field name="disabled" title="Disabled" disabled component={Input} />
  </FieldsWrapperExample>
)
