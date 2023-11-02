
/* eslint-disable react/prop-types */

import crypto from 'crypto'

import React from 'react'

import Input from '..'

import { Example } from './shared'

const generateRandom16Bytes = () => crypto.randomBytes(16).toString('hex').toUpperCase()

export default {
  title: 'Input/Normal',
}

export const Default = () => (
  <div>
    <Example label="Username" />
    <Example label="Username" warning />
    <Example label="Username" error />
  </div>
)

export const WithPlaceholder = () => <Example placeholder="Placeholder..." />
export const WithIcon = () => <Example icon="search" />
export const WithAppend = () => <Example error append="test" />

WithIcon.story = {
  name: 'With icon',
}

export const Valid = () => <Example valid />
export const Disabled = () => <Example value="1234" disabled />
export const Readonly = () => <Example value="1234" readOnly />
export const Password = () => <Example type="password" />
export const Number = () => <Example type="number" />

export const Toggled = () => (
  <Example component={Input.Toggled} type="toggled" enabledMessage="Enabled" />
)

export const Textarea = () => <Example type="textarea" />
export const WithSpinner = () => <Example icon="search" loading />
export const Sensitive = () => <Example sensitive max={5} />

export const WithAction = () => (
  <div>
    <Example action={{ icon: 'build', secondary: true }} />
    <Example action={{ icon: 'build', secondary: true }} warning />
    <Example action={{ icon: 'build', secondary: true }} error />
  </div>
)

export const Generate = () => (
  <Example
    type="byte"
    component={Input.Generate}
    onGenerateValue={generateRandom16Bytes}
    min={16}
    max={16}
  />
)
