
/* eslint-disable react/prop-types */

import React from 'react'
import { action } from '@storybook/addon-actions'

import Yup from '@ttn-lw/lib/yup'

import Form from '../..'

const handleSubmit = (data, { resetForm }) => {
  action('Submit')(data)
  setTimeout(() => resetForm({ values: data }), 1000)
}

const info = {
  inline: true,
  header: false,
  source: false,
  propTables: [Form.Field],
}

const errorSchema = Yup.string().test('error', 'Something went wrong.', () => false)
const validationSchema = Yup.object().shape({
  error: errorSchema,
})

class FieldsWrapperExample extends React.Component {
  form = React.createRef()

  componentDidMount() {
    if (this.form.current) {
      this.form.current.setFieldError('error', 'Something went wrong.')
      this.form.current.setFieldTouched('error')
    }
  }

  render() {
    return (
      <Form
        onSubmit={handleSubmit}
        initialValues={this.props.initialValues}
        formikRef={this.form}
        validationSchema={validationSchema}
      >
        {this.props.children}
      </Form>
    )
  }
}

export { info, FieldsWrapperExample }
