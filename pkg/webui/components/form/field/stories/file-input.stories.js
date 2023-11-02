
/* eslint-disable react/prop-types */

import React from 'react'
import { withInfo } from '@storybook/addon-info'

import FileInput from '@ttn-lw/components/file-input'
import Form from '@ttn-lw/components/form'

import { FieldsWrapperExample, info } from './shared'

export default {
  title: 'Fields/FileInput',
  decorators: [withInfo(info)],
}

export const Default = () => (
  <FieldsWrapperExample
    initialValues={{
      default: '',
      withValue: 'base64-value-goes-here',
      error: '',
    }}
  >
    <Form.Field name="default" title="Default" component={FileInput} />
    <Form.Field
      name="description"
      title="With Description"
      description="A file input field."
      component={FileInput}
    />
    <Form.Field name="withValue" title="With initially attached file" component={FileInput} />
    <Form.Field name="error" title="With error" component={FileInput} />
    <Form.Field
      name="warning"
      title="With warning"
      component={FileInput}
      warning="A file input field."
    />
    <Form.Field name="disabled" title="Disabled" component={FileInput} disabled />
  </FieldsWrapperExample>
)
