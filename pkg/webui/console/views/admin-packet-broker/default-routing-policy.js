
import React, { useCallback, useState } from 'react'
import { Col } from 'react-grid-system'
import { useSelector, useDispatch } from 'react-redux'

import toast from '@ttn-lw/components/toast'

import Message from '@ttn-lw/lib/components/message'

import RoutingPolicyForm from '@console/components/routing-policy-form'

import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'

import { isValidPolicy } from '@console/lib/packet-broker/utils'

import {
  setHomeNetworkDefaultRoutingPolicy,
  deleteHomeNetworkDefaultRoutingPolicy,
} from '@console/store/actions/packet-broker'

import { selectHomeNetworkDefaultRoutingPolicy } from '@console/store/selectors/packet-broker'

import m from './messages'

import style from './admin-packet-broker.styl'

const DefaultRoutingPolicyView = () => {
  const dispatch = useDispatch()
  const defaultRoutingPolicy = useSelector(selectHomeNetworkDefaultRoutingPolicy)
  const initialValues = { _use_default_policy: isValidPolicy(defaultRoutingPolicy) }
  initialValues.policy = initialValues._use_default_policy
    ? defaultRoutingPolicy
    : { uplink: {}, downlink: {} }
  const [formError, setFormError] = useState(undefined)
  const handleDefaultRoutingPolicySubmit = useCallback(
    async ({ _use_default_policy, policy }) => {
      try {
        if (_use_default_policy) {
          await dispatch(attachPromise(setHomeNetworkDefaultRoutingPolicy(policy)))
        } else {
          await dispatch(attachPromise(deleteHomeNetworkDefaultRoutingPolicy()))
        }
        toast({
          message: m.defaultRoutingPolicySet,
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
        content={m.routingPolicyInformation}
        component="p"
        className={style.routingPolicyInformation}
      />
      <RoutingPolicyForm
        onSubmit={handleDefaultRoutingPolicySubmit}
        initialValues={initialValues}
        error={formError}
      />
    </Col>
  )
}

export default DefaultRoutingPolicyView
