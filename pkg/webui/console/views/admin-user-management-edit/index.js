

import React from 'react'
import { useParams } from 'react-router-dom'

import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'

import RequireRequest from '@ttn-lw/lib/components/require-request'

import UserDataFormEdit from '@console/containers/user-data-form/edit'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { getUser } from '@console/store/actions/users'

const UserManagementEdit = () => {
  const { userId } = useParams()

  useBreadcrumbs(
    'admin-panel.user-management.edit',
    <Breadcrumb path={`./${userId}`} content={sharedMessages.edit} />,
  )

  return (
    <RequireRequest
      requestAction={getUser(userId, [
        'name',
        'primary_email_address',
        'state',
        'admin',
        'description',
      ])}
    >
      <UserDataFormEdit />
    </RequireRequest>
  )
}

export default UserManagementEdit
