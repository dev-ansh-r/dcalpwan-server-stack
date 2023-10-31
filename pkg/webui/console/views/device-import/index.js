

import React from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { useSelector } from 'react-redux'
import { defineMessages } from 'react-intl'

import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import Notification from '@ttn-lw/components/notification'
import PageTitle from '@ttn-lw/components/page-title'

import RequireRequest from '@ttn-lw/lib/components/require-request'

import DeviceImporter from '@console/containers/device-importer'

import SubViewError from '@console/views/sub-view-error'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { listBrands } from '@console/store/actions/device-repository'

import {
  selectDeviceTemplateFormats,
  selectDeviceTemplateFormatsFetching,
} from '@console/store/selectors/device-template-formats'
import { selectSelectedApplicationId } from '@console/store/selectors/applications'

const m = defineMessages({
  noTemplatesTitle: 'No end device templates found',
  noTemplates:
    'There are currently no end device templates set up. Please set up an end device template to make use of the bulk device import feature. For more information please refer to the documentation.',
})

const DeviceAddBulk = () => {
  const appId = useSelector(selectSelectedApplicationId)
  useBreadcrumbs(
    'devices.import',
    <Breadcrumb path={`/applications/${appId}/devices/import`} content={sharedMessages.import} />,
  )

  const deviceTemplateFormats = useSelector(selectDeviceTemplateFormats)
  const deviceTemplateFormatsFetching = useSelector(selectDeviceTemplateFormatsFetching)
  const showEmptyWarning =
    !deviceTemplateFormatsFetching && Object.keys(deviceTemplateFormats).length === 0

  return (
    <RequireRequest
      requestAction={listBrands(appId, {}, ['name', 'lora_alliance_vendor_id'])}
      errorRenderFunction={SubViewError}
      spinnerProps={{ center: false, inline: true }}
    >
      <Container>
        <PageTitle title={sharedMessages.importDevices} />
        <Row>
          <Col md={12}>
            {showEmptyWarning && (
              <Notification warning title={m.noTemplatesTitle} content={m.noTemplates} />
            )}
            <DeviceImporter />
          </Col>
        </Row>
      </Container>
    </RequireRequest>
  )
}

export default DeviceAddBulk
