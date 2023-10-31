

import React from 'react'
import { defineMessages } from 'react-intl'
import { useParams } from 'react-router-dom'

import PageTitle from '@ttn-lw/components/page-title'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import WithRootClass from '@ttn-lw/lib/components/with-root-class'

import OrganizationEvents from '@console/containers/organization-events'

import style from '@console/views/app/app.styl'

import sharedMessages from '@ttn-lw/lib/shared-messages'

const m = defineMessages({
  orgData: 'Organization data',
})

const Data = () => {
  const { orgId } = useParams()

  useBreadcrumbs(
    'orgs.single.data',
    <Breadcrumb path={`/organizations/${orgId}/data`} content={sharedMessages.liveData} />,
  )

  return (
    <WithRootClass className={style.stageFlex} id="stage">
      <PageTitle hideHeading title={m.orgData} />
      <OrganizationEvents orgId={orgId} />
    </WithRootClass>
  )
}

export default Data
