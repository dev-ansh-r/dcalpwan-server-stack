

import React from 'react'
import { defineMessages } from 'react-intl'

import Notification from '@ttn-lw/components/notification'
import Link from '@ttn-lw/components/link'

import PropTypes from '@ttn-lw/lib/prop-types'

const m = defineMessages({
  hintTitle: 'Your end device will be added soon!',
  hintMessage: `We're sorry, but your device is not yet part of The LoRaWAN Device Repository. You can use <b>enter end device specifics manually</b> option above, using the information your end device manufacturer provided e.g. in the product's data sheet. Please also refer to our documentation on <GuideLink>Adding Devices</GuideLink>.`,
})

const OtherHint = React.memo(props => {
  const { manualGuideDocsPath } = props

  return (
    <Notification
      info
      small
      title={m.hintTitle}
      content={m.hintMessage}
      messageValues={{
        GuideLink: msg => (
          <Link.DocLink secondary key="manual-guide-link" path={manualGuideDocsPath}>
            {msg}
          </Link.DocLink>
        ),
        b: msg => <b>{msg}</b>,
      }}
    />
  )
})

OtherHint.propTypes = {
  manualGuideDocsPath: PropTypes.string,
}

OtherHint.defaultProps = {
  manualGuideDocsPath: undefined,
}

export default OtherHint
