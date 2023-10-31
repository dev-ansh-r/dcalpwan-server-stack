

import React from 'react'
import { Col, Row } from 'react-grid-system'
import classnames from 'classnames'

import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'
import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './page-title.styl'

const PageTitle = ({
  title,
  values,
  tall,
  hideHeading,
  children,
  className,
  colProps,
  rowProps,
  noGrid,
}) => {
  const titleClass = classnames(style.title, className, { [style.tall]: tall })
  const pageTitle = <IntlHelmet title={title} values={values} />

  if (noGrid) {
    return (
      <>
        {pageTitle}
        {!hideHeading && (
          <Message component="h1" className={titleClass} content={title} values={values} />
        )}
        {children}
      </>
    )
  }

  return hideHeading ? (
    pageTitle
  ) : (
    <Row {...rowProps}>
      <Col {...colProps}>
        {pageTitle}
        {!hideHeading && (
          <Message component="h1" className={titleClass} content={title} values={values} />
        )}
        {children}
      </Col>
    </Row>
  )
}

PageTitle.propTypes = {
  children: PropTypes.node,
  className: PropTypes.string,
  colProps: PropTypes.shape({}),
  hideHeading: PropTypes.bool,
  noGrid: PropTypes.bool,
  rowProps: PropTypes.shape({}),
  tall: PropTypes.bool,
  title: PropTypes.message.isRequired,
  values: PropTypes.shape({}),
}

PageTitle.defaultProps = {
  children: null,
  className: undefined,
  colProps: {},
  hideHeading: false,
  noGrid: false,
  rowProps: {},
  tall: false,
  values: undefined,
}

export default PageTitle
