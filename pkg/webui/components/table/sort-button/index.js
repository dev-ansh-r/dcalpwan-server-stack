

import React, { useCallback } from 'react'
import classnames from 'classnames'

import Icon from '@ttn-lw/components/icon'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './sort-button.styl'

const SortButton = ({ name, onSort, className, active, direction, title }) => {
  const handleSort = useCallback(() => {
    onSort(name)
  }, [name, onSort])

  const buttonClassNames = classnames(className, style.button, {
    [style.buttonActive]: active,
    [style.buttonDesc]: active && direction === 'desc',
  })

  let icon = 'sort_order'
  if (active && direction) {
    icon += `_${direction}`
  }

  return (
    <button className={buttonClassNames} type="button" onClick={handleSort}>
      <Message content={title} />
      <Icon className={style.icon} icon={icon} nudgeUp />
    </button>
  )
}

SortButton.propTypes = {
  /** A flag identifying whether the button is active. */
  active: PropTypes.bool.isRequired,
  className: PropTypes.string,
  /** The current ordering (ascending/descending/none). */
  direction: PropTypes.string,
  /** The name of the column that the sort button represents. */
  name: PropTypes.string.isRequired,
  /** Function to be called when the button is pressed. */
  onSort: PropTypes.func.isRequired,
  /** The text of the button. */
  title: PropTypes.message.isRequired,
}

SortButton.defaultProps = {
  className: undefined,
  direction: undefined,
}

export default SortButton
