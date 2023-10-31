

import React from 'react'
import { Col, Row, Container } from 'react-grid-system'

import PageTitle from '@ttn-lw/components/page-title'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { withBreadcrumb } from '@ttn-lw/components/breadcrumbs/context'

import GatewayLocationForm from '@console/containers/gateway-location-form'

import withFeatureRequirement from '@console/lib/components/with-feature-requirement'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { mayViewOrEditGatewayLocation } from '@console/lib/feature-checks'

const GatewayLocation = () => (
  <Container>
    <PageTitle title={sharedMessages.location} />
    <Row>
      <Col lg={8} md={12}>
        <GatewayLocationForm />
      </Col>
    </Row>
  </Container>
)

export default withBreadcrumb('gateway.single.data', props => {
  const { gtwId } = props
  return <Breadcrumb path={`/gateways/${gtwId}/location`} content={sharedMessages.location} />
})(
  withFeatureRequirement(mayViewOrEditGatewayLocation, {
    redirect: ({ gtwId }) => `/gateways/${gtwId}`,
  })(GatewayLocation),
)
