

import React from 'react'

import Notification from '.'

export default {
  title: 'Notification/Notification',
}

export const Default = () => (
  <div>
    <Notification title="example message title" content="This is an example message" />
    <Notification content="This is an example message" />
    <Notification title="example message title" content="This is an example message" small />
    <Notification content="This is an example message" small />
  </div>
)

export const Info = () => (
  <div>
    <Notification title="example message title" info content="This message is good to know" />
    <Notification info content="This message is good to know" />
    <Notification title="example message title" info content="This message is good to know" small />
    <Notification info content="This message is good to know" small />
  </div>
)

export const Warning = () => (
  <div>
    <Notification title="example message title" warning content="This issue should be addressed!" />
    <Notification warning content="This issue should be addressed!" />
    <Notification
      title="example message title"
      warning
      content="This issue should be addressed!"
      small
    />
    <Notification warning content="This issue should be addressed!" small />
  </div>
)

export const Success = () => (
  <div>
    <Notification title="example message title" success content="Successful action!" />
    <Notification success content="Successful action!" />
    <Notification title="example message title" success content="Successful action!" small />
    <Notification success content="Successful action!" small />
  </div>
)
