

import React from 'react'

import Spinner from '.'

export default {
  title: 'Spinner',
}

export const Default = () => <Spinner />
export const WithChildren = () => <Spinner>This is a message</Spinner>

WithChildren.story = {
  name: 'With children',
}

export const WithInlineChildren = () => <Spinner inline>This is a message</Spinner>

WithInlineChildren.story = {
  name: 'With inline children',
}

export const Centered = () => <Spinner center>This is a message</Spinner>
export const Faded = () => <Spinner faded />
export const Small = () => <Spinner small />
export const Micro = () => <Spinner micro />
