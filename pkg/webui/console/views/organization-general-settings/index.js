

import React, { useCallback, useEffect, useState } from 'react'
import { Col, Row, Container } from 'react-grid-system'
import { useDispatch, useSelector } from 'react-redux'
import { useNavigate } from 'react-router-dom'

import PageTitle from '@ttn-lw/components/page-title'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import Spinner from '@ttn-lw/components/spinner'

import { FullViewErrorInner } from '@ttn-lw/lib/components/full-view-error'
import Message from '@ttn-lw/lib/components/message'
import RequireRequest from '@ttn-lw/lib/components/require-request'

import OrganizationUpdateForm from '@console/containers/organization-form/update'

import Require from '@console/lib/components/require'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import { getCollaboratorsList } from '@ttn-lw/lib/store/actions/collaborators'
import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'

import {
  mayEditBasicOrganizationInformation,
  mayDeleteOrganization,
  mayViewOrEditOrganizationApiKeys,
  mayViewOrEditOrganizationCollaborators,
} from '@console/lib/feature-checks'
import { checkFromState } from '@account/lib/feature-checks'

import { getApiKeysList } from '@console/store/actions/api-keys'
import { getIsConfiguration } from '@console/store/actions/identity-server'

import { selectSelectedOrganizationId } from '@console/store/selectors/organizations'

const GeneralSettings = () => {
  const [error, setError] = useState()
  const [fetching, setFetching] = useState(true)
  const dispatch = useDispatch()
  const navigate = useNavigate()
  const mayDeleteOrg = useSelector(state => checkFromState(mayDeleteOrganization, state))
  const mayViewApiKeys = useSelector(state =>
    checkFromState(mayViewOrEditOrganizationApiKeys, state),
  )
  const mayViewCollaborators = useSelector(state =>
    checkFromState(mayViewOrEditOrganizationCollaborators, state),
  )
  const orgId = useSelector(selectSelectedOrganizationId)

  useBreadcrumbs(
    'orgs.single.general-settings',
    <Breadcrumb
      path={`/organizations/${orgId}/general-settings`}
      content={sharedMessages.generalSettings}
    />,
  )

  // Conditionally load API Keys and Collaborators to determine whether
  // deleting is possible.
  useEffect(() => {
    try {
      if (mayDeleteOrg) {
        if (mayViewApiKeys) {
          dispatch(attachPromise(getApiKeysList('organization', orgId)))
        }

        if (mayViewCollaborators) {
          dispatch(attachPromise(getCollaboratorsList('organization', orgId)))
        }
      }
    } catch (error) {
      setError(error)
    }
    setFetching(false)
  }, [dispatch, mayDeleteOrg, mayViewApiKeys, mayViewCollaborators, orgId])

  const handleDeleteSuccess = useCallback(() => navigate(`/organizations`), [navigate])

  if (fetching) {
    return (
      <Spinner inline center>
        <Message content={sharedMessages.fetching} />
      </Spinner>
    )
  }

  if (error) {
    return <FullViewErrorInner error={error} />
  }

  return (
    <Require
      featureCheck={mayEditBasicOrganizationInformation}
      otherwise={{ redirect: `/organizations/${orgId}` }}
    >
      <RequireRequest requestAction={getIsConfiguration()}>
        <Container>
          <PageTitle title={sharedMessages.generalSettings} />
          <Row>
            <Col lg={8} md={12}>
              <OrganizationUpdateForm onDeleteSuccess={handleDeleteSuccess} />
            </Col>
          </Row>
        </Container>
      </RequireRequest>
    </Require>
  )
}

export default GeneralSettings
