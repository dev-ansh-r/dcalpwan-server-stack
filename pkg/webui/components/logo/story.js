

import React from 'react'

import TtsLogo from '@assets/static/logo.svg'

import ExampleLogo from './story-logo.svg'
import ExampleSquareLogo from './story-logo-2.svg'

import Logo from '.'

export default {
  title: 'Logo',
}

export const Default = () => <Logo logo={{ src: TtsLogo, alt: 'Logo' }} />

export const WithSecondaryLogo = () => (
  <Logo
    logo={{ src: TtsLogo, alt: 'Logo' }}
    brandLogo={{ src: ExampleLogo, alt: 'Secondary Logo' }}
  />
)

WithSecondaryLogo.story = {
  name: 'With secondary Logo',
}

export const WithSquareShapeSecondaryLogo = () => (
  <Logo
    logo={{ src: TtsLogo, alt: 'Logo' }}
    brandLogo={{ src: ExampleSquareLogo, alt: 'Secondary Logo' }}
  />
)

WithSquareShapeSecondaryLogo.story = {
  name: 'With square-shape secondary Logo',
}

export const WithSecondaryLogoVertical = () => (
  <Logo
    vertical
    logo={{ src: TtsLogo, alt: 'Logo' }}
    brandLogo={{ src: ExampleLogo, alt: 'Secondary Logo' }}
  />
)

WithSecondaryLogoVertical.story = {
  name: 'With secondary logo, vertical',
}

export const WithSquareShapeSecondaryLogoVertical = () => (
  <Logo
    vertical
    logo={{ src: TtsLogo, alt: 'Logo' }}
    brandLogo={{ src: ExampleSquareLogo, alt: 'Secondary Logo' }}
  />
)

WithSquareShapeSecondaryLogoVertical.story = {
  name: 'With square-shape secondary logo, vertical',
}
