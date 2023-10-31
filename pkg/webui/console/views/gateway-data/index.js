

import React from 'react'
import { defineMessages } from 'react-intl'
import { useParams } from 'react-router-dom'

import PageTitle from '@ttn-lw/components/page-title'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import WithRootClass from '@ttn-lw/lib/components/with-root-class'

import GatewayEvents from '@console/containers/gateway-events'

import Require from '@console/lib/components/require'

import style from '@console/views/app/app.styl'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { mayViewGatewayEvents } from '@console/lib/feature-checks'

const m = defineMessages({
  gtwData: 'Gateway data',
})

const GatewayData = () => {
  const { gtwId } = useParams()

  useBreadcrumbs(
    'gtws.single.data',
    <Breadcrumb path={`/gateways/${gtwId}/data`} content={sharedMessages.liveData} />,
  )

  return (
    <Require featureCheck={mayViewGatewayEvents} otherwise={{ redirect: `/gateways/${gtwId}` }}>
      <WithRootClass className={style.stageFlex} id="stage">
        <PageTitle hideHeading title={m.gtwData} />
        <GatewayEvents gtwId={gtwId} />
      </WithRootClass>
    </Require>
  )
}

export default GatewayData
