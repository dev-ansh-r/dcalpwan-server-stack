
import React from 'react'

import Link from '@ttn-lw/components/link'

import Message from '@ttn-lw/lib/components/message'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'

import style from './doc.styl'

import Tooltip from '.'

const DocTooltip = ({ content, docPath, docTitle, interactive, ...rest }) => (
  <Tooltip
    {...rest}
    interactive
    content={
      <>
        {content}
        <div className={style.docLink}>
          <Link.DocLink primary path={docPath}>
            <Message content={docTitle} />
          </Link.DocLink>
        </div>
      </>
    }
    appendTo={document.body}
  />
)

DocTooltip.propTypes = {
  ...Tooltip.propTypes,
  docPath: PropTypes.string.isRequired,
  docTitle: PropTypes.message,
}

DocTooltip.defaultProps = {
  ...Tooltip.defaultProps,
  docTitle: sharedMessages.moreInformation,
}

export default DocTooltip
