

import React from 'react'
import { defineMessages, useIntl } from 'react-intl'

import Tag from '@ttn-lw/components/tag'
import TagGroup from '@ttn-lw/components/tag/group'

import FetchTable from '@ttn-lw/containers/fetch-table'

import DateTime from '@ttn-lw/lib/components/date-time'
import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import style from './api-keys-table.styl'

const m = defineMessages({
  grantedRights: 'Granted Rights',
})

const RIGHT_TAG_MAX_WIDTH = 160

const ApiKeysTable = props => {
  const { pageSize, baseDataSelector, getItemsAction } = props
  const intl = useIntl()

  const headers = [
    {
      name: 'id',
      displayName: sharedMessages.keyId,
      width: 20,
      sortKey: 'api_key_id',
      render: id => <span className={style.keyId}>{id}</span>,
      sortable: true,
    },
    {
      name: 'name',
      displayName: sharedMessages.name,
      width: 20,
      sortable: true,
    },
    {
      name: 'rights',
      displayName: m.grantedRights,
      width: 50,
      render: (rights = []) => {
        if (rights.length === 0) {
          return <Message className={style.none} content={sharedMessages.none} lowercase />
        }
        const tags = rights.map(r => {
          let rightLabel = intl.formatMessage({ id: `enum:${r}` })
          rightLabel = rightLabel.charAt(0).toUpperCase() + rightLabel.slice(1)
          return <Tag className={style.rightTag} content={rightLabel} key={r} />
        })

        return <TagGroup tagMaxWidth={RIGHT_TAG_MAX_WIDTH} tags={tags} />
      },
    },
    {
      name: 'created_at',
      displayName: sharedMessages.createdAt,
      sortable: true,
      width: 10,
      render: date => <DateTime.Relative value={date} />,
    },
  ]

  return (
    <FetchTable
      entity="keys"
      defaultOrder="-created_at"
      headers={headers}
      addMessage={sharedMessages.addApiKey}
      pageSize={pageSize}
      baseDataSelector={baseDataSelector}
      getItemsAction={getItemsAction}
      tableTitle={<Message content={sharedMessages.apiKeys} />}
    />
  )
}

ApiKeysTable.propTypes = {
  baseDataSelector: PropTypes.func.isRequired,
  getItemsAction: PropTypes.func.isRequired,
  pageSize: PropTypes.number,
}

ApiKeysTable.defaultProps = {
  pageSize: undefined,
}

export default ApiKeysTable
