

import React from 'react'
import { Navigate, useSearchParams } from 'react-router-dom'
import { defineMessages } from 'react-intl'
import { Container, Col, Row } from 'react-grid-system'

import SafeInspector from '@ttn-lw/components/safe-inspector'
import Button from '@ttn-lw/components/button'
import PageTitle from '@ttn-lw/components/page-title'

import Message from '@ttn-lw/lib/components/message'

import { selectApplicationSiteTitle } from '@ttn-lw/lib/selectors/env'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import style from './code.styl'

const m = defineMessages({
  codeDescription: 'Your authorization code is:',
})

const siteTitle = selectApplicationSiteTitle()

const Code = () => {
  const [searchParams] = useSearchParams()
  const code = searchParams.get('code')

  if (!code) {
    return <Navigate to="/" />
  }

  return (
    <Container>
      <Row>
        <Col lg={4} md={6} sm={12}>
          <PageTitle title={sharedMessages.authorizationCode} />
          <Message
            content={m.codeDescription}
            component="label"
            className={style.codeDescription}
          />
          <SafeInspector
            data={code}
            initiallyVisible
            hideable={false}
            isBytes={false}
            className={style.code}
          />
          <Button.Link
            to="/"
            icon="keyboard_arrow_left"
            message={{ ...sharedMessages.backTo, values: { siteTitle } }}
          />
        </Col>
      </Row>
    </Container>
  )
}

export default Code
