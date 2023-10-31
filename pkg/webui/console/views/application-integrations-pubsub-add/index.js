

import React, { useCallback } from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { useSelector, useDispatch } from 'react-redux'
import { useNavigate, useParams } from 'react-router-dom'

import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import PageTitle from '@ttn-lw/components/page-title'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'

import PubsubForm from '@console/components/pubsub-form'

import { isNotFoundError } from '@ttn-lw/lib/errors/utils'
import sharedMessages from '@ttn-lw/lib/shared-messages'
import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'

import { createPubsub, getPubsub } from '@console/store/actions/pubsubs'

import {
  selectMqttProviderDisabled,
  selectNatsProviderDisabled,
} from '@console/store/selectors/application-server'

// Inner Function
const ApplicationPubsubAdd = () => {
  const { appId } = useParams()
  const mqttDisabled = useSelector(selectMqttProviderDisabled)
  const natsDisabled = useSelector(selectNatsProviderDisabled)
  const dispatch = useDispatch()
  const navigate = useNavigate()

  useBreadcrumbs(
    'apps.single.integrations.add',
    <Breadcrumb path={`/applications/${appId}/integrations/add`} content={sharedMessages.add} />,
  )

  const existCheck = useCallback(
    async pubsubId => {
      try {
        await dispatch(attachPromise(getPubsub(appId, pubsubId, [])))
        return true
      } catch (error) {
        if (isNotFoundError(error)) {
          return false
        }
        throw error
      }
    },
    [appId, dispatch],
  )

  const handleSubmit = useCallback(
    async pubsub => {
      await dispatch(attachPromise(createPubsub(appId, pubsub)))
      navigate(`/applications/${appId}/integrations/pubsubs`)
    },
    [appId, dispatch, navigate],
  )

  return (
    <Container>
      <PageTitle title={sharedMessages.addPubsub} className="mb-0" />
      <Row>
        <Col lg={8} md={12}>
          <PubsubForm
            appId={appId}
            update={false}
            onSubmit={handleSubmit}
            existCheck={existCheck}
            mqttDisabled={mqttDisabled}
            natsDisabled={natsDisabled}
          />
        </Col>
      </Row>
    </Container>
  )
}

export default ApplicationPubsubAdd
