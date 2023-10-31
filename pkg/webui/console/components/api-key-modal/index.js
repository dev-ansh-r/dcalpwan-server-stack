

import React from 'react'
import { defineMessages } from 'react-intl'

import PortalledModal from '@ttn-lw/components/modal/portalled'
import SafeInspector from '@ttn-lw/components/safe-inspector'
import Icon from '@ttn-lw/components/icon'

import Message from '@ttn-lw/lib/components/message'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'

import style from './api-key-modal.styl'

const m = defineMessages({
  title: 'Please copy newly created API key',
  subtitle: "You won't be able to view the key afterward",
  buttonMessage: 'I have copied the key',
  grantedRights: 'Granted rights',
  description: `Your API key has been created successfully.
Note: After closing this window, the value of the key secret will not be accessible anymore.
Make sure to copy and store it in a safe place now.`,
})

const ApiKeyModal = props => {
  const { visible, secret, rights, ...rest } = props

  if (!visible) {
    return null
  }

  return (
    <PortalledModal
      visible={visible}
      {...rest}
      title={m.title}
      subtitle={m.subtitle}
      approval={false}
      buttonMessage={m.buttonMessage}
    >
      <div className={style.left}>
        <Message component="h4" content={m.grantedRights} />
        <ul>
          {rights.map(right => (
            <li key={right}>
              <Icon icon="check" className={style.icon} />
              <Message className={style.rightName} content={{ id: `enum:${right}` }} firstToUpper />
            </li>
          ))}
        </ul>
      </div>
      <div className={style.right}>
        <Message className={style.description} component="p" content={m.description} />
        <Message component="h3" content={sharedMessages.apiKey} />
        <SafeInspector
          className={style.secretInspector}
          data={secret}
          isBytes={false}
          disableResize
        />
      </div>
    </PortalledModal>
  )
}

ApiKeyModal.propTypes = {
  rights: PropTypes.rights,
  secret: PropTypes.string,
  visible: PropTypes.bool.isRequired,
}

ApiKeyModal.defaultProps = {
  rights: undefined,
  secret: undefined,
}

export default ApiKeyModal
