

import React from 'react'
import classnames from 'classnames'
import { useIntl } from 'react-intl'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import style from './entity-title-section.styl'

const EntityTitleSection = props => {
  const { className, id, name, icon, iconAlt, children } = props
  const { formatMessage } = useIntl()

  return (
    <section className={classnames(className, style.container)}>
      <div className={style.titleSectionContainer}>
        <img className={style.icon} src={icon} alt={formatMessage(iconAlt)} />
        <div className={style.titleSection}>
          <h1 className={style.title}>{name || id}</h1>
          <span className={style.id}>
            <Message className={style.idPrefix} content={sharedMessages.id} uppercase />
            {id}
          </span>
        </div>
      </div>
      {children}
    </section>
  )
}

EntityTitleSection.propTypes = {
  children: PropTypes.oneOfType([PropTypes.arrayOf(PropTypes.node), PropTypes.node]),
  className: PropTypes.string,
  icon: PropTypes.string.isRequired,
  iconAlt: PropTypes.message.isRequired,
  id: PropTypes.string.isRequired,
  name: PropTypes.string,
}

EntityTitleSection.defaultProps = {
  className: undefined,
  name: undefined,
  children: null,
}

export default EntityTitleSection
