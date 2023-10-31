

export const composeOption = value => {
  const data = value.ids || value
  return {
    value: 'user_ids' in data ? data.user_ids?.user_id : data.organization_ids?.organization_id,
    label: 'user_ids' in data ? data.user_ids?.user_id : data.organization_ids?.organization_id,
    icon: 'user_ids' in data ? 'user' : 'organization',
  }
}

export const encodeContact = value =>
  value
    ? {
        [`${value.icon}_ids`]: {
          [`${value.icon}_id`]: value.value,
        },
      }
    : null

export const decodeContact = value => (value ? composeOption(value) : null)
