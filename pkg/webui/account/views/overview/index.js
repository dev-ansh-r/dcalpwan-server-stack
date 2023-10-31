

import React from 'react'
import { useSelector } from 'react-redux'
import { Container, Col, Row } from 'react-grid-system'
import { defineMessages } from 'react-intl'

import PageTitle from '@ttn-lw/components/page-title'
import Button from '@ttn-lw/components/button'

import Message from '@ttn-lw/lib/components/message'

import ProfileCard from '@account/containers/profile-card'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { selectConsoleUrl } from '@account/lib/selectors/app-config'

import { selectUser } from '@account/store/selectors/user'

import style from './overview.styl'

const m = defineMessages({
  accountAppInfoTitle: 'Welcome, {userId}!',
  accountAppInfoMessage: `<p>You have successfully logged into the Account App. The Account App is the official user account management application of The Things Stack.</p>`,
  accountAppConsoleInfo:
    'If you wish to manage your applications, end devices and/or gateways, you can use the button below to head over to the Console.',
  goToConsole: 'Go to the Console',
})

const consoleUrl = selectConsoleUrl()

const Overview = () => {
  const {
    name: userName,
    ids: { user_id: userId },
  } = useSelector(selectUser)

  return (
    <Container>
      <Row>
        <Col className={style.top}>
          <PageTitle title={sharedMessages.overview} hideHeading />
          <ProfileCard />
        </Col>
      </Row>
      <Row justify="center">
        <Col sm={6}>
          <Message
            component="h1"
            content={m.accountAppInfoTitle}
            values={{ userId: userName || userId }}
          />
          <Message
            content={m.accountAppInfoMessage}
            values={{
              p: msg => <p key={msg}>{msg}</p>,
              ul: msg => <ol key="list">{msg}</ol>,
              li: msg => <li>{msg}</li>,
            }}
          />
          <Message component="p" content={m.accountAppConsoleInfo} />
          <Button.AnchorLink primary href={consoleUrl} message={m.goToConsole} />
        </Col>
      </Row>
    </Container>
  )
}

export default Overview
