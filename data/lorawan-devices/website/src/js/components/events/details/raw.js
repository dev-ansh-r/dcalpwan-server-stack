
import React from 'react'
import PropTypes from 'prop-types'

import CodeEditor from '../../code-editor'

const RawEventDetails = React.memo((props) => {
  const { className, id, details } = props

  const formattedDetails = JSON.stringify(details, null, 2)

  return (
    <CodeEditor className={className} readOnly name={id} language="json" value={formattedDetails} />
  )
})

RawEventDetails.propTypes = {
  className: PropTypes.string,
  details: PropTypes.shape({}).isRequired,
  id: PropTypes.string.isRequired,
}

RawEventDetails.defaultProps = {
  className: undefined,
}

export default RawEventDetails
