

import React from 'react'

import Button from '@ttn-lw/components/button'

import Message from '@ttn-lw/lib/components/message'
import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'

import style from '@account/views/front/front.styl'

import { selectApplicationSiteName, selectApplicationSiteTitle } from '@ttn-lw/lib/selectors/env'
import errorMessages from '@ttn-lw/lib/errors/error-messages'
import statusCodeMessages from '@ttn-lw/lib/errors/status-code-messages'
import sharedMessages from '@ttn-lw/lib/shared-messages'

const siteName = selectApplicationSiteName()
const siteTitle = selectApplicationSiteTitle()

const FrontNotFound = () => (
  <React.Fragment>
    <div className={style.form}>
      <IntlHelmet title={statusCodeMessages['404']} />
      <h1 className={style.title}>
        {siteName}
        <br />
        <Message content={statusCodeMessages['404']} component="strong" />
      </h1>
      <hr className={style.hRule} />
      <Message
        content={errorMessages.genericNotFound}
        component="p"
        className={style.errorDescription}
      />
      <Button.Link
        to="/login"
        icon="keyboard_arrow_left"
        message={{ ...sharedMessages.backTo, values: { siteTitle } }}
      />
    </div>
  </React.Fragment>
)

export default FrontNotFound
