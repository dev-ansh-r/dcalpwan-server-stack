

import React from 'react'
import DOM from 'react-dom'
import { Provider } from 'react-redux'
import * as Sentry from '@sentry/browser'

import sentryConfig from '@ttn-lw/constants/sentry'
import store, { history } from '@account/store'

import { BreadcrumbsProvider } from '@ttn-lw/components/breadcrumbs/context'
import Header from '@ttn-lw/components/header'

import WithLocale from '@ttn-lw/lib/components/with-locale'
import { ErrorView } from '@ttn-lw/lib/components/error-view'
import { FullViewError } from '@ttn-lw/lib/components/full-view-error'
import Init from '@ttn-lw/lib/components/init'

import Logo from '@account/containers/logo'

import App from '@account/views/app'

import { selectSentryDsnConfig } from '@ttn-lw/lib/selectors/env'

import '@ttn-lw/lib/yup'

// Initialize sentry before creating store
if (selectSentryDsnConfig()) {
  Sentry.init(sentryConfig)
}

// Error renderer for the outermost error boundary.
// Do not use any components that depend on context
// e.g. Intl, Router, Redux store.
const errorRender = error => (
  <FullViewError error={error} header={<Header logo={<Logo safe />} />} safe />
)

DOM.render(
  <ErrorView errorRender={errorRender}>
    <Provider store={store}>
      <WithLocale>
        <div id="modal-container" />
        <Init>
          <BreadcrumbsProvider>
            <App history={history} />
          </BreadcrumbsProvider>
        </Init>
      </WithLocale>
    </Provider>
  </ErrorView>,
  document.getElementById('app'),
)
