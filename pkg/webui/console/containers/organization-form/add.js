

import React, { useCallback, useState } from 'react'
import { useDispatch } from 'react-redux'

import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'
import { getOrganizationId } from '@ttn-lw/lib/selectors/id'
import PropTypes from '@ttn-lw/lib/prop-types'

import { createOrganization } from '@console/store/actions/organizations'

import OrganizationForm from './form'

const OrganizationAddForm = ({ onSuccess }) => {
  const [error, setError] = useState(undefined)
  const dispatch = useDispatch()

  const handleSubmit = useCallback(
    async values => {
      try {
        const organization = await dispatch(attachPromise(createOrganization(values)))
        onSuccess(getOrganizationId(organization))
      } catch (error) {
        setError(error)
      }
    },
    [dispatch, onSuccess],
  )

  return <OrganizationForm onSubmit={handleSubmit} error={error} />
}

OrganizationAddForm.propTypes = {
  onSuccess: PropTypes.func.isRequired,
}

export default OrganizationAddForm
