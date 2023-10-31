

import PropTypes from '@ttn-lw/lib/prop-types'

import useRootClass from '../hooks/use-root-class'

const WithRootClass = ({ children, className, id }) => {
  useRootClass(className, id)

  return children
}

WithRootClass.propTypes = {
  children: PropTypes.node.isRequired,
  className: PropTypes.string,
  id: PropTypes.string,
}

WithRootClass.defaultProps = {
  className: undefined,
  id: 'app',
}

export default WithRootClass
