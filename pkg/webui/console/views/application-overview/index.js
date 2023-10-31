

import React from 'react'
import { Col, Row, Container } from 'react-grid-system'
import { defineMessages } from 'react-intl'
import { useParams } from 'react-router-dom'
import { useSelector } from 'react-redux'

import PAGE_SIZES from '@ttn-lw/constants/page-sizes'

import DataSheet from '@ttn-lw/components/data-sheet'

import DateTime from '@ttn-lw/lib/components/date-time'
import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'

import DevicesTable from '@console/containers/devices-table'
import ApplicationEvents from '@console/containers/application-events'
import ApplicationTitleSection from '@console/containers/application-title-section'

import Require from '@console/lib/components/require'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { isOtherClusterApp } from '@console/lib/application-utils'
import { mayViewApplicationInfo } from '@console/lib/feature-checks'
import { checkFromState } from '@account/lib/feature-checks'

import { selectSelectedApplication } from '@console/store/selectors/applications'

import style from './application-overview.styl'

const m = defineMessages({
  failedAccessOtherHostApplication:
    'The application you attempted to visit is registered on a different cluster and needs to be accessed using its host Console.',
})

const ApplicationOverview = () => {
  const { appId } = useParams()
  const application = useSelector(selectSelectedApplication)
  const may = useSelector(state => checkFromState(mayViewApplicationInfo, state))
  const { created_at, updated_at } = application
  const shouldRedirect = isOtherClusterApp(application)
  const condition = !shouldRedirect && may

  const sheetData = [
    {
      header: sharedMessages.generalInformation,
      items: [
        { key: sharedMessages.appId, value: appId, type: 'code', sensitive: false },
        { key: sharedMessages.createdAt, value: <DateTime value={created_at} /> },
        { key: sharedMessages.updatedAt, value: <DateTime value={updated_at} /> },
      ],
    },
  ]

  const otherwise = {
    redirect: '/applications',
    message: m.failedAccessOtherHostApplication,
  }

  return (
    <Require condition={condition} otherwise={otherwise}>
      <div className={style.titleSection}>
        <Container>
          <IntlHelmet title={sharedMessages.overview} />
          <Row>
            <Col sm={12}>
              <ApplicationTitleSection appId={appId} />
            </Col>
          </Row>
        </Container>
      </div>
      <Container>
        <Row>
          <Col sm={12} lg={6}>
            <DataSheet data={sheetData} className={style.generalInformation} />
          </Col>
          <Col sm={12} lg={6}>
            <ApplicationEvents appId={appId} widget />
          </Col>
        </Row>
        <Row>
          <Col sm={12} className={style.table}>
            <DevicesTable pageSize={PAGE_SIZES.SMALL} devicePathPrefix="/devices" />
          </Col>
        </Row>
      </Container>
    </Require>
  )
}

export default ApplicationOverview
