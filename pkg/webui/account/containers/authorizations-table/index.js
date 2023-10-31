

import React from 'react'
import { useSelector } from 'react-redux'
import { createSelector } from 'reselect'

import FetchTable from '@ttn-lw/containers/fetch-table'

import Message from '@ttn-lw/lib/components/message'
import DateTime from '@ttn-lw/lib/components/date-time'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { getAuthorizationsList } from '@account/store/actions/authorizations'

import {
  selectAuthorizations,
  selectAuthorizationsTotalCount,
} from '@account/store/selectors/authorizations'
import { selectUserId } from '@account/store/selectors/user'

const getItemPathPrefix = item => `${item.client_ids.client_id}`

const OAuthClientAuthorizationsTable = () => {
  const userId = useSelector(selectUserId)

  const headers = React.useMemo(() => {
    const baseHeaders = [
      {
        name: 'client_ids.client_id',
        displayName: sharedMessages.clientId,
        width: 20,
      },
      {
        name: 'user_ids.user_id',
        displayName: sharedMessages.userId,
        width: 20,
      },
      {
        name: 'created_at',
        displayName: sharedMessages.created,
        width: 40,
        sortable: true,
        render: created_at => <DateTime.Relative value={created_at} />,
      },
    ]
    return baseHeaders
  }, [])

  const baseDataSelector = createSelector(
    [selectAuthorizations, selectAuthorizationsTotalCount],
    (authorizations, totalCount) => ({
      authorizations,
      totalCount,
      mayAdd: false,
    }),
  )

  const getItems = React.useCallback(filters => getAuthorizationsList(userId, filters), [userId])

  return (
    <FetchTable
      entity="authorizations"
      defaultOrder="-created_at"
      headers={headers}
      getItemsAction={getItems}
      baseDataSelector={baseDataSelector}
      getItemPathPrefix={getItemPathPrefix}
      tableTitle={<Message content={sharedMessages.oauthClientAuthorizations} />}
      clickable
    />
  )
}

export default OAuthClientAuthorizationsTable
