

import React from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { Routes, Route, Navigate, useParams } from 'react-router-dom'
import { useSelector } from 'react-redux'

import Tabs from '@ttn-lw/components/tabs'

import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'

import DownlinkForm from '@console/components/downlink-form'
import UplinkForm from '@console/components/uplink-form'

import Require from '@console/lib/components/require'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import {
  mayWriteTraffic,
  mayScheduleDownlinks,
  maySendUplink,
  checkFromState,
} from '@console/lib/feature-checks'

import style from './device-messaging.styl'

const DeviceMessaging = () => {
  const { appId, devId } = useParams()
  const mayScheduleDown = useSelector(state => checkFromState(mayScheduleDownlinks, state))
  const maySendUp = useSelector(state => checkFromState(maySendUplink, state))
  const tabs =
    mayScheduleDown && maySendUp
      ? [
          { title: sharedMessages.uplink, name: 'uplink', link: 'uplink' },
          { title: sharedMessages.downlink, name: 'downlink', link: 'downlink' },
        ]
      : []

  return (
    <Require
      featureCheck={mayWriteTraffic}
      otherwise={{ redirect: `/applications/${appId}/devices/${devId}` }}
    >
      <Container>
        <IntlHelmet title={sharedMessages.messaging} />
        <Row>
          {tabs.length > 0 && (
            <Col sm={12}>
              <Tabs className={style.tabs} tabs={tabs} divider />
            </Col>
          )}
          <Col lg={8} md={12}>
            <Routes>
              {maySendUp && <Route path="uplink" Component={UplinkForm} />}
              {mayScheduleDown && <Route path="downlink" Component={DownlinkForm} />}
              <Route index element={<Navigate to="uplink" replace />} />
            </Routes>
          </Col>
        </Row>
      </Container>
    </Require>
  )
}

export default DeviceMessaging
