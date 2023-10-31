

import React from 'react'
import { useSelector } from 'react-redux'
import { Routes, Route, Navigate } from 'react-router-dom'

import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import Tab from '@ttn-lw/components/tabs'

import GenericNotFound from '@ttn-lw/lib/components/full-view-error/not-found'

import DeviceDownlinkPayloadFormatters from '@console/containers/device-payload-formatters/downlink'
import DeviceUplinkPayloadFormatters from '@console/containers/device-payload-formatters/uplink'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { selectSelectedApplicationId } from '@console/store/selectors/applications'
import { selectSelectedDeviceId } from '@console/store/selectors/devices'

import style from './device-payload-formatters.styl'

const DevicePayloadFormatters = () => {
  const appId = useSelector(selectSelectedApplicationId)
  const devId = useSelector(selectSelectedDeviceId)

  useBreadcrumbs(
    'device.single.payload-formatters',
    <Breadcrumb
      path={`/applications/${appId}/devices/${devId}/payload-formatters`}
      content={sharedMessages.payloadFormatters}
    />,
  )

  const tabs = [
    { title: sharedMessages.uplink, name: 'uplink', link: 'uplink' },
    { title: sharedMessages.downlink, name: 'downlink', link: 'downlink' },
  ]

  return (
    <div className={style.fullWidth}>
      <Tab className={style.tabs} tabs={tabs} divider />
      <Routes>
        <Route path="uplink" Component={DeviceUplinkPayloadFormatters} />
        <Route path="downlink" Component={DeviceDownlinkPayloadFormatters} />
        <Route index element={<Navigate to="uplink" replace />} />
        <Route path="*" element={<GenericNotFound />} />
      </Routes>
    </div>
  )
}

export default DevicePayloadFormatters
