

import React, { useCallback } from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { defineMessages } from 'react-intl'
import { useNavigate } from 'react-router-dom'

import PageTitle from '@ttn-lw/components/page-title'
import Link from '@ttn-lw/components/link'

import RequireRequest from '@ttn-lw/lib/components/require-request'
import Message from '@ttn-lw/lib/components/message'

import ApplicationForm from '@console/containers/applications-form'

import Require from '@console/lib/components/require'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { mayCreateApplications } from '@console/lib/feature-checks'

import { getOrganizationsList } from '@console/store/actions/organizations'

const m = defineMessages({
  appDescription:
    'Within applications, you can register and manage end devices and their network data. After setting up your device fleet, use one of our many integration options to pass relevant data to your external services.{break}Learn more in our guide on <Link>Adding Applications</Link>.',
})

const ApplicationAdd = () => {
  const navigate = useNavigate()
  const handleSuccess = useCallback(
    gtwId => {
      navigate(`/applications/${gtwId}`)
    },
    [navigate],
  )

  return (
    <Require featureCheck={mayCreateApplications} otherwise={{ redirect: '/applications' }}>
      <RequireRequest requestAction={getOrganizationsList()}>
        <Container>
          <PageTitle
            colProps={{ md: 10, lg: 9 }}
            className="mb-cs-s"
            title={sharedMessages.createApplication}
          >
            <Message
              component="p"
              content={m.appDescription}
              values={{
                Link: content => (
                  <Link.DocLink secondary path="/integrations/adding-applications">
                    {content}
                  </Link.DocLink>
                ),
                break: <br />,
              }}
            />
            <hr className="mb-ls-s" />
          </PageTitle>
          <Row>
            <Col md={10} lg={9}>
              <ApplicationForm onSuccess={handleSuccess} />
            </Col>
          </Row>
        </Container>
      </RequireRequest>
    </Require>
  )
}

export default ApplicationAdd
