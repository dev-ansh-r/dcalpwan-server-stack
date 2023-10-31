

import React, { useCallback } from 'react'
import PropTypes from 'prop-types'
import classnames from 'classnames'
import Paginate from 'react-paginate'

import Icon from '@ttn-lw/components/icon'

import style from './pagination.styl'

const Pagination = ({
  onPageChange,
  className,
  forcePage,
  pageRangeDisplayed,
  marginPagesDisplayed,
  hideIfOnlyOnePage,
  pageCount,
  ...rest
}) => {
  const handlePageChange = useCallback(
    page => {
      onPageChange(page.selected + 1)
    },
    [onPageChange],
  )

  // Don't show pagination if there is only one page.
  if (hideIfOnlyOnePage && pageCount === 1) {
    return null
  }

  const containerClassNames = classnames(style.pagination, className)
  const breakClassNames = classnames(style.item, style.itemBreak)
  const navigationNextClassNames = classnames(style.item, style.itemNavigationNext)
  const navigationPrevClassNames = classnames(style.item, style.itemNavigationPrev)

  return (
    <Paginate
      previousClassName={navigationPrevClassNames}
      previousLinkClassName={style.link}
      previousLabel={<Icon icon="navigate_before" small aria-label="Go to the previous page" />}
      nextClassName={navigationNextClassNames}
      nextLinkClassName={style.link}
      nextLabel={<Icon icon="navigate_next" small aria-label="Go to the next page" />}
      containerClassName={containerClassNames}
      pageClassName={style.item}
      breakClassName={breakClassNames}
      pageLinkClassName={style.link}
      disabledClassName={style.itemDisabled}
      activeClassName={style.itemActive}
      forcePage={forcePage - 1}
      pageRangeDisplayed={pageRangeDisplayed}
      marginPagesDisplayed={marginPagesDisplayed}
      onPageChange={handlePageChange}
      pageCount={pageCount}
      {...rest}
    />
  )
}

Pagination.propTypes = {
  className: PropTypes.string,
  /** Page to be displayed immediately. */
  forcePage: PropTypes.number,
  /** A flag indicating whether the pagination should be hidden when there is
   * only one page.
   */
  hideIfOnlyOnePage: PropTypes.bool,
  /**
   * The number of pages to be displayed in the beginning/end of
   * the component. For example, marginPagesDisplayed = 2, then the
   * component will display at most two pages as margins:
   * [<][1][2]...[10]...[19][20][>].
   *
   */
  marginPagesDisplayed: PropTypes.number,
  /** An onClick handler that gets called with the new page number. */
  onPageChange: PropTypes.func,
  /** The total number of pages. */
  pageCount: PropTypes.number.isRequired,
  /**
   * The number of pages to be displayed. If is bigger than
   * pageCount, then all pages will be displayed without gaps.
   */
  pageRangeDisplayed: PropTypes.number,
}

Pagination.defaultProps = {
  className: undefined,
  forcePage: 1,
  hideIfOnlyOnePage: true,
  marginPagesDisplayed: 1,
  onPageChange: () => null,
  pageRangeDisplayed: 1,
}

export default Pagination
