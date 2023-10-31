

import React from 'react'
import ReactDom from 'react-dom'
import classnames from 'classnames'
import { Container } from 'react-grid-system'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './breadcrumbs.styl'

const Breadcrumbs = ({ className, breadcrumbs }) => (
  <nav className={classnames(className, style.breadcrumbs)}>
    {breadcrumbs.map((component, index) =>
      React.cloneElement(component, {
        key: index,
        isLast: index === breadcrumbs.length - 1,
      }),
    )}
  </nav>
)

Breadcrumbs.propTypes = {
  /** A list of breadcrumb elements. */
  breadcrumbs: PropTypes.arrayOf(PropTypes.oneOfType([PropTypes.func, PropTypes.element])),
  className: PropTypes.string,
}

Breadcrumbs.defaultProps = {
  breadcrumbs: undefined,
  className: undefined,
}

const PortalledBreadcrumbs = ({ className, ...rest }) => {
  // Breadcrumbs can be rendered into multiple containers.
  const containers = document.querySelectorAll('.breadcrumbs, #breadcrumbs')
  if (containers.length) {
    const nodes = []
    containers.forEach(element => {
      nodes.push(
        ReactDom.createPortal(
          <div className={classnames(className, style.breadcrumbsContainer)}>
            <Container>
              <Breadcrumbs {...rest} />
            </Container>
          </div>,
          element,
        ),
      )
    })
    return nodes
  }

  return null
}

PortalledBreadcrumbs.propTypes = {
  className: PropTypes.string,
}

PortalledBreadcrumbs.defaultProps = {
  className: undefined,
}

export { PortalledBreadcrumbs as default, Breadcrumbs }
