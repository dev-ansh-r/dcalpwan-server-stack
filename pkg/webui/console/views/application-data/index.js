

import React from 'react'
import { useParams } from 'react-router-dom'

import PageTitle from '@ttn-lw/components/page-title'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import ApplicationEvents from '@console/containers/application-events'

import Require from '@console/lib/components/require'

import style from '@console/views/app/app.styl'

import useRootClass from '@ttn-lw/lib/hooks/use-root-class'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import { mayViewApplicationEvents } from '@console/lib/feature-checks'

const ApplicationData = () => {
  const { appId } = useParams()

  useRootClass(style.stageFlex, 'stage')

  useBreadcrumbs(
    'apps.single.data',
    <Breadcrumb path={`/applications/${appId}/data`} content={sharedMessages.appData} />,
  )

  return (
    <Require
      featureCheck={mayViewApplicationEvents}
      otherwise={{ redirect: `/applications/${appId}` }}
    >
      <PageTitle hideHeading title={sharedMessages.appData} />
      <ApplicationEvents appId={appId} />
    </Require>
  )
}

export default ApplicationData
