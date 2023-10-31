

import React from 'react'
import { Routes, Route, Navigate, useParams } from 'react-router-dom'
import { Container, Col, Row } from 'react-grid-system'
import { useSelector } from 'react-redux'

import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'

import RequireRequest from '@ttn-lw/lib/components/require-request'

import ApplicationUplinkPayloadFormatters from '@console/containers/application-payload-formatters/uplink'
import ApplicationDownlinkPayloadFormatters from '@console/containers/application-payload-formatters/downlink'

import Require from '@console/lib/components/require'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import {
  maySetApplicationPayloadFormatters,
  mayViewApplicationLink,
} from '@console/lib/feature-checks'
import { checkFromState } from '@account/lib/feature-checks'

import { getApplicationLink } from '@console/store/actions/link'

const ApplicationPayloadFormatters = () => {
  const { appId } = useParams()
  const mayViewLink = useSelector(state => checkFromState(mayViewApplicationLink, state))
  const getLink = mayViewLink ? getApplicationLink(appId, ['default_formatters']) : []

  return (
    <Require
      featureCheck={maySetApplicationPayloadFormatters}
      otherwise={{ redirect: `/applications/${appId}` }}
    >
      <RequireRequest requestAction={getLink}>
        <ApplicationPayloadFormattersInner />
      </RequireRequest>
    </Require>
  )
}

const ApplicationPayloadFormattersInner = () => {
  const { appId } = useParams()

  useBreadcrumbs(
    'apps.single.payload-formatters',
    <Breadcrumb
      path={`/applications/${appId}/payload-formatters`}
      content={sharedMessages.payloadFormatters}
    />,
  )

  return (
    <Container>
      <Row>
        <Col>
          <Routes>
            <Route index element={<Navigate to="uplink" replace />} />
            <Route path="uplink" Component={ApplicationUplinkPayloadFormatters} />
            <Route path="downlink" Component={ApplicationDownlinkPayloadFormatters} />
          </Routes>
        </Col>
      </Row>
    </Container>
  )
}

export default ApplicationPayloadFormatters
