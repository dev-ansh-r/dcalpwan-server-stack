

import { isPlainObject } from 'lodash'

import { TokenError } from './errors/custom-errors'
import * as cache from './cache'

export default fetchToken => {
  let tokenPromise
  let finishedRetrieval = true

  const retrieveToken = async () => {
    try {
      const response = await fetchToken()
      const token = response.data
      if (!isPlainObject(token) || !('access_token' in token)) {
        throw new TokenError('Received invalid token')
      }

      cache.set('accessToken', token)

      return token
    } catch (error) {
      throw new TokenError('Could not fetch token', error)
    } finally {
      finishedRetrieval = true
    }
  }

  return () => {
    const token = cache.get('accessToken')

    if (!token || Date.parse(token.expiry) < Date.now()) {
      // If we don't have a token stored or it's expired, we want to retrieve it.

      // Prevent issuing more than one request at a time.
      if (finishedRetrieval) {
        finishedRetrieval = false

        // Remove stored, invalid token.
        clear()

        // Retrieve new token and store it.
        tokenPromise = retrieveToken()
      }
      return tokenPromise
    }

    // If we have a stored token and its valid, we want to use it.
    return token
  }
}

export const clear = () => {
  cache.remove('accessToken')
}
