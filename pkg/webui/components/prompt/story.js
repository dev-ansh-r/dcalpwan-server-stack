

import React from 'react'
import { withInfo } from '@storybook/addon-info'
import { action } from '@storybook/addon-actions'

import Link from '@ttn-lw/components/link'

import Prompt from '.'

const shouldBlockNavigation = location => location.pathname.endsWith('should-block')

const linkContainerStyle = {
  margin: '10px 0',
}

export default {
  title: 'Prompt',

  decorators: [
    withInfo({
      inline: true,
      header: false,
      source: true,
      propTables: [Prompt],
    }),
  ],
}

export const Default = () => (
  <div>
    Navigate using the links below to trigger the `Prompt` component to appear.
    <div style={linkContainerStyle}>
      <Link to="/should-not-block" title="should not block">
        should not block
      </Link>
    </div>
    <div style={linkContainerStyle}>
      <Link to="/should-block" title="should block">
        should block
      </Link>
    </div>
    <div style={linkContainerStyle}>
      <Link to="/should-block" title="should also block">
        should block
      </Link>
    </div>
    <Prompt
      when
      onApprove={action('onApprove')}
      onCancel={action('onCancel')}
      shouldBlockNavigation={shouldBlockNavigation}
      modal={{
        title: 'Block modal title',
      }}
    >
      <p>
        This modal prompts the user to confirm current route change. By pressing `Cancel` the user
        blocks navigation, while pressing `Approve` allows navigation to happen.
      </p>
    </Prompt>
  </div>
)
