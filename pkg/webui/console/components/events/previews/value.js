
import React from 'react'

import PropTypes from '@ttn-lw/lib/prop-types'

import DescriptionList from './shared/description-list'
import JSONPayload from './shared/json-payload'

const Value = React.memo(({ event }) => {
  const { data } = event
  return (
    <DescriptionList>
      {Array.isArray(data.value) || typeof data.value === 'object' ? (
        <JSONPayload data={data.value} />
      ) : (
        <DescriptionList.Item data={data.value} />
      )}
    </DescriptionList>
  )
})

Value.propTypes = {
  event: PropTypes.event.isRequired,
}

export default Value
