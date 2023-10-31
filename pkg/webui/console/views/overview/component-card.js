

import React from 'react'

import ServerIcon from '@assets/auxiliary-icons/server.svg'

import Status from '@ttn-lw/components/status'

import Message from '@ttn-lw/lib/components/message'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'

import style from './overview.styl'

const ComponentCard = ({ name, enabled, host }) => (
  <div className={style.componentCard}>
    <img src={ServerIcon} className={style.componentCardIcon} />
    <div className={style.componentCardDesc}>
      <div className={style.componentCardName}>
        <Status label={name} status={enabled ? 'good' : 'unknown'} flipped />
      </div>
      <span className={style.componentCardHost} title={host}>
        {enabled ? host : <Message content={sharedMessages.disabled} />}
      </span>
    </div>
  </div>
)

ComponentCard.propTypes = {
  enabled: PropTypes.bool.isRequired,
  host: PropTypes.string,
  name: PropTypes.message.isRequired,
}

ComponentCard.defaultProps = {
  host: undefined,
}

export default ComponentCard
