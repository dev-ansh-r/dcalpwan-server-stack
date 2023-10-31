

import React, { useCallback } from 'react'
import { defineMessages } from 'react-intl'
import classnames from 'classnames'

import Button from '@ttn-lw/components/button'
import Dropdown from '@ttn-lw/components/dropdown'
import ProfilePicture from '@ttn-lw/components/profile-picture'

import Message from '@ttn-lw/lib/components/message'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'

import style from './mobile-menu.styl'

const m = defineMessages({
  loggedInAs: 'Logged in as <b>{userId}</b>',
})

const MobileMenu = ({ className, children, user, onItemsClick, onLogout, bottomArea }) => {
  const handleLogoutClick = useCallback(() => {
    onItemsClick()
    onLogout()
  }, [onItemsClick, onLogout])

  return (
    <div className={classnames(className, style.container)}>
      <Dropdown
        className={style.mobileDropdown}
        itemClassName={style.mobileDropdownItem}
        onItemsClick={onItemsClick}
        larger
      >
        {children}
      </Dropdown>
      <div className={style.bottom}>
        <div className={style.deployment}>{bottomArea}</div>
        {Boolean(user) && (
          <div className={style.userHeader}>
            <div className={style.userMessage}>
              <ProfilePicture
                profilePicture={user.profile_picture}
                className={style.profilePictureMobile}
              />
              <Message
                className={style.userMessage}
                content={m.loggedInAs}
                values={{ userId: user.ids.user_id, b: (...chunks) => <b key="1"> {chunks}</b> }}
              />
            </div>
            <div>
              <Button
                message={sharedMessages.logout}
                icon="logout"
                onClick={handleLogoutClick}
                naked
              />
            </div>
          </div>
        )}
      </div>
    </div>
  )
}

MobileMenu.propTypes = {
  bottomArea: PropTypes.node,
  children: PropTypes.node.isRequired,
  className: PropTypes.string,
  onItemsClick: PropTypes.func.isRequired,
  onLogout: PropTypes.func.isRequired,
  user: PropTypes.user.isRequired,
}

MobileMenu.defaultProps = {
  bottomArea: null,
  className: undefined,
}

export default MobileMenu
