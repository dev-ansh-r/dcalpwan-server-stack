

import React from 'react'
import { withInfo } from '@storybook/addon-info'

import SafeInspector from '.'

export default {
  title: 'Safe Inspector',

  decorators: [
    withInfo({
      inline: true,
      header: false,
      propTables: [SafeInspector],
    }),
  ],
}

export const Default = () => <SafeInspector data="ab01f46d2f" />

export const Multiple = () => (
  <>
    <SafeInspector data="ab01f46d2f" />
    <br />
    <SafeInspector data="ff0000" />
    <br />
    <SafeInspector data="f8a683c1d9b2" />
    <br />
    <SafeInspector data="f8a683c1d9b2" noCopy />
    <br />
    <SafeInspector data="f8a683c1d9b2" noCopy noTransform />
    <br />
    <SafeInspector data="f8a683c1d9b2" noTransform />
  </>
)

export const Small = () => <SafeInspector small data="ab01f46d2f" />

export const NonByte = () => (
  <SafeInspector data="somerandomhash.du9d8321-9dk2-edf2398efh8" isBytes={false} />
)

export const InitiallyVisible = () => <SafeInspector data="ab01f46d2f" initiallyVisible />
export const NotHideable = () => <SafeInspector data="ab01f46d2f" hideable={false} />

NotHideable.story = {
  name: 'Not hideable',
}
