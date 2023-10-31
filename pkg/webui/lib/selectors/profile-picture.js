

export const isValidProfilePictureObject = profilePicture =>
  Boolean(profilePicture) && Boolean(profilePicture.sizes)

export const getProfilePictureSizes = profilePicture => {
  if (!isValidProfilePictureObject(profilePicture)) {
    return {}
  }

  return profilePicture.sizes
}

export const getSmallestAvailableProfilePicture = profilePicture =>
  getClosestProfilePictureBySize(profilePicture, 64)

export const getOriginalSizeProfilePicture = profilePicture => {
  const sizes = getProfilePictureSizes(profilePicture)

  return sizes[0]
}

export const getClosestProfilePictureBySize = (profilePicture, size) => {
  const sizes = getProfilePictureSizes(profilePicture)
  if (sizes[size]) {
    return sizes[size]
  }

  const closestSize = Object.keys(sizes).sort((a, b) => Math.abs(size - a) - Math.abs(size - b))[0]

  return sizes[closestSize]
}

export const isGravatarProfilePicture = profilePicture => {
  if (!isValidProfilePictureObject(profilePicture)) return false
  const sizes = profilePicture.sizes

  return sizes[64] ? sizes[64].startsWith('https://www.gravatar.com') : false
}

export const convertUriToProfilePicture = uri => ({ sizes: { 0: uri } })
