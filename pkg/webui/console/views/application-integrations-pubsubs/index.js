

import React from 'react'
import { Routes, Route, useParams } from 'react-router-dom'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import ErrorView from '@ttn-lw/lib/components/error-view'
import ValidateRouteParam from '@ttn-lw/lib/components/validate-route-param'
import GenericNotFound from '@ttn-lw/lib/components/full-view-error/not-found'

import Require from '@console/lib/components/require'

import ApplicationPubsubEdit from '@console/views/application-integrations-pubsub-edit'
import ApplicationPubsubAdd from '@console/views/application-integrations-pubsub-add'
import ApplicationPubsubsList from '@console/views/application-integrations-pubsubs-list'
import SubViewError from '@console/views/sub-view-error'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import { pathId as pathIdRegexp } from '@ttn-lw/lib/regexp'

import { mayViewApplicationEvents } from '@console/lib/feature-checks'

const ApplicationPubsubs = () => {
  const { appId } = useParams()
  useBreadcrumbs(
    'apps.single.integrations.pubsubs',
    <Breadcrumb
      path={`/applications/${appId}/integrations/pubsubs`}
      content={sharedMessages.pubsubs}
    />,
  )
  return (
    <Require
      featureCheck={mayViewApplicationEvents}
      otherwise={{ redirect: `/applications/${appId}` }}
    >
      <ErrorView errorRender={SubViewError}>
        <Routes>
          <Route index Component={ApplicationPubsubsList} />
          <Route path="add" Component={ApplicationPubsubAdd} />
          <Route
            path=":pubsubId"
            element={
              <ValidateRouteParam
                check={{ pubsubId: pathIdRegexp }}
                Component={ApplicationPubsubEdit}
              />
            }
          />
          <Route path="*" element={<GenericNotFound />} />
        </Routes>
      </ErrorView>
    </Require>
  )
}

export default ApplicationPubsubs
