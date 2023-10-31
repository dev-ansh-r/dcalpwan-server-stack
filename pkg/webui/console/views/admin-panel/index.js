

import React from 'react'
import { useSelector } from 'react-redux'
import { defineMessages } from '@formatjs/intl'

import Breadcrumbs from '@ttn-lw/components/breadcrumbs'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'

import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'

import PanelView from '@console/components/panel-view'

import Require from '@console/lib/components/require'

import UserManagement from '@console/views/admin-user-management'
import PacketBrokerRouter from '@console/views/admin-packet-broker'
import NetworkInformation from '@console/views/admin-panel-network-information'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import {
  checkFromState,
  mayConfigurePacketBroker,
  mayManageUsers,
  mayPerformAdminActions,
} from '@console/lib/feature-checks'

const m = defineMessages({
  adminPanel: 'Admin panel',
  userManagement: 'User management',
  globalNetworkSettings: 'Global network settings',
  peeringSettings: 'Peering settings',
})

const AdminPanel = () => {
  useBreadcrumbs('admin-panel', <Breadcrumb path="/admin-panel" content={m.adminPanel} />)
  const showUserManagement = useSelector(state => checkFromState(mayManageUsers, state))
  const showPacketBroker = useSelector(state => checkFromState(mayConfigurePacketBroker, state))

  return (
    <Require featureCheck={mayPerformAdminActions} otherwise={{ redirect: '/' }}>
      <Breadcrumbs />
      <IntlHelmet title={m.adminPanel} />
      <PanelView>
        <PanelView.Item
          title={sharedMessages.networkInformation}
          icon="view_compact"
          path="network-information"
          Component={NetworkInformation}
          exact
        />
        {showUserManagement && (
          <PanelView.Item
            title={m.userManagement}
            icon="user_management"
            path="user-management"
            Component={UserManagement}
            condition={showUserManagement}
          />
        )}
        {showPacketBroker && (
          <PanelView.Item
            title={m.peeringSettings}
            icon="packet_broker"
            path="packet-broker"
            Component={PacketBrokerRouter}
            condition={showPacketBroker}
          />
        )}
      </PanelView>
    </Require>
  )
}

export default AdminPanel
