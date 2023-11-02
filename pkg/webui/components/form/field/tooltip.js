
import React from 'react'
import { defineMessages } from 'react-intl'

import Tooltip from '@ttn-lw/components/tooltip'
import Icon from '@ttn-lw/components/icon'
import Link from '@ttn-lw/components/link'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'
import { descriptions, links } from '@ttn-lw/lib/field-description-messages'

import style from './field.styl'

const m = defineMessages({
  descriptionTitle: 'What is this?',
  locationTitle: 'What should I enter here?',
  absenceTitle: 'What if I cannot find the correct value?',
  viewGlossaryPage: 'View glossary page',
  readMore: 'Read more',
})

const Content = props => {
  const { tooltipDescription, glossaryTerm, children } = props
  const { description, location, absence, glossaryId } = tooltipDescription

  const hasLocation = Boolean(location)
  const hasAbsence = Boolean(absence)
  const hasChildren = Boolean(children)
  const hasGlossary = Boolean(glossaryId)

  return (
    <div className={style.tooltipContent}>
      <Message className={style.tooltipTitle} content={m.descriptionTitle} component="h4" />
      <Message
        className={style.tooltipDescription}
        content={description}
        component="p"
        convertBackticks
      />
      {hasLocation && (
        <>
          <Message className={style.tooltipTitle} content={m.locationTitle} component="h4" />
          <Message
            className={style.tooltipDescription}
            content={location}
            component="p"
            convertBackticks
          />
        </>
      )}
      {hasAbsence && (
        <>
          <Message className={style.tooltipTitle} content={m.absenceTitle} component="h4" />
          <Message
            className={style.tooltipDescription}
            content={absence}
            component="p"
            convertBackticks
          />
        </>
      )}
      {(hasChildren || hasGlossary) && (
        <div className={style.tooltipLinks}>
          {children}
          {hasGlossary && (
            <Link.GlossaryLink
              term={glossaryTerm}
              glossaryId={glossaryId}
              title={m.viewGlossaryPage}
              primary
            />
          )}
        </div>
      )}
    </div>
  )
}

Content.propTypes = {
  children: PropTypes.node,
  glossaryTerm: PropTypes.message,
  tooltipDescription: PropTypes.shape({
    description: PropTypes.message.isRequired,
    location: PropTypes.message,
    absence: PropTypes.message,
    glossaryId: PropTypes.string,
  }).isRequired,
}

Content.defaultProps = {
  children: null,
  glossaryTerm: undefined,
}

const FieldTooltip = React.memo(props => {
  const { id, glossaryTerm } = props

  const tooltipDescription = descriptions[id]
  if (!tooltipDescription) {
    return null
  }

  const tooltipAdditionalLink = links[id]
  let link = null
  if (tooltipAdditionalLink) {
    const { documentationPath, externalUrl } = tooltipAdditionalLink
    if (documentationPath) {
      link = (
        <Link.DocLink primary path={documentationPath}>
          <Message content={m.readMore} />
        </Link.DocLink>
      )
    } else if (externalUrl) {
      link = (
        <Link.Anchor primary href={externalUrl} target="_blank">
          <Message content={m.readMore} />
        </Link.Anchor>
      )
    }
  }

  return (
    <Tooltip
      className={style.tooltip}
      placement="bottom-start"
      interactive
      content={
        <Content
          glossaryTerm={glossaryTerm}
          tooltipDescription={tooltipDescription}
          children={link}
        />
      }
    >
      <Icon className={style.tooltipIcon} icon="help_outline" />
    </Tooltip>
  )
})

FieldTooltip.propTypes = {
  glossaryTerm: PropTypes.message,
  id: PropTypes.string.isRequired,
}

FieldTooltip.defaultProps = {
  glossaryTerm: undefined,
}

export default FieldTooltip
