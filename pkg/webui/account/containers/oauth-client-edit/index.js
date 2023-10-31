

import React, { useState, useCallback } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { useNavigate } from 'react-router-dom'
import { defineMessages } from 'react-intl'

import toast from '@ttn-lw/components/toast'

import OAuthClientForm from '@account/components/oauth-client-form'

import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'
import PropTypes from '@ttn-lw/lib/prop-types'
import diff from '@ttn-lw/lib/diff'

import { deleteClient, updateClient } from '@account/store/actions/clients'

import { selectIsConfiguration } from '@account/store/selectors/identity-server'

const m = defineMessages({
  deleteSuccess: 'OAuth client deleted',
  deleteFail: 'There was an error and the OAuth client could not be deleted',
  updateSuccess: 'OAuth client updated',
  updateFailure: 'There was an error updating this client',
})

const checkChanged = (changed, values) => {
  if ('redirect_uris' in changed) {
    return {
      ...changed,
      redirect_uris: values.redirect_uris,
    }
  } else if ('logout_redirect_uris' in changed) {
    return {
      ...changed,
      logout_redirect_uris: values.logout_redirect_uris,
    }
  } else if ('grants' in changed) {
    return {
      ...changed,
      grants: values.grants,
    }
  }

  return changed
}

const ClientAdd = props => {
  const { userId, isAdmin, rights, pseudoRights, initialValues } = props

  const [error, setError] = useState()
  const navigate = useNavigate()
  const dispatch = useDispatch()
  const isConfig = useSelector(selectIsConfiguration)
  const isResctrictedUser =
    isConfig && isConfig.collaborator_rights?.set_others_as_contacts === false

  const navigateToOAuthClient = useCallback(
    clientId => {
      navigate(`/oauth-clients/${clientId}`)
    },
    [navigate],
  )
  const handleSubmit = useCallback(
    async (values, resetForm, setSubmitting) => {
      const { client_id } = values.ids
      setError(undefined)

      const changed = diff(initialValues, values)

      // If there is a change in `redirect_uris`, `logout_redirect_uris` or `grants`,
      // copy all values so they don't get overwritten.
      const update = checkChanged(changed, values)

      const { owner_id, ...newClient } = update

      try {
        await dispatch(attachPromise(updateClient(client_id, newClient)))
        resetForm({ values })
        toast({
          title: client_id,
          message: m.updateSuccess,
          type: toast.types.SUCCESS,
        })
      } catch (error) {
        setSubmitting(false)
        setError(error)
        toast({
          title: client_id,
          message: m.updateFailure,
          type: toast.types.ERROR,
        })
      }
    },
    [dispatch, initialValues],
  )

  const handleDelete = useCallback(
    async (shouldPurge, clientId) => {
      setError(undefined)

      try {
        await dispatch(attachPromise(deleteClient(clientId, { purge: shouldPurge || false })))
        navigate('/oauth-clients')
        toast({
          title: clientId,
          message: m.deleteSuccess,
          type: toast.types.SUCCESS,
        })
      } catch (error) {
        setError(error)
        toast({
          title: clientId,
          message: m.deleteFail,
          type: toast.types.ERROR,
        })
      }
    },
    [dispatch, navigate],
  )

  return (
    <OAuthClientForm
      update
      initialValues={initialValues}
      onSubmit={handleSubmit}
      onDelete={handleDelete}
      navigateToOAuthClient={navigateToOAuthClient}
      error={error}
      userId={userId}
      isAdmin={isAdmin}
      rights={rights}
      pseudoRights={pseudoRights}
      isResctrictedUser={isResctrictedUser}
    />
  )
}

ClientAdd.propTypes = {
  initialValues: PropTypes.shape({
    grants: PropTypes.arrayOf(PropTypes.string),
  }).isRequired,
  isAdmin: PropTypes.bool.isRequired,
  pseudoRights: PropTypes.rights.isRequired,
  rights: PropTypes.rights,
  userId: PropTypes.string.isRequired,
}

ClientAdd.defaultProps = {
  rights: undefined,
}

export default ClientAdd
