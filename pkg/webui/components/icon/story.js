

import React from 'react'
import { withInfo } from '@storybook/addon-info'

import style from './story.styl'

import Icon from '.'

const icons = [
  'devices',
  'integration',
  'settings',
  'lock',
  'lock_open',
  'close',
  'menu',
  'dashboard',
  'transform',
  'data',
  'sort',
  'overview',
  'application',
  'gateway',
  'organization',
]

const doc = `Icons can be used using \`display: {flex|inline-block}\`.
\`inline-block\` is used by default. To use \`flex\` instead, overwrite the
display value of the wrapping \`<span />\`in your local scoped css. The
positioning will differ slightly, so the nudge props can be used to fine-tune
the appearance.`

const iconElement = icons.map(icon => (
  <div className={style.wrapper} key={icon}>
    <Icon icon={icon} />
    {icon}
  </div>
))

export default {
  title: 'Icon',

  decorators: [
    withInfo({
      inline: true,
      header: false,
      propTables: [Icon],
      text: doc,
    }),
  ],
}

export const Icons = () => <div>{iconElement}</div>

export const Usage = () => (
  <div className={style.wrapper}>
    <div className={style.block}>
      <Icon icon="devices" />
      <span>{'display: inline-block'}</span>
    </div>
    <br />
    <div className={style.flex}>
      <Icon icon="devices" />
      <span>{'display: flex'}</span>
    </div>
  </div>
)
