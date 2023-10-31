

import React, { useCallback } from 'react'
import { defineMessages } from 'react-intl'
import { useParams } from 'react-router-dom'
import { createSelector } from 'reselect'

import FetchTable from '@ttn-lw/containers/fetch-table'

import Message from '@ttn-lw/lib/components/message'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { natsUrl as natsUrlRegexp } from '@console/lib/regexp'

import { getPubsubsList } from '@console/store/actions/pubsubs'

import { selectPubsubs, selectPubsubsTotalCount } from '@console/store/selectors/pubsubs'

const m = defineMessages({
  host: 'Server host',
})

const headers = [
  {
    name: 'ids.pub_sub_id',
    displayName: sharedMessages.id,
    width: 40,
    sortable: true,
  },
  {
    getValue: row => {
      if (row.nats) {
        const res = row.nats.server_url.match(natsUrlRegexp)
        return res ? res[8] : ''
      } else if (row.mqtt) {
        return row.mqtt.server_url
      }
      return ''
    },
    displayName: m.host,
    width: 33,
  },
  {
    name: 'base_topic',
    displayName: sharedMessages.pubsubBaseTopic,
    width: 9,
    sortable: true,
  },
  {
    getValue: row => {
      if (row.nats) {
        return 'NATS'
      } else if (row.mqtt) {
        return 'MQTT'
      }
      return 'Not set'
    },
    displayName: sharedMessages.provider,
    width: 9,
  },
  {
    name: 'format',
    displayName: sharedMessages.format,
    width: 9,
    sortable: true,
  },
]

const getItemPathPrefix = item => `${item.ids.pub_sub_id}`

const PubsubsTable = () => {
  const { appId } = useParams()

  const getItemsAction = useCallback(() => getPubsubsList(appId), [appId])

  const baseDataSelector = createSelector(
    [selectPubsubs, selectPubsubsTotalCount],
    (pubsubs, totalCount) => ({
      pubsubs,
      totalCount,
    }),
  )

  return (
    <FetchTable
      entity="pubsubs"
      defaultOrder="-ids.pub_sub_id"
      addMessage={sharedMessages.addPubsub}
      headers={headers}
      getItemsAction={getItemsAction}
      baseDataSelector={baseDataSelector}
      tableTitle={<Message content={sharedMessages.pubsubs} />}
      getItemPathPrefix={getItemPathPrefix}
      paginated={false}
      handlesSorting
    />
  )
}

export default PubsubsTable
