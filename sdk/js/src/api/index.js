

import apiDefinition from '../../generated/api-definition.json'
import { STACK_COMPONENTS } from '../util/constants'

import Http from './http'

/**
 * Api Class is an abstraction on the API connection which can use either the
 * HTTP or gRPC connector to communicate with The Things Stack for LoraWAN API
 * in order to expose the same class API for both.
 */
class Api {
  constructor(connectionType = 'http', authorization, stackConfig, axiosConfig) {
    this.connectionType = connectionType

    if (this.connectionType !== 'http') {
      throw new Error('Only http connection type is supported')
    }

    this._connector = new Http(authorization, stackConfig, axiosConfig)
    const connector = this._connector

    for (const serviceName of Object.keys(apiDefinition)) {
      const service = apiDefinition[serviceName]

      this[serviceName] = {}

      for (const rpcName of Object.keys(service)) {
        const rpc = service[rpcName]

        this[serviceName][rpcName] = ({ routeParams = {}, component } = {}, payload) => {
          const componentType = typeof component
          if (componentType === 'string' && !STACK_COMPONENTS.includes(component)) {
            throw new Error(`Unknown stack component: ${component}`)
          }
          if (component && componentType !== 'string') {
            throw new Error(`Invalid component argument type: ${typeof componentType}`)
          }

          const paramSignature = Object.keys(routeParams).sort().join()
          const endpoint = rpc.http.find(
            prospect => prospect.parameters.sort().join() === paramSignature,
          )

          if (!endpoint) {
            throw new Error(`The parameter signature did not match the one of the rpc.
Rpc: ${serviceName}.${rpcName}()
Signature tried: ${paramSignature}`)
          }

          let route = endpoint.pattern
          const isStream = Boolean(endpoint.stream)

          for (const parameter of endpoint.parameters) {
            route = route.replace(`{${parameter}}`, routeParams[parameter])
          }

          return connector.handleRequest(endpoint.method, route, component, payload, isStream)
        }

        this[serviceName][`${rpcName}AllowedFieldMaskPaths`] = rpc.allowedFieldMaskPaths
      }
    }
  }
}

export default Api
