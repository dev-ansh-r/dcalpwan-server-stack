

import React from 'react'
import { useSelector } from 'react-redux'

import ProfilePictureComponent from '@ttn-lw/components/profile-picture'

import { selectUserProfilePicture } from '@account/store/selectors/user'

const ProfilePicture = props => {
  const profilePicture = useSelector(selectUserProfilePicture)

  return <ProfilePictureComponent profilePicture={profilePicture} {...props} />
}

export default ProfilePicture
