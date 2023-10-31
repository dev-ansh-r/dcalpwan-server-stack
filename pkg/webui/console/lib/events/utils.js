

import { createNetworkErrorEvent, createUnknownErrorEvent } from './definitions'

export const defineSyntheticEvent = name => data => ({
  time: new Date().toISOString(),
  name,
  isError: name.startsWith('synthetic.error'),
  isSynthetic: true,
  unique_id: `synthetic.${Date.now()}`,
  data,
})

export const createSyntheticEventFromError = error => {
  if (error instanceof Error) {
    const errorString = error.toString()
    if (error.message === 'network error' || error.message === 'Error in body stream') {
      return createNetworkErrorEvent({ error: errorString })
    }

    return createUnknownErrorEvent({ error: errorString })
  }
}
