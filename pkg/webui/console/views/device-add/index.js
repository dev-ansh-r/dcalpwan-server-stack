

import React, { useCallback } from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { useParams } from 'react-router-dom'

import PageTitle from '@ttn-lw/components/page-title'

import RequireRequest from '@ttn-lw/lib/components/require-request'

import DeviceOnboardingForm from '@console/containers/device-onboarding-form'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import { selectJsConfig } from '@ttn-lw/lib/selectors/env'
import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'

import { listBrands } from '@console/store/actions/device-repository'
import { getJoinEUIPrefixes } from '@console/store/actions/join-server'

const DeviceAdd = () => {
  const { appId } = useParams()
  const { enabled: jsEnabled } = selectJsConfig()
  const requestAction = useCallback(
    async dispatch => {
      if (jsEnabled) {
        await dispatch(attachPromise(getJoinEUIPrefixes()))
      }
      await dispatch(attachPromise(listBrands(appId, {}, ['name', 'lora_alliance_vendor_id'])))
    },
    [appId, jsEnabled],
  )

  return (
    <RequireRequest requestAction={requestAction}>
      <Container>
        <Row>
          <Col>
            <PageTitle tall title={sharedMessages.registerEndDevice} className="mb-cs-m" />
            <DeviceOnboardingForm />
          </Col>
        </Row>
      </Container>
    </RequireRequest>
  )
}

export default DeviceAdd
