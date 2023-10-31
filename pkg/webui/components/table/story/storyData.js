

import React from 'react'

import Button from '@ttn-lw/components/button'

const headers = [
  {
    name: 'appId',
    displayName: 'Application ID',
  },
  {
    name: 'desc',
    displayName: 'Description',
  },
  {
    name: 'devices',
    displayName: 'Devices',
    align: 'center',
  },
  {
    name: 'lastActivity',
    displayName: 'Last Activity',
  },
]

const rows = [
  {
    appId: 'my-app1',
    desc: 'Test Application',
    devices: '1',
    lastActivity: '10 sec. ago',
    tabs: ['all', 'starred'],
    clickable: false,
  },
  {
    appId: 'my-app2',
    desc: 'Test Application 2',
    devices: '2',
    lastActivity: '10 sec. ago',
    tabs: ['all'],
    clickable: false,
  },
  {
    appId: 'my-app3',
    desc: 'Test Application 3',
    devices: '3',
    lastActivity: '10 sec. ago',
    tabs: ['all', 'starred'],
    clickable: false,
  },
  {
    appId: 'my-app4',
    desc: 'Test Application 4',
    devices: '5',
    lastActivity: '10 sec. ago',
    tabs: ['all', 'starred'],
    clickable: false,
  },
  {
    appId: 'my-app5',
    desc: 'Test Application 5',
    devices: '4',
    lastActivity: '10 sec. ago',
    tabs: ['all', 'starred'],
    clickable: false,
  },
  {
    appId: 'my-app6',
    desc: 'Test Application 6',
    devices: '3',
    lastActivity: '10 sec. ago',
    tabs: ['all'],
    clickable: false,
  },
]

export default {
  defaultExample: {
    headers,
    rows,
  },
  loadingExample: {
    headers,
    rows,
  },
  paginatedExample: {
    headers,
    rows: rows.concat(rows).concat(rows).concat(rows).concat(rows),
  },
  clickableRowsExample: {
    headers,
    rows: rows.slice(0, 3).map(row =>
      Object.assign({}, row, {
        clickable: true,
      }),
    ),
  },
  customCellExample: {
    headers: [
      ...headers,
      {
        name: 'options',
        displayName: 'Options',
        align: 'center',
      },
    ],
    rows: rows.map(r =>
      Object.assign({}, r, {
        options: (
          <div>
            <Button icon="settings" />
            <Button danger icon="delete" />
          </div>
        ),
      }),
    ),
  },
  sortableExample: {
    headers: headers.map(header => {
      if (header.name === 'devices' || header.name === 'appId') {
        return Object.assign({}, header, {
          sortable: true,
        })
      }

      return header
    }),
    rows,
  },
  emptyExample: {
    headers,
    rows: [],
  },
  customWrapperExample: {
    headers,
    rows,
  },
}
