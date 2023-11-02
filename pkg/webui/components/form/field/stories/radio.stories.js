
/* eslint-disable react/prop-types */

import React from 'react'
import { withInfo } from '@storybook/addon-info'

import Radio from '@ttn-lw/components/radio-button'
import Form from '@ttn-lw/components/form'

import { FieldsWrapperExample, info } from './shared'

export default {
  title: 'Fields/Radio',
  decorators: [withInfo(info)],
}

export const Default = () => (
  <FieldsWrapperExample
    initialValues={{
      default: '1',
      description: '2',
      warning: '3',
      error: '4',
      disabled: '5',
    }}
  >
    <Form.Field name="default" title="Default" label="Radio" checked component={Radio} />
    <Form.Field
      name="description"
      title="With Description"
      description="A select field."
      label="Radio"
      checked
      component={Radio}
    />
    <Form.Field name="error" title="With Error" label="Radio" checked component={Radio} />
    <Form.Field
      name="warning"
      title="With Warning"
      warning="A select field."
      label="Radio"
      checked
      component={Radio}
    />
    <Form.Field name="disabled" title="Disabled" label="Radio" disabled checked component={Radio} />
  </FieldsWrapperExample>
)

export const HorizontalGroup = () => (
  <FieldsWrapperExample
    initialValues={{
      default: '1',
      description: '1',
      warning: '1',
      error: '1',
      disabled: '1',
    }}
  >
    <Form.Field name="default" title="Default" component={Radio.Group} horizontal>
      <Radio label="Radio 1" value="1" />
      <Radio label="Radio 2" value="2" />
      <Radio label="Radio 3" value="3" />
    </Form.Field>
    <Form.Field
      name="description"
      title="With Description"
      description="A select field."
      component={Radio.Group}
      horizontal
    >
      <Radio label="Radio 1" value="1" />
      <Radio label="Radio 2" value="2" />
      <Radio label="Radio 3" value="3" />
    </Form.Field>
    <Form.Field name="error" title="With Error" component={Radio.Group} horizontal>
      <Radio label="Radio 1" value="1" />
      <Radio label="Radio 2" value="2" />
      <Radio label="Radio 3" value="3" />
    </Form.Field>
    <Form.Field
      name="warning"
      title="With Warning"
      warning="A select field."
      component={Radio.Group}
      horizontal
    >
      <Radio label="Radio 1" value="1" />
      <Radio label="Radio 2" value="2" />
      <Radio label="Radio 3" value="3" />
    </Form.Field>
    <Form.Field name="disabled" title="Disabled" disabled component={Radio.Group} horizontal>
      <Radio label="Radio 1" value="1" />
      <Radio label="Radio 2" value="2" />
      <Radio label="Radio 3" value="3" />
    </Form.Field>
  </FieldsWrapperExample>
)

export const RowGroup = () => (
  <FieldsWrapperExample
    initialValues={{
      default: '1',
      description: '1',
      warning: '1',
      error: '1',
      disabled: '1',
    }}
  >
    <Form.Field name="default" title="Default" component={Radio.Group}>
      <Radio label="Radio 1" value="1" />
      <Radio label="Radio 2" value="2" />
      <Radio label="Radio 3" value="3" />
    </Form.Field>
    <Form.Field
      name="description"
      title="With Description"
      description="A select field."
      component={Radio.Group}
    >
      <Radio label="Radio 1" value="1" />
      <Radio label="Radio 2" value="2" />
      <Radio label="Radio 3" value="3" />
    </Form.Field>
    <Form.Field name="error" title="With Error" component={Radio.Group}>
      <Radio label="Radio 1" value="1" />
      <Radio label="Radio 2" value="2" />
      <Radio label="Radio 3" value="3" />
    </Form.Field>
    <Form.Field
      name="warning"
      title="With Warning"
      warning="A select field."
      component={Radio.Group}
    >
      <Radio label="Radio 1" value="1" />
      <Radio label="Radio 2" value="2" />
      <Radio label="Radio 3" value="3" />
    </Form.Field>
    <Form.Field name="disabled" title="Disabled" disabled component={Radio.Group}>
      <Radio label="Radio 1" value="1" />
      <Radio label="Radio 2" value="2" />
      <Radio label="Radio 3" value="3" />
    </Form.Field>
  </FieldsWrapperExample>
)
