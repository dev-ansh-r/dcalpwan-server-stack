
import React from 'react'
import { useLocation } from 'react-router-dom'

import { ingestError } from '@ttn-lw/lib/errors/utils'
import PropTypes from '@ttn-lw/lib/prop-types'

class ErrorView extends React.Component {
  static propTypes = {
    children: PropTypes.node.isRequired,
    errorRender: PropTypes.func.isRequired,
    location: PropTypes.location,
  }

  static defaultProps = {
    location: undefined,
  }

  state = {
    error: undefined,
    hasCaught: false,
  }

  unlisten = () => null

  componentWillUnmount() {
    this.unlisten()
  }

  componentDidUpdate({ location: prevLocation }) {
    const { location } = this.props
    if (location !== prevLocation) {
      this.setState({ hasCaught: false, error: undefined })
    }
  }

  componentDidCatch(error) {
    ingestError(error, { ingestedBy: 'ErrorView' })

    this.setState({
      hasCaught: true,
      error,
    })
  }

  render() {
    const { children, errorRender } = this.props
    const { hasCaught, error } = this.state

    if (hasCaught) {
      return errorRender(error)
    }

    return React.Children.only(children)
  }
}

const ErrorViewWithRouter = props => {
  const location = useLocation()
  return <ErrorView {...props} location={location} />
}

export { ErrorViewWithRouter as default, ErrorView }
