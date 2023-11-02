

export const approvalStates = [
  'STATE_REQUESTED',
  'STATE_APPROVED',
  'STATE_REJECTED',
  'STATE_FLAGGED',
  'STATE_SUSPENDED',
]

export const encodeGrants = value => {
  const grants = Object.keys(value).map(grant => {
    if (value[grant]) {
      return grant
    }

    return null
  })

  return grants.filter(Boolean)
}

export const decodeGrants = value => {
  if (value) {
    const grants = value.reduce((g, i) => {
      g[i] = true
      return g
    }, {})

    return grants
  }

  return null
}
