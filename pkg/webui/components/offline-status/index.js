

import React, { useEffect, useRef } from 'react'
import { defineMessages } from 'react-intl'
import classnames from 'classnames'

import ONLINE_STATUS from '@ttn-lw/constants/online-status'

import toast from '@ttn-lw/components/toast'
import Icon from '@ttn-lw/components/icon'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'
import { selectApplicationSiteTitle } from '@ttn-lw/lib/selectors/env'

import style from './offline.styl'

const siteTitle = selectApplicationSiteTitle()

const m = defineMessages({
  checking: 'Connection issues detected. Attempting to reconnect…',
  offline: '{applicationName} is offline. Please check your internet connection.',
  online: '{applicationName} is back online',
})

const handleMessage = (message, type) => {
  // Don't show a toast when the tab is not in focus
  // to prevent flooding the toast queue.
  if (document.hidden) {
    return
  }

  toast({
    messageGroup: 'offline-status',
    message: { ...message, values: { applicationName: siteTitle } },
    preventConsecutive: true,
    type,
  })
}

const OfflineStatus = ({ showOfflineOnly, showWarnings, onlineStatus }) => {
  const initialUpdate = useRef(true)
  const wasOffline = useRef(false)
  const isOnline = onlineStatus === ONLINE_STATUS.ONLINE
  const isOffline = onlineStatus === ONLINE_STATUS.OFFLINE
  const isChecking = onlineStatus === ONLINE_STATUS.CHECKING

  useEffect(() => {
    if (initialUpdate.current) {
      initialUpdate.current = false
      return
    }
    if (showWarnings) {
      if (isOnline && wasOffline.current) {
        handleMessage(m.online, toast.types.INFO)
        wasOffline.current = false
      } else if (isOffline) {
        handleMessage(m.offline, toast.types.ERROR)
        wasOffline.current = true
      }
    }
  }, [showWarnings, isOnline, isChecking, isOffline])

  if (showOfflineOnly && isOnline) {
    return null
  }

  const icon = isOnline ? 'info' : isChecking ? 'warning' : 'error'
  const message = isOnline
    ? sharedMessages.online
    : isChecking
    ? sharedMessages.connectionIssues
    : sharedMessages.offline
  const cls = classnames(style.status, {
    [style.online]: isOnline,
    [style.offline]: isOffline,
    [style.checking]: isChecking,
  })

  return (
    <span className={cls}>
      <Icon className={style.icon} icon={icon} />
      <Message content={message} />
    </span>
  )
}

OfflineStatus.propTypes = {
  onlineStatus: PropTypes.onlineStatus.isRequired,
  showOfflineOnly: PropTypes.bool,
  showWarnings: PropTypes.bool,
}

OfflineStatus.defaultProps = {
  showOfflineOnly: false,
  showWarnings: false,
}

export default OfflineStatus
