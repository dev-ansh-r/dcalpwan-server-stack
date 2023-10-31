

import React from 'react'
import classnames from 'classnames'

import Spinner from '@ttn-lw/components/spinner'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import style from './content.styl'

const EntityTitleSectionContent = props => {
  const { className, children, bottomBarLeft, bottomBarRight, fetching } = props

  return (
    <div className={classnames(className, style.container)}>
      {fetching ? (
        <Spinner after={0} faded micro inline>
          <Message content={sharedMessages.fetching} />
        </Spinner>
      ) : (
        <div className={style.content}>
          {Boolean(bottomBarLeft || bottomBarRight) && (
            <div className={style.bottomBar}>
              <div>{bottomBarLeft}</div>
              <div>{bottomBarRight}</div>
            </div>
          )}
          {children}
        </div>
      )}
    </div>
  )
}

EntityTitleSectionContent.propTypes = {
  bottomBarLeft: PropTypes.node,
  bottomBarRight: PropTypes.node,
  children: PropTypes.oneOfType([PropTypes.arrayOf(PropTypes.node), PropTypes.node]),
  className: PropTypes.string,
  fetching: PropTypes.bool,
}

EntityTitleSectionContent.defaultProps = {
  bottomBarLeft: undefined,
  bottomBarRight: undefined,
  className: undefined,
  children: undefined,
  fetching: false,
}

export default EntityTitleSectionContent
