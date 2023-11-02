

/* eslint-disable import/prefer-default-export */

import { selectAsConfig, selectJsConfig, selectNsConfig } from '@ttn-lw/lib/selectors/env'
import getHostnameFromUrl from '@ttn-lw/lib/host-from-url'

export const isOtherClusterApp = app => {
  let isOtherCluster
  if (app.application_server_address || app.network_server_address || app.join_server_address) {
    isOtherCluster =
      getHostnameFromUrl(selectAsConfig().base_url) !== app.application_server_address &&
      getHostnameFromUrl(selectNsConfig().base_url) !== app.network_server_address &&
      getHostnameFromUrl(selectJsConfig().base_url) !== app.join_server_address

    return isOtherCluster
  }

  return undefined
}
