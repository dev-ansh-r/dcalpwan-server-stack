

import React, { useCallback } from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { defineMessages } from 'react-intl'
import { useNavigate } from 'react-router-dom'

import PageTitle from '@ttn-lw/components/page-title'
import Link from '@ttn-lw/components/link'

import Message from '@ttn-lw/lib/components/message'

import OrganizationAddForm from '@console/containers/organization-form/add'

import Require from '@console/lib/components/require'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { mayCreateOrganizations } from '@console/lib/feature-checks'

const m = defineMessages({
  orgDescription:
    'Organizations are used to group multiple users and assigning collective rights for them. An organization can then be set as collaborator of applications or gateways. This makes it easy to grant or revoke rights to entities for a group of users.{break} Learn more in our guide on <Link>Organization Management</Link>.',
})

const OrganizationAdd = () => {
  const navigate = useNavigate()

  const handleSuccess = useCallback(
    orgId => {
      navigate(`/organizations/${orgId}`)
    },
    [navigate],
  )

  return (
    <Require featureCheck={mayCreateOrganizations} otherwise={{ redirect: '/organizations' }}>
      <Container>
        <PageTitle
          colProps={{ md: 10, lg: 9 }}
          className="mb-cs-s"
          title={sharedMessages.createOrganization}
        >
          <Message
            component="p"
            content={m.orgDescription}
            values={{
              Link: content => (
                <Link.DocLink secondary path="/the-things-stack/management/user-management/org/">
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
            <OrganizationAddForm onSuccess={handleSuccess} />
          </Col>
        </Row>
      </Container>
    </Require>
  )
}

export default OrganizationAdd
