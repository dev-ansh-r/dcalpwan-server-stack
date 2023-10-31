

import React from 'react'
import classnames from 'classnames'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './cell.styl'

const Cell = ({ className, align, small, width, children, ...rest }) => {
  const cellClassNames = classnames(className, style.cell, {
    [style.cellCentered]: align === 'center',
    [style.cellLeft]: align === 'left',
    [style.cellRight]: align === 'right',
    [style.cellSmall]: small,
  })

  const widthStyle = width ? { width: `${width}%` } : undefined

  return (
    <div {...rest} style={widthStyle} className={cellClassNames} role="cell">
      {children}
    </div>
  )
}

Cell.propTypes = {
  /** A flag indicating how the text in the cell should be aligned. */
  align: PropTypes.oneOf(['left', 'right', 'center']),
  children: PropTypes.node,
  className: PropTypes.string,
  /** The number of columns that the cell should occupy. */
  colSpan: PropTypes.number,
  /** A flag indicating whether the row take less height. */
  small: PropTypes.bool,
  /** The width of the cell in percentages. */
  width: PropTypes.number,
}

Cell.defaultProps = {
  align: undefined,
  children: undefined,
  className: undefined,
  colSpan: 1,
  small: false,
  width: undefined,
}

const HeadCell = ({ className, content, children, ...rest }) => (
  <Cell className={classnames(className, style.cellHead)} {...rest}>
    {Boolean(content) && <Message content={content} />}
    {!Boolean(content) && children}
  </Cell>
)

HeadCell.propTypes = {
  children: PropTypes.node,
  className: PropTypes.string,
  /** The title of the head cell. */
  content: PropTypes.message,
}

HeadCell.defaultProps = {
  children: undefined,
  className: undefined,
  content: undefined,
}

const DataCell = ({ className, children, ...rest }) => (
  <Cell className={classnames(className, style.cellData)} {...rest}>
    {children}
  </Cell>
)

DataCell.propTypes = {
  children: PropTypes.node,
  className: PropTypes.string,
}

DataCell.defaultProps = {
  children: undefined,
  className: undefined,
}

export { Cell, HeadCell, DataCell }
