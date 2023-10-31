

import React from 'react'
import { useSelector } from 'react-redux'
import { useParams } from 'react-router-dom'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import GenericNotFound from '@ttn-lw/lib/components/full-view-error/not-found'
import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'
import WithRootClass from '@ttn-lw/lib/components/with-root-class'

import DeviceEvents from '@console/containers/device-events'

import appStyle from '@console/views/app/app.styl'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { selectSelectedDevice, selectSelectedDeviceId } from '@console/store/selectors/devices'

import style from './device-data.styl'

const Data = () => {
  const { appId } = useParams()

  const device = useSelector(selectSelectedDevice)
  const devId = useSelector(selectSelectedDeviceId)

  useBreadcrumbs(
    'device.single.data',
    <Breadcrumb
      path={`/applications/${appId}/devices/${devId}/data`}
      content={sharedMessages.liveData}
    />,
  )

  if (!device) {
    return <GenericNotFound />
  }

  const { ids } = device

  return (
    <WithRootClass className={appStyle.stageFlex} id="stage">
      <div className={style.overflowContainer}>
        <IntlHelmet hideHeading title={sharedMessages.liveData} />
        <DeviceEvents devIds={ids} />
      </div>
    </WithRootClass>
  )
}

export default Data
