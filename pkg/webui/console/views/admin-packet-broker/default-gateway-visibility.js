

import React, { useCallback, useState } from 'react'
import { Col } from 'react-grid-system'
import { useSelector, useDispatch } from 'react-redux'

import toast from '@ttn-lw/components/toast'

import Message from '@ttn-lw/lib/components/message'

import GatewayVisibilityForm from '@console/components/gateway-visibility-form'

import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'

import { setHomeNetworkDefaultGatewayVisibility } from '@console/store/actions/packet-broker'

import { selectHomeNetworkDefaultGatewayVisibility } from '@console/store/selectors/packet-broker'

import m from './messages'

import style from './admin-packet-broker.styl'

const DefaultGatewayVisibilityView = () => {
  const dispatch = useDispatch()
  const defaultGatewayVisibility = useSelector(selectHomeNetworkDefaultGatewayVisibility)
  const initialValues = {
    visibility: defaultGatewayVisibility.visibility || {},
  }
  const [formError, setFormError] = useState(undefined)
  const handleDefaultGatewayVisibilitySubmit = useCallback(
    async ({ visibility }) => {
      try {
        await dispatch(attachPromise(setHomeNetworkDefaultGatewayVisibility(visibility)))
        toast({
          message: m.defaultGatewayVisibilitySet,
          type: toast.types.SUCCESS,
        })
      } catch (error) {
        setFormError(error)
      }
    },
    [dispatch, setFormError],
  )

  return (
    <Col md={12}>
      <Message
        content={m.gatewayVisibilityInformation}
        component="p"
        className={style.routingPolicyInformation}
      />
      <GatewayVisibilityForm
        onSubmit={handleDefaultGatewayVisibilitySubmit}
        initialValues={initialValues}
        error={formError}
      />
    </Col>
  )
}

export default DefaultGatewayVisibilityView
