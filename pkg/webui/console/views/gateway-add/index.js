

import React, { useCallback } from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { defineMessages } from 'react-intl'
import { useNavigate } from 'react-router-dom'

import PageTitle from '@ttn-lw/components/page-title'
import Link from '@ttn-lw/components/link'

import Message from '@ttn-lw/lib/components/message'
import RequireRequest from '@ttn-lw/lib/components/require-request'

import GatewayOnboardingForm from '@console/containers/gateway-onboarding-form'

import Require from '@console/lib/components/require'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { mayCreateGateways } from '@console/lib/feature-checks'

import { getOrganizationsList } from '@console/store/actions/organizations'

const m = defineMessages({
  gtwOnboardingDescription:
    'Register your gateway to enable data traffic between nearby end devices and the network. {break} Learn more in our guide on <Link>Adding Gateways</Link>.',
})

const GatewayGuideLink = content => (
  <Link.DocLink secondary path="/gateways/adding-gateways">
    {content}
  </Link.DocLink>
)

const GatewayAdd = () => {
  const navigate = useNavigate()
  const handleSuccess = useCallback(
    gtwId => {
      navigate(`/gateways/${gtwId}`)
    },
    [navigate],
  )

  return (
    <Require featureCheck={mayCreateGateways} otherwise={{ redirect: '/gateways' }}>
      <RequireRequest requestAction={getOrganizationsList()}>
        <Container>
          <PageTitle
            colProps={{ md: 10, lg: 9 }}
            className="mb-cs-s"
            title={sharedMessages.registerGateway}
          >
            <Message
              component="p"
              content={m.gtwOnboardingDescription}
              values={{ Link: GatewayGuideLink, break: <br /> }}
            />
            <hr className="mb-ls-s" />
          </PageTitle>
          <Row>
            <Col md={10} lg={9}>
              <GatewayOnboardingForm onSuccess={handleSuccess} />
            </Col>
          </Row>
        </Container>
      </RequireRequest>
    </Require>
  )
}

export default GatewayAdd
