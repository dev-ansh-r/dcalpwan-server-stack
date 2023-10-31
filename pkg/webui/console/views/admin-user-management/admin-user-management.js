

import React from 'react'

import PageTitle from '@ttn-lw/components/page-title'

import UsersTable from '@console/containers/users-table'

import sharedMessages from '@ttn-lw/lib/shared-messages'

const UserManagementAdd = () => (
  <>
    <PageTitle title={sharedMessages.userManagement} className="panel-title mb-0" />
    <UsersTable />
  </>
)

export default UserManagementAdd
