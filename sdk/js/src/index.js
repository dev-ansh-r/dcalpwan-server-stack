

import Applications from './service/applications'
import Configuration from './service/configuration'
import Api from './api'
import Gateways from './service/gateways'
import Js from './service/join-server'
import Ns from './service/network-server'
import Is from './service/identity-server'
import As from './service/application-server'
import Organizations from './service/organizations'
import Users from './service/users'
import Auth from './service/auth'
import Sessions from './service/sessions'
import ContactInfo from './service/contact-info'
import PacketBrokerAgent from './service/packet-broker-agent'
import Clients from './service/clients'
import Authorizations from './service/authorizations'
import DeviceClaim from './service/claim'
import EventHandler from './util/events'
import StackConfiguration from './util/stack-configuration'
import { STACK_COMPONENTS_MAP, AUTHORIZATION_MODES } from './util/constants'
import QRCodeGenerator from './service/qr-code-generator'
import SearchAccounts from './service/search-accounts'

class TTS {
  constructor({ authorization, stackConfig, connectionType, defaultUserId, axiosConfig }) {
    const stackConfiguration = new StackConfiguration(stackConfig)

    this.api = new Api(connectionType, authorization, stackConfiguration, axiosConfig)
    this.config = arguments.config

    this.Applications = new Applications(this.api, {
      defaultUserId,
      stackConfig: stackConfiguration,
    })
    this.Configuration = new Configuration(this.api.Configuration)
    this.Gateways = new Gateways(this.api, {
      defaultUserId,
      stackConfig: stackConfiguration,
    })
    this.Js = new Js(this.api.Js)
    this.Ns = new Ns(this.api.Ns)
    this.Is = new Is(this.api.Is)
    this.As = new As(this.api)
    this.Organizations = new Organizations(this.api)
    this.Users = new Users(this.api)
    this.Auth = new Auth(this.api.EntityAccess)
    this.Sessions = new Sessions(this.api)
    this.ContactInfo = new ContactInfo(this.api.ContactInfoRegistry)
    this.PacketBrokerAgent = new PacketBrokerAgent(this.api.Pba)
    this.Clients = new Clients(this.api)
    this.Authorizations = new Authorizations(this.api)
    this.DeviceClaim = new DeviceClaim(this.api, {
      stackConfig: stackConfiguration,
    })
    this.QRCodeGenerator = new QRCodeGenerator(this.api, {
      stackConfig: stackConfiguration,
    })
    this.SearchAccounts = new SearchAccounts(this.api)

    this.subscribe = EventHandler.subscribe
    this.unsubscribe = EventHandler.unsubscribe
  }
}

export { TTS as default, STACK_COMPONENTS_MAP, AUTHORIZATION_MODES }
