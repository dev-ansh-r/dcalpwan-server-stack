

import React from 'react'
import { withInfo } from '@storybook/addon-info'
import { IntlProvider } from 'react-intl'

import doc from './message.md'

import Message from '.'

const exampleMessage = {
  id: '$.some.id',
  defaultMessage: 'This is the default message',
}

const exampleLowercaseMessage = {
  id: '$.some.other.id',
  defaultMessage: 'this is the default message',
}

const placeholderMessage = {
  id: '$.placeholder',
  defaultMessage: 'There are {number} gateways.',
}

const messages = {
  '$.some.id': 'This is the translated message',
  '$.some.other.id': 'this is the translated message',
  '$.placeholder': 'There are {number} gateways.',
}

const IntlDecorator = storyFn => (
  <IntlProvider key="key" messages={messages} locale="en-US">
    {storyFn()}
  </IntlProvider>
)

export default {
  title: 'Utility Components/Message',

  decorators: [
    withInfo({
      inline: true,
      header: false,
      text: doc,
      propTables: [Message],
    }),
    IntlDecorator,
  ],
}

export const Default = () => <Message content={exampleMessage} />
export const Placeholder = () => <Message content={placeholderMessage} values={{ number: 5 }} />
export const String = () => (
  <Message content="I can also be just a string, but will issue a warning" />
)

export const Transforms = () => (
  <ul>
    <Message capitalize component="li" content={exampleLowercaseMessage} />
    <Message firstToUpper component="li" content={exampleLowercaseMessage} />
    <Message firstToLower component="li" content={exampleMessage} />
    <Message uppercase component="li" content={exampleLowercaseMessage} />
    <Message lowercase component="li" content={exampleMessage} />
  </ul>
)
