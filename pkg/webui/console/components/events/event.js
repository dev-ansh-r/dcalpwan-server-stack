

import React, { useMemo, useCallback } from 'react'
import classnames from 'classnames'
import { useIntl } from 'react-intl'
import { upperFirst } from 'lodash'

import Icon from '@ttn-lw/components/icon'

import DateTime from '@ttn-lw/lib/components/date-time'

import PropTypes from '@ttn-lw/lib/prop-types'

import { eventMessages } from '@console/lib/events/definitions'

import { getEventIconByName, getPreviewComponent, getEntityId } from './utils'
import ErrorBoundary from './error-boundary'

import style from './events.styl'

const Event = React.memo(({ event, scoped, widget, rowStyle, onRowClick, eventId, active }) => {
  const intl = useIntl()

  const icon = useMemo(() => getEventIconByName(event.name), [event.name])
  const typeValue = event.isSynthetic
    ? intl.formatMessage(eventMessages[`${event.name}:type`])
    : upperFirst(intl.formatMessage({ id: `event:${event.name}` }))
  const PreviewComponent = useMemo(() => getPreviewComponent(event), [event])
  const entityId = useMemo(
    () => (!event.isSynthetic && !scoped ? getEntityId(event.identifiers[0]) : undefined),
    [scoped, event.identifiers, event.isSynthetic],
  )

  const handleRowClick = useCallback(() => {
    onRowClick(eventId)
  }, [eventId, onRowClick])
  const eventClasses = classnames(style.event, {
    [style.widget]: widget,
    [style.active]: active,
    [style.synthetic]: event.isSynthetic,
  })
  return (
    <li className={eventClasses} style={rowStyle} onClick={handleRowClick}>
      <ErrorBoundary>
        <div className={style.cellTime} title={`${event.time}: ${typeValue}`}>
          <Icon icon={icon} className={style.eventIcon} />
          <div>
            <DateTime value={event.time} date={false} />
          </div>
        </div>
        {!scoped && (
          <div className={style.cellId} title={entityId}>
            <span>{entityId}</span>
          </div>
        )}
        <div className={style.cellType} title={typeValue}>
          <span>{typeValue}</span>
        </div>
        {(!widget || (widget && scoped)) && (
          <div className={style.cellPreview}>
            <PreviewComponent event={event} />
          </div>
        )}
      </ErrorBoundary>
    </li>
  )
})

Event.propTypes = {
  active: PropTypes.bool,
  event: PropTypes.event.isRequired,
  eventId: PropTypes.string.isRequired,
  onRowClick: PropTypes.func,
  rowStyle: PropTypes.shape({}),
  scoped: PropTypes.bool,
  widget: PropTypes.bool,
}

Event.defaultProps = {
  active: false,
  onRowClick: () => null,
  rowStyle: undefined,
  scoped: false,
  widget: false,
}

export default Event
