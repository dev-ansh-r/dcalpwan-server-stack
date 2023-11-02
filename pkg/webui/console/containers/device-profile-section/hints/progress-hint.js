

import React from 'react'
import { defineMessages } from 'react-intl'

import Link from '@ttn-lw/components/link'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'

const m = defineMessages({
  hintMessage:
    'Cannot find your exact end device? <SupportLink>Get help here</SupportLink> and try <b>enter end device specifics manually</b> option above.',
  hintNoSupportMessage:
    'Cannot find your exact end device? Try <b>enter end device specifics manually</b> option above.',
})

const ProgressHint = React.memo(props => {
  const { supportLink } = props

  return (
    <Message
      className="mb-ls-xs"
      component="div"
      content={Boolean(supportLink) ? m.hintMessage : m.hintNoSupportMessage}
      values={{
        SupportLink: msg => (
          <Link.Anchor secondary key="support-link" href={supportLink} target="_blank">
            {msg}
          </Link.Anchor>
        ),
        b: msg => <b>{msg}</b>,
      }}
    />
  )
})

ProgressHint.propTypes = {
  supportLink: PropTypes.string,
}

ProgressHint.defaultProps = {
  supportLink: undefined,
}

export default ProgressHint
