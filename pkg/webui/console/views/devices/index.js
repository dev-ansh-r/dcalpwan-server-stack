

import React from 'react'
import { Route, Routes } from 'react-router-dom'
import { useSelector } from 'react-redux'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import ValidateRouteParam from '@ttn-lw/lib/components/validate-route-param'
import GenericNotFound from '@ttn-lw/lib/components/full-view-error/not-found'

import Require from '@console/lib/components/require'

import Device from '@console/views/device'
import DeviceImport from '@console/views/device-import'
import DeviceAdd from '@console/views/device-add'
import DeviceList from '@console/views/device-list'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import { pathId as pathIdRegexp } from '@ttn-lw/lib/regexp'

import { mayViewApplicationDevices } from '@console/lib/feature-checks'

import { selectSelectedApplicationId } from '@console/store/selectors/applications'

const Devices = () => {
  const appId = useSelector(selectSelectedApplicationId)

  useBreadcrumbs(
    'apps.single.devices',
    <Breadcrumb path={`/applications/${appId}/devices`} content={sharedMessages.devices} />,
  )

  return (
    <Require
      featureCheck={mayViewApplicationDevices}
      otherwise={{ redirect: `/applications/${appId}` }}
    >
      <Routes>
        <Route index Component={DeviceList} />
        <Route path="add" Component={DeviceAdd} />
        <Route path="import" Component={DeviceImport} />
        <Route
          path=":devId/*"
          element={<ValidateRouteParam check={{ devId: pathIdRegexp }} Component={Device} />}
        />
        <Route path="*" element={<GenericNotFound />} />
      </Routes>
    </Require>
  )
}

export default Devices
