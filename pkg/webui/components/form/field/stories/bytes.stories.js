
/* eslint-disable react/prop-types */

import React from 'react'
import { withInfo } from '@storybook/addon-info'

import Input from '@ttn-lw/components/input'
import Form from '@ttn-lw/components/form'

import { FieldsWrapperExample, info } from './shared'

export default {
  title: 'Fields/Byte',
  decorators: [withInfo(info)],
}

export const Default = () => (
  <FieldsWrapperExample
    initialValues={{
      default: 'ADADADAD',
      'xxs-size': 'ADAD',
      'xs-size': 'ADADADAD',
      's-size': 'ADADADADADADADAD',
      'm-size': 'ADADADADADADADADADADADADADADAD',
      'l-size': 'ADADADADADADADADADADADADADADADADADADADAD',
      description: 'ADADADAD',
      warning: 'ADADADAD',
      error: 'ADADADAD',
      disabled: 'ADADADAD',
    }}
  >
    <Form.Field
      name="default"
      title="Default"
      type="byte"
      placeholder="default"
      min={4}
      max={4}
      component={Input}
    />
    <Form.Field
      name="xxs-size"
      title="XXS Size"
      type="byte"
      placeholder="default"
      min={2}
      max={2}
      component={Input}
      inputWidth="xxs"
    />
    <Form.Field
      name="xs-size"
      title="XS Size"
      type="byte"
      placeholder="default"
      min={4}
      max={4}
      component={Input}
      inputWidth="xs"
    />
    <Form.Field
      name="s-size"
      title="S Size"
      type="byte"
      placeholder="default"
      min={8}
      max={8}
      component={Input}
      inputWidth="s"
    />
    <Form.Field
      name="m-size"
      title="M Size"
      type="byte"
      placeholder="default"
      min={15}
      max={15}
      component={Input}
      inputWidth="m"
    />
    <Form.Field
      name="l-size"
      title="L Size"
      type="byte"
      placeholder="default"
      min={20}
      max={20}
      component={Input}
      inputWidth="l"
    />
    <Form.Field
      name="description"
      title="With Description"
      description="A select field."
      type="byte"
      placeholder="description"
      min={4}
      max={4}
      component={Input}
    />
    <Form.Field
      name="error"
      title="With Error"
      type="byte"
      placeholder="error"
      min={4}
      max={4}
      component={Input}
    />
    <Form.Field
      name="warning"
      title="With Warning"
      warning="A select field."
      type="byte"
      placeholder="warning"
      min={4}
      max={4}
      component={Input}
    />
    <Form.Field
      name="disabled"
      title="Disabled"
      disabled
      placeholder="disabled"
      type="byte"
      min={4}
      max={4}
      component={Input}
    />
  </FieldsWrapperExample>
)
