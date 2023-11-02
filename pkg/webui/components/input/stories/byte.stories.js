
/* eslint-disable react/prop-types */

import crypto from 'crypto'

import React from 'react'

import Input from '..'

import { Example } from './shared'

const generateRandom16Bytes = () => crypto.randomBytes(16).toString('hex').toUpperCase()

export default {
  title: 'Input/Byte',
}

export const Byte = () => <Example type="byte" min={1} max={5} />
export const ByteReadOnly = () => (
  <Example type="byte" min={1} max={5} value="A0BF49A464" readOnly />
)

ByteReadOnly.story = {
  name: 'Byte read-only',
}

export const Sensitive = () => <Example type="byte" sensitive max={5} />

export const Generate = () => (
  <Example
    type="byte"
    component={Input.Generate}
    onGenerateValue={generateRandom16Bytes}
    min={16}
    max={16}
  />
)
