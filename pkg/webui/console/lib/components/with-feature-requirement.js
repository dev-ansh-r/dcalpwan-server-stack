

import React from 'react'

import Require from './require'

/**
 * `withFeatureRequirement` is a HOC that checks whether the current has the
 * necessary authorization to view the wrapped component. It can be set up to
 * either redirect to another route, to render something different or to not
 * render anything if the requirement is not met.
 *
 * @param {object} featureCheck - The feature check object containing the right
 * selector as well as the check itself.
 * @param {object} otherwise - A configuration object determining what should be
 * rendered if the requirement was not met. If not set, nothing will be
 * rendered.
 * @returns {Function} - An instance of the `withFeatureRequirement` HOC.
 */
const withFeatureRequirement = (featureCheck, otherwise) => Component =>
  class WithFeatureRequirement extends React.Component {
    constructor(props) {
      super(props)

      if (
        typeof otherwise === 'object' &&
        'redirect' in otherwise &&
        typeof otherwise.redirect === 'function'
      ) {
        this.otherwise = { ...otherwise, redirect: otherwise.redirect(props) }
      } else {
        this.otherwise = otherwise
      }
    }

    render() {
      return (
        <Require featureCheck={featureCheck} otherwise={this.otherwise}>
          <Component {...this.props} />
        </Require>
      )
    }
  }

export default withFeatureRequirement
