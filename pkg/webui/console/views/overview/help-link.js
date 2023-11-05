import React from 'react'
import { defineMessages } from 'react-intl'

import Link from '@ttn-lw/components/link'

import Message from '@ttn-lw/lib/components/message'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'

import style from './overview.styl'

const m = defineMessages({
  needHelp: 'Need help? Have a look at {githubRepositoryLink}.',
  needHelpShort: 'Check github documentation for support.',
})

const HelpLink = ({ supportLink, documentationLink }) => {
  if (!supportLink && !documentationLink) return null

  const githubRepositoryLink = "https://github.com/yourusername/yourrepository"; // Replace with your GitHub repository URL

  const documentation = (
    <a href={githubRepositoryLink} target="_blank" rel="noopener noreferrer">
      GitHub Repository
    </a>
  );
  

  const support = (
    <Link.Anchor secondary href={supportLink || ''} external>
      <Message content={sharedMessages.getSupport} />
    </Link.Anchor>
  )

  return (
    <Message
      className={style.getStarted}
      content={documentationLink && supportLink ? m.needHelp : m.needHelpShort}
      values={{
        documentationLink: documentation,
        supportLink: support,
        link: documentationLink ? documentation : support,
      }}
      component="h2"
    />
  )
}

HelpLink.propTypes = {
  documentationLink: PropTypes.string,
  supportLink: PropTypes.string,
}

HelpLink.defaultProps = {
  supportLink: undefined,
  documentationLink: undefined,
}

export default HelpLink
