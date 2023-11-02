
import React, { useRef, useCallback } from 'react'
import classnames from 'classnames'

import missingProfilePicture from '@assets/img/placeholder/missing-profile-picture.png'

import PropTypes from '@ttn-lw/lib/prop-types'
import {
  getClosestProfilePictureBySize,
  isValidProfilePictureObject,
} from '@ttn-lw/lib/selectors/profile-picture'

import styles from './profile-picture.styl'

const ProfilePicture = ({ profilePicture, className, size }) => {
  const imageRef = useRef()
  const handleImageError = useCallback(error => {
    error.target.src = missingProfilePicture
  }, [])
  return (
    <div className={classnames(className, styles.container)}>
      <img
        onError={handleImageError}
        src={
          isValidProfilePictureObject(profilePicture)
            ? getClosestProfilePictureBySize(profilePicture, size)
            : missingProfilePicture
        }
        alt="Profile picture"
        ref={imageRef}
      />
    </div>
  )
}

ProfilePicture.propTypes = {
  className: PropTypes.string,
  profilePicture: PropTypes.profilePicture,
  size: PropTypes.number,
}

ProfilePicture.defaultProps = {
  profilePicture: undefined,
  className: undefined,
  size: 128,
}

export default ProfilePicture
