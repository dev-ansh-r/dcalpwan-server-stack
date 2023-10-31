

import React from 'react'
import { defineMessages } from 'react-intl'

import Notification from '@ttn-lw/components/notification'
import Button from '@ttn-lw/components/button'

import RegistryTotals from './registry-totals'

const m = defineMessages({
  openSourceInfo:
    'You are currently using The Things Stack Open Source. More features can be unlocked by using The Things Stack Cloud.',
  plansButton: 'Get started with The Things Stack Cloud',
})

const NetworkInformationContainer = () => (
  <>
    <Notification
      info
      small
      content={m.openSourceInfo}
      children={
        <Button.AnchorLink
          primary
          href={'https://www.thethingsindustries.com/stack/plans/'}
          message={m.plansButton}
          target="_blank"
          external
        />
      }
      className="mt-cs-l"
    />
    <RegistryTotals />
  </>
)

export default NetworkInformationContainer
