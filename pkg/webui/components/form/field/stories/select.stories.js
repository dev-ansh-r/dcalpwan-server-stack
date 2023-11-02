
/* eslint-disable react/prop-types */

import React from 'react'
import { withInfo } from '@storybook/addon-info'

import Select from '@ttn-lw/components/select'
import Form from '@ttn-lw/components/form'

import { FieldsWrapperExample, info } from './shared'

export default {
  title: 'Fields/Select',
  decorators: [withInfo(info)],
}

export const Default = () => (
  <FieldsWrapperExample
    initialValues={{
      default: 'amsterdam',
      description: 'amsterdam',
      warning: 'amsterdam',
      error: 'amsterdam',
      disabled: 'amsterdam',
    }}
  >
    <Form.Field
      name="default"
      title="Default"
      component={Select}
      options={[
        { value: 'amsterdam', label: 'Amsterdam' },
        { value: 'berlin', label: 'Berlin' },
        { value: 'dusseldorf', label: 'Düsseldorf' },
      ]}
    />
    <Form.Field
      name="description"
      title="With Description"
      description="A select field."
      component={Select}
      options={[
        { value: 'amsterdam', label: 'Amsterdam' },
        { value: 'berlin', label: 'Berlin' },
        { value: 'dusseldorf', label: 'Düsseldorf' },
      ]}
    />
    <Form.Field
      name="error"
      title="With Error"
      component={Select}
      options={[
        { value: 'amsterdam', label: 'Amsterdam' },
        { value: 'berlin', label: 'Berlin' },
        { value: 'dusseldorf', label: 'Düsseldorf' },
      ]}
    />
    <Form.Field
      name="warning"
      title="With Warning"
      warning="A select field."
      component={Select}
      options={[
        { value: 'amsterdam', label: 'Amsterdam' },
        { value: 'berlin', label: 'Berlin' },
        { value: 'dusseldorf', label: 'Düsseldorf' },
      ]}
    />
    <Form.Field
      name="disabled"
      title="Disabled"
      disabled
      component={Select}
      options={[
        { value: 'amsterdam', label: 'Amsterdam' },
        { value: 'berlin', label: 'Berlin' },
        { value: 'dusseldorf', label: 'Düsseldorf' },
      ]}
    />
  </FieldsWrapperExample>
)
