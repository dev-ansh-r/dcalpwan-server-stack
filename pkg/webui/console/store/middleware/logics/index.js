

import status from '@ttn-lw/lib/store/logics/status'

import user from './logout'
import users from './users'
import init from './init'
import applications from './applications'
import collaborators from './collaborators'
import claim from './claim'
import devices from './devices'
import gateways from './gateways'
import configuration from './configuration'
import organizations from './organizations'
import js from './join-server'
import apiKeys from './api-keys'
import webhooks from './webhooks'
import pubsubs from './pubsubs'
import applicationPackages from './application-packages'
import is from './identity-server'
import as from './application-server'
import deviceRepository from './device-repository'
import packetBroker from './packet-broker'
import networkServer from './network-server'
import qrCodeGenerator from './qr-code-generator'
import searchAccounts from './search-accounts'

export default [
  ...status,
  ...user,
  ...users,
  ...init,
  ...applications,
  ...claim,
  ...devices,
  ...gateways,
  ...configuration,
  ...organizations,
  ...js,
  ...apiKeys,
  ...webhooks,
  ...pubsubs,
  ...applicationPackages,
  ...is,
  ...as,
  ...deviceRepository,
  ...packetBroker,
  ...collaborators,
  ...networkServer,
  ...qrCodeGenerator,
  ...searchAccounts,
]
