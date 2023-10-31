

/* eslint-disable import/prefer-default-export */

import React from 'react'
import { Provider } from 'react-redux'
import { ConnectedRouter } from 'connected-react-router'
import { IntlProvider } from 'react-intl'
import { createMemoryHistory } from 'history'

import messages from '@ttn-lw/locales/en.json'
import backendMessages from '@ttn-lw/locales/.backend/en.json'

import { EnvProvider } from '@ttn-lw/lib/components/env'

import '../../pkg/webui/styles/main.styl'
import 'focus-visible/dist/focus-visible'
import createStore from './store'
import Center from './center'
import env from './env'

const history = createMemoryHistory()
const store = createStore(history)

export const decorators = [
  Story => (
    <EnvProvider env={env}>
      <Provider store={store}>
        <IntlProvider key="key" messages={{ ...messages, ...backendMessages }} locale="en-US">
          <ConnectedRouter history={history}>
            <Center>{Story()}</Center>
          </ConnectedRouter>
        </IntlProvider>
      </Provider>
    </EnvProvider>
  ),
]
