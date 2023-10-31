

import React from 'react'
import { useSelector } from 'react-redux'
import { defineMessages } from 'react-intl'

import Button from '@ttn-lw/components/button'
import ProfilePicture from '@ttn-lw/components/profile-picture'

import {
  selectUserName,
  selectUserId,
  selectUserProfilePicture,
} from '@account/store/selectors/user'

import style from './profile-card.styl'

const m = defineMessages({
  editProfileSettings: 'Edit profile settings',
})

const ProfileCard = () => {
  const userId = useSelector(selectUserId)
  const userName = useSelector(selectUserName)
  const profilePicture = useSelector(selectUserProfilePicture)

  return (
    <section className={style.container} data-test-id="profile-card">
      <ProfilePicture profilePicture={profilePicture} className={style.profilePicture} />
      <div className={style.panel}>
        <div className={style.name}>
          <h3>{userName || userId}</h3>
          {Boolean(userName) && <span className={style.userId}>{userId}</span>}
        </div>
        <Button.Link to="/profile-settings" icon="edit" message={m.editProfileSettings} />
      </div>
    </section>
  )
}

export default ProfileCard
