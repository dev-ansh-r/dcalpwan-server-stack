

import * as Integrations from '@sentry/integrations'

import env from '@ttn-lw/lib/env'

const sentryConfig = {
  dsn: env.sentryDsn,
  release: process.env.VERSION,
  normalizeDepth: 10,
  integrations: [new Integrations.Dedupe()],
  beforeSend: event => {
    if (event.extra.state && event.extra.state.user) {
      delete event.extra.state.user.user.name
    }
    return event
  },
}

export default sentryConfig
